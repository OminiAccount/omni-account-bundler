package state

import (
	"context"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/utils/chains"
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/merkletree"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jackc/pgx/v4/pgxpool"
	"math/big"
	"time"
)

type State struct {
	cfg    Config
	ctx    context.Context
	cancel func()

	ethereum EthereumInterface
	tree     *merkletree.SMT
	his      *HistoryManager
	db       *PostgresStorage
}

func NewState(ctx context.Context, cfg Config, s *merkletree.SMT, ether EthereumInterface, pg *pgxpool.Pool) (*State, error) {
	stateCtx, cancel := context.WithCancel(ctx)
	state := &State{
		cfg:      cfg,
		ctx:      stateCtx,
		cancel:   cancel,
		ethereum: ether,
		tree:     s,
		db:       NewPostgresStorage(pg),
	}
	state.his = NewHistoryManager(ctx, cfg, state)

	state.LoadCache()

	return state, nil
}

func (s *State) Start() {
	log.Info("state start")
	go s.his.Start()
	go func() {
		for {
			time.Sleep(time.Second * 5)
			s.Persistence()
		}
	}()
	go func() {
		testPrk := "82693fc767eb00e3288a01b7516f3a98269882e951be74403e4061d898ea0929"
		testUser := common.HexToAddress("0x7c38C1646213255f62dB509688B8fA062e0Ed8e4")
		var testAcc *common.Address
		adr := s.GetAccountAdr(testUser)
		if adr != nil {
			testAcc = adr
		} else {
			testAcc = s.CreateAccount(testUser)
		}
		for {
			// deposit
			time.Sleep(time.Minute * 5)
			log.Infof("add test deposit...")
			accInfo, _ := s.GetAccountInfo(testUser, *testAcc)
			suo := &SignedUserOperation{
				UserOperation: &UserOperation{
					Uid:            accInfo.Uid,
					Owner:          testUser,
					OperationType:  DepositAction,
					OperationValue: (*hexutil.Big)(big.NewInt(20000000000000000)),
					Sender:         *testAcc,
					Exec: ExecData{
						ChainId:                28516,
						Nonce:                  hexutil.Uint64(accInfo.Chain[28516].Nonce + 1),
						CallData:               hexutil.Bytes(""),
						MainChainGasPrice:      (*hexutil.Big)(big.NewInt(20)),
						MainChainGasLimit:      hexutil.Uint64(10),
						DestChainGasPrice:      (*hexutil.Big)(big.NewInt(20)),
						DestChainGasLimit:      hexutil.Uint64(10),
						ZkVerificationGasLimit: hexutil.Uint64(10),
					},
					InnerExec: ExecData{},
					Status:    PendingStatus,
				},
			}
			privateKey, err := crypto.HexToECDSA(testPrk)
			if err != nil {
				log.Errorf("Error generating private key: %v", err)
				return
			}
			signature, err := crypto.Sign(accounts.TextHash(suo.CalculateEIP712TypeDataHash()), privateKey)
			if err != nil {
				log.Errorf("Failed to sign hash typed data: %v", err)
				return
			}
			signature[crypto.RecoveryIDOffset] += 27
			suo.Signature = signature
			suo.Did = getDepositID(suo)
			dbTx, err := s.db.BeginDBTransaction(s.ctx)
			if err != nil {
				log.Errorf("%+v, use db transaction err: %+v", suo.UserOperation, err)
				return
			}
			tNonce, tGas, err := s.db.GetUserInfoLock(s.ctx, suo.Uid, suo.Exec.ChainId.Uint64(), dbTx)
			if err != nil {
				log.Errorf("%+v, chain: %d, get userinfo err: %+v", suo.UserOperation, suo.Exec.ChainId, err)
				_ = dbTx.Rollback(s.ctx)
				return
			}
			tGas = tGas.Add(tGas, suo.OperationValue.ToInt())
			err = s.db.ModUserInfo(s.ctx, suo.Uid, suo.Exec.ChainId.Uint64(), tNonce+1, tGas.String(), dbTx)
			if err != nil {
				_ = dbTx.Rollback(s.ctx)
				log.Errorf("%+v, chain: %d, mod userinfo err: %+v", suo.UserOperation, suo.Exec.ChainId, err)
				return
			}
			err = s.db.AddOperation(s.ctx, suo.Uid, suo, dbTx)
			if err != nil {
				_ = dbTx.Rollback(s.ctx)
				log.Errorf("%+v, add op err: %+v", suo.UserOperation, err)
				return
			}
			err = dbTx.Commit(s.ctx)
			if err != nil {
				log.Errorf("%+v, add op commit err: %+v", suo.UserOperation, err)
				return
			}

			// withdraw
			time.Sleep(time.Minute * 5)
			log.Infof("add test withdraw...")
			accInfo, _ = s.GetAccountInfo(testUser, *testAcc)
			suo2 := &SignedUserOperation{
				UserOperation: &UserOperation{
					Uid:            accInfo.Uid,
					Owner:          testUser,
					OperationType:  WithdrawAction,
					OperationValue: (*hexutil.Big)(big.NewInt(2000000000000000)),
					Sender:         *testAcc,
					Exec: ExecData{
						ChainId:                28516,
						Nonce:                  hexutil.Uint64(accInfo.Chain[28516].Nonce + 1),
						CallData:               hexutil.Bytes(""),
						MainChainGasPrice:      (*hexutil.Big)(big.NewInt(20)),
						MainChainGasLimit:      hexutil.Uint64(10),
						DestChainGasPrice:      (*hexutil.Big)(big.NewInt(20)),
						DestChainGasLimit:      hexutil.Uint64(10),
						ZkVerificationGasLimit: hexutil.Uint64(10),
					},
					InnerExec: ExecData{},
					Status:    PendingStatus,
				},
			}
			privateKey, err = crypto.HexToECDSA(testPrk)
			if err != nil {
				log.Errorf("Error generating private key: %v", err)
				return
			}
			signature, err = crypto.Sign(accounts.TextHash(suo2.CalculateEIP712TypeDataHash()), privateKey)
			if err != nil {
				log.Errorf("Failed to sign hash typed data: %v", err)
				return
			}
			signature[crypto.RecoveryIDOffset] += 27
			suo2.Signature = signature
			hashBytes := crypto.Keccak256Hash(suo2.Encode())
			suo2.Did = hashBytes.Hex()
			dbTx, err = s.db.BeginDBTransaction(s.ctx)
			if err != nil {
				log.Errorf("%+v, use db transaction err: %+v", suo2.UserOperation, err)
				return
			}
			tNonce, tGas, err = s.db.GetUserInfoLock(s.ctx, suo2.Uid, suo2.Exec.ChainId.Uint64(), dbTx)
			if err != nil {
				log.Errorf("%+v, chain: %d, get userinfo err: %+v", suo2.UserOperation, suo2.Exec.ChainId, err)
				_ = dbTx.Rollback(s.ctx)
				return
			}
			tGas = tGas.Sub(tGas, suo2.OperationValue.ToInt())
			tGas = tGas.Sub(tGas, suo2.UserOperation.CalculateGasUsed())
			err = s.db.ModUserInfo(s.ctx, suo2.Uid, suo2.Exec.ChainId.Uint64(), tNonce+1, tGas.String(), dbTx)
			if err != nil {
				_ = dbTx.Rollback(s.ctx)
				log.Errorf("%+v, chain: %d, mod userinfo err: %+v", suo2.UserOperation, suo2.Exec.ChainId, err)
				return
			}
			err = s.db.AddOperation(s.ctx, suo2.Uid, suo2, dbTx)
			if err != nil {
				_ = dbTx.Rollback(s.ctx)
				log.Errorf("%+v, add op err: %+v", suo2.UserOperation, err)
				return
			}
			err = dbTx.Commit(s.ctx)
			if err != nil {
				log.Errorf("%+v, add op commit err: %+v", suo2.UserOperation, err)
				return
			}
		}
	}()
}

func (s *State) Stop() {
	log.Info("state stop")
	s.cancel()
	time.Sleep(time.Second * 2)
	s.Persistence()
}

func (s *State) IsSupportChain(nid hexutil.Uint64) bool {
	cli := s.ethereum.GetChainCli(chains.ChainId(nid))
	if cli == nil {
		return false
	}
	return true
}

func (s *State) GetHisIns() *HistoryManager {
	return s.his
}

func (s *State) GetTree() *merkletree.SMT {
	return s.tree
}

func (s *State) Persistence() {
	//log.Info("cache state data into disk")
	if err := s.tree.Db.Cache(); err != nil {
		log.Error("cache smt into disk error", "error", err)
	}
}

func (s *State) LoadCache() {
	log.Info("load state data from disk")
	err := s.tree.Db.LoadCache()
	if err != nil {
		log.Fatalf("load smt data from disk error: %v", err)
	}
}
