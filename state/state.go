package state

import (
	"context"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/pool"
	"github.com/OAB/utils/db"
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/merkletree"
	"github.com/OAB/utils/queue"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"math/big"
	"time"
)

type State struct {
	cfg    Config
	ctx    context.Context
	cancel func()

	pool     PoolInterface
	ethereum EthereumInterface
	storage  *Storage
	tree     *merkletree.SMT
	his      *HistoryManager

	proofQueue  *queue.ConcurrentQueue[Batch]
	provenQueue *queue.ConcurrentQueue[ProofResult]
}

func NewState(ctx context.Context, cfg Config, pool PoolInterface, ether EthereumInterface, ethDb ethdb.Database) (*State, error) {
	stateCtx, cancel := context.WithCancel(ctx)
	state := &State{
		cfg:         cfg,
		ctx:         stateCtx,
		cancel:      cancel,
		pool:        pool,
		ethereum:    ether,
		tree:        merkletree.NewSMT(db.NewMemDb(ethDb), false),
		proofQueue:  queue.NewConcurrentQueue[Batch](),
		provenQueue: queue.NewConcurrentQueue[ProofResult](),
		storage:     NewStorage(ethDb),
	}
	state.his = NewHistoryManager(ctx, cfg, state)

	state.LoadCache()

	return state, nil
}

func (s *State) Start() {
	log.Info("state start")
	go s.ProcessBatch()
	go s.ExecuteBatch()
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
		adr, _ := s.GetAccountInfo(testUser)
		if adr != nil {
			testAcc = adr
		} else {
			testAcc = s.CreateAccount(testUser)
		}
		for {
			// deposit
			time.Sleep(time.Minute)
			log.Infof("add test deposit...")
			_, accInfo := s.GetAccountInfo(testUser)
			suo := &pool.SignedUserOperation{
				UserOperation: &pool.UserOperation{
					Owner:          testUser,
					OperationType:  pool.DepositAction,
					OperationValue: (*hexutil.Big)(big.NewInt(20000000000000000)),
					Sender:         *testAcc,
					Exec: pool.ExecData{
						ChainId:                28516,
						Nonce:                  hexutil.Uint64(accInfo.Nonce[28516]+1),
						CallData:               hexutil.Bytes(""),
						MainChainGasPrice:      (*hexutil.Big)(big.NewInt(20)),
						MainChainGasLimit:      hexutil.Uint64(10),
						DestChainGasPrice:      (*hexutil.Big)(big.NewInt(20)),
						DestChainGasLimit:      hexutil.Uint64(10),
						ZkVerificationGasLimit: hexutil.Uint64(10),
					},
					InnerExec: pool.ExecData{},
				},
			}
			privateKey, err := crypto.HexToECDSA(testPrk)
			if err != nil {
				log.Errorf("Error generating private key: %v", err)
				continue
			}
			signature, err := crypto.Sign(accounts.TextHash(suo.CalculateEIP712TypeDataHash()), privateKey)
			if err != nil {
				log.Errorf("Failed to sign hash typed data: %v", err)
			}
			signature[crypto.RecoveryIDOffset] += 27
			suo.Signature = signature
			if err := s.AddAccountGas(suo); err != nil {
				log.Errorf("deposit signedUserOperation, add account gas error: %+v", err)
				return
			}
			s.pool.AddSignedUserOperation(suo)

			// withdraw
			time.Sleep(time.Minute * 10)
			log.Infof("add test withdraw...")
			_, accInfo = s.GetAccountInfo(testUser)
			suo2 := &pool.SignedUserOperation{
				UserOperation: &pool.UserOperation{
					Owner:          testUser,
					OperationType:  pool.WithdrawAction,
					OperationValue: (*hexutil.Big)(big.NewInt(2000000000000000)),
					Sender:         *testAcc,
					Exec: pool.ExecData{
						ChainId:                28516,
						Nonce:                  hexutil.Uint64(accInfo.Nonce[28516]+1),
						CallData:               hexutil.Bytes(""),
						MainChainGasPrice:      (*hexutil.Big)(big.NewInt(20)),
						MainChainGasLimit:      hexutil.Uint64(10),
						DestChainGasPrice:      (*hexutil.Big)(big.NewInt(20)),
						DestChainGasLimit:      hexutil.Uint64(10),
						ZkVerificationGasLimit: hexutil.Uint64(10),
					},
					InnerExec: pool.ExecData{},
				},
			}
			privateKey, err = crypto.HexToECDSA(testPrk)
			if err != nil {
				log.Errorf("Error generating private key: %v", err)
				continue
			}
			signature, err = crypto.Sign(accounts.TextHash(suo2.CalculateEIP712TypeDataHash()), privateKey)
			if err != nil {
				log.Errorf("Failed to sign hash typed data: %v", err)
			}
			signature[crypto.RecoveryIDOffset] += 27
			suo2.Signature = signature
			if err := s.AddAccountGas(suo2); err != nil {
				log.Errorf("withdraw signedUserOperation, add account gas error: %+v", err)
				return
			}
			s.pool.AddSignedUserOperation(suo2)
		}
	}()
}

func (s *State) Stop() {
	log.Info("state stop")
	s.cancel()
	time.Sleep(time.Second * 2)
	s.Persistence()
}

func (s *State) ProcessBatch() {
	log.Info("ProcessBatch start")
	for {
		select {
		case batchContext := <-s.pool.BatchContext():
			log.Debugf("Processing batch, SignedUserOperations count: %d", len(batchContext.SignedUserOperations()))
			err := s.processBatch(batchContext)
			if err != nil {
				log.Error("failed process a batch", "error", err)
			}
		case <-s.ctx.Done():
			if len(s.pool.BatchContext()) > 0 {
				continue
			}
			log.Warn("Stopping PreProcessBatch due to context cancellation")
			return
		}
	}
}

func (s *State) ExecuteBatch() {
	log.Info("ExecuteBatch start")
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			res, err := s.executeBatch()
			if err != nil {
				log.Errorf("send batch to verify error: %s", err.Error())
			}
			if res != nil {
				s.provenQueue.Lpush(*res)
			}
		case <-s.ctx.Done():
			if !s.provenQueue.IsEmpty() {
				continue
			}
			log.Warn("Stopping ExecuteBatch due to context cancellation")
			return
		}
	}
}

func (s *State) GetHisIns() *HistoryManager {
	return s.his
}

func (s *State) Persistence() {
	//log.Info("cache state data into disk")
	if err := s.storage.cache(); err != nil {
		log.Error("cache state data into disk error", "error", err)
	}
	if err := s.tree.Db.Cache(); err != nil {
		log.Error("cache smt into disk error", "error", err)
	}
	if err := s.pool.Cache(); err != nil {
		log.Error("cache pool into disk error", "error", err)
	}
}

func (s *State) LoadCache() {
	log.Info("load state data from disk")
	if err := s.storage.loadCache(); err != nil {
		log.Fatalf("load state data from disk error: %v", err)
	}
	err := s.tree.Db.LoadCache()
	if err != nil {
		log.Fatalf("load smt data from disk error: %v", err)
	}
	for i := s.storage.VerifyBatchNumber + 1; i <= s.storage.BatchNumber; i++ {
		batch, err := s.storage.loadBatch(i)
		if err != nil && err != NotBatch {
			log.Fatalf("load batch info from disk error: %v", err)
		}
		if batch != nil {
			log.Infof("from cache load batch number(%d)", batch.NewNumBatch)
			s.proofQueue.Enqueue(*batch)
		}
	}
}
