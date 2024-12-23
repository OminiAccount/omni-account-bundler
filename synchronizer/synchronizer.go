package synchronizer

import (
	"context"
	"github.com/OAB/etherman"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/state"
	"github.com/OAB/synchronizer/types"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Synchronizer struct {
	ether     types.EthereumInterface
	pool      types.PoolInterface
	state     types.StateInterface
	ctx       context.Context
	cancelCtx context.CancelFunc
	storage   ethdb.Database
	db        *PostgresStorage
}

func NewSynchronizer(ctx context.Context, ethereum types.EthereumInterface, p types.PoolInterface,
	state types.StateInterface, db ethdb.Database, pg *pgxpool.Pool) (*Synchronizer, error) {
	syncCtx, cancel := context.WithCancel(ctx)
	return &Synchronizer{
		ether:     ethereum,
		pool:      p,
		state:     state,
		ctx:       syncCtx,
		cancelCtx: cancel,
		storage:   db,
		db:        NewPostgresStorage(pg),
	}, nil
}

func (s *Synchronizer) Start() {
	log.Info("Sync start......")
	acFunc := func(acc etherman.AccountCreateData) {
		log.Infof("chainID: %d, sync to a new account, user: %s, account: %s",
			acc.ChainID, acc.Owner, acc.Account)
		s.state.CreateAccountInfo(acc.ChainID, state.AccountMapping{
			User:    acc.Owner,
			Account: acc.Account,
		})
	}
	dpFunc := func(dp etherman.DepositData) {
		log.Infof("sync to a new value deposit ticket, data: %+v", dp)
		if dp.Account.Hex() == "" || dp.Account.Hex() == "0x" {
			log.Errorf("invalid deposit operation, data: %+v", dp)
			return
		}
		if dp.User.Hex() == "" || dp.User.Hex() == "0x" {
			log.Errorf("invalid deposit operation, data: %+v", dp)
			return
		}
		if dp.Amount.Uint64() <= 0 {
			log.Errorf("invalid deposit operation, data: %+v", dp)
			return
		}
		ai, err := s.state.GetAccountInfo(dp.User, dp.Account)
		if err != nil {
			log.Errorf("data: %+v, get account err: %+v", dp, err)
			return
		}
		uoHis := state.ToUserOperationHis(dp.TxHash, &state.UserOperation{
			Uid:            ai.Uid,
			Did:            dp.Did,
			OperationType:  state.DepositAction,
			OperationValue: (*hexutil.Big)(dp.Amount),
			Owner:          dp.User,
			Sender:         dp.Account,
			Exec: state.ExecData{
				ChainId: hexutil.Uint64(dp.ChainID),
			},
			Status: state.SuccessStatus,
		})
		err = s.state.GetHisIns().SaveAccountHis(&uoHis)
		if err != nil {
			log.Errorf("add account(%s, %s) history error: %v", dp.User, dp.Account, err)
		}
	}
	vdpFunc := func(dp etherman.DepositData) {
		log.Infof("sync to a new vizing deposit ticket, data: %+v", dp)
		ticker, err := s.state.GetSignedUserOp(dp.User, dp.Account, dp.Did, state.NormalStatus)
		if err != nil {
			log.Errorf("get deposit operation, err: %+v", err)
			return
		}
		if ticker == nil {
			log.Errorf("invalid deposit operation, data: %+v", dp)
			return
		}
		if ticker.OperationValue.Uint64() != dp.Amount.Uint64() {
			log.Errorf("invalid deposit signedUserOperation, amount mismatch, data: %+v", dp)
			return
		}
		dbTx, err := s.db.BeginDBTransaction(s.ctx)
		if err != nil {
			log.Errorf("use db transaction err: %+v", err)
			return
		}
		if err := s.state.AddAccountGas(ticker, dbTx); err != nil {
			_ = dbTx.Rollback(s.ctx)
			log.Errorf("deposit signedUserOperation, add account gas error: %+v", err)
			return
		}
		if err := s.state.UpdateUserOpStatus(ticker.OpId, state.PendingStatus, dbTx); err != nil {
			_ = dbTx.Rollback(s.ctx)
			log.Errorf("update deposit signedUserOperation status error: %+v", err)
			return
		}
		uoHis := state.ToUserOperationHis(dp.TxHash, &state.UserOperation{
			Uid:            ticker.Uid,
			Did:            dp.Did,
			OperationType:  state.DepositAction,
			OperationValue: (*hexutil.Big)(dp.Amount),
			Owner:          dp.User,
			Sender:         dp.Account,
			Exec: state.ExecData{
				ChainId: hexutil.Uint64(dp.ChainID),
			},
			Status: state.SuccessStatus,
		})
		err = s.state.GetHisIns().SaveAccountHis(&uoHis)
		if err != nil {
			_ = dbTx.Rollback(s.ctx)
			log.Errorf("add account(%s, %s) history error: %v", dp.User, dp.Account, err)
			return
		}
		err = dbTx.Commit(s.ctx)
		if err != nil {
			log.Errorf("db commit error: %+v", err)
			return
		}
	}
	wtFunc := func(wt etherman.WithdrawData) {
		log.Info("sync to a new withdraw ticket，data: %+v", wt)
	}
	eoFunc := func(eo etherman.ExecOpData) {
		if eo.Phase != state.PhaseSecond || !eo.InnerExec || !eo.Success {
			return
		}
		log.Info("sync to a new exec op evnet，data: %+v", eo)
		uop, err := s.state.GetSignedUserOp(eo.Owner, eo.Account, eo.ID, state.SuccessStatus)
		if err != nil {
			log.Errorf("get signedUserOperation by id, error: %+v", err)
			return
		}
		if uop.InnerExec.ChainId < 1 || uop.InnerExec.CallData.String() == "0x" {
			err = s.state.UpdateUserOpStatus(uop.OpId, state.FailedStatus, nil)
			log.Warnf("no inner exec data, %+v, err: %+v", uop, err)
			return
		}
		dbTx, err := s.db.BeginDBTransaction(s.ctx)
		if err != nil {
			log.Errorf("use db transaction err: %+v", err)
			return
		}
		err = s.state.UpdateUserOpStatus(uop.OpId, state.PendingStatus, dbTx)
		if err != nil {
			_ = dbTx.Rollback(s.ctx)
			log.Errorf("update op status err: %+v", err)
			return
		}
		err = s.state.UpdateUserOpPhase(uop.OpId, state.PhaseSecond, dbTx)
		if err != nil {
			_ = dbTx.Rollback(s.ctx)
			log.Errorf("update signedUserOperation phase error: %+v", err)
			return
		}
		_ = dbTx.Commit(s.ctx)
		s.pool.CheckFlush(true)
	}
	for _, cli := range s.ether.GetChains() {
		if !cli.IsNeedSync() {
			continue
		}
		chainSync, err := etherman.NewSynchronizer(s.ctx, s.db.GetPool(), cli, acFunc, vdpFunc, dpFunc, wtFunc, eoFunc)
		if err != nil {
			log.Fatal(err.Error())
		}
		go chainSync.Sync()
	}
}

func (s *Synchronizer) Stop() {
	log.Info("sync stop")
	s.cancelCtx()
}
