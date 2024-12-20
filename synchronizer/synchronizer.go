package synchronizer

import (
	"context"
	"github.com/OAB/etherman"
	"github.com/OAB/pool"
	stateTypes "github.com/OAB/state/types"
	"github.com/OAB/synchronizer/types"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Synchronizer struct {
	ether     types.EthereumInterface
	state     types.StateInterface
	ctx       context.Context
	cancelCtx context.CancelFunc
	storage   ethdb.Database
	db        *PostgresStorage
}

func NewSynchronizer(ctx context.Context, ethereum types.EthereumInterface,
	state types.StateInterface, db ethdb.Database, pg *pgxpool.Pool) (*Synchronizer, error) {
	syncCtx, cancel := context.WithCancel(ctx)
	return &Synchronizer{
		ether:     ethereum,
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
		s.state.CreateAccountInfo(acc.ChainID, stateTypes.AccountMapping{
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
		/*uoHis := stateTypes.ToUserOperationHis(dp.TxHash, &pool.UserOperation{
			Did:            dp.Did,
			OperationType:  pool.DepositAction,
			OperationValue: (*hexutil.Big)(dp.Amount),
			Owner:          dp.User,
			Sender:         dp.Account,
			Exec: pool.ExecData{
				ChainId: hexutil.Uint64(dp.ChainID),
			},
		})
		err := s.state.GetHisIns().SaveAccountHis(dp.User, dp.Account, &uoHis)
		if err != nil {
			log.Errorf("cache account(%s, %s) history error: %v", dp.User, dp.Account, err)
		}*/
	}
	vdpFunc := func(dp etherman.DepositData) {
		log.Infof("sync to a new vizing deposit ticket, data: %+v", dp)
		ticker, err := s.state.GetSignedUserOp(dp.User, dp.Account, dp.Did, pool.NormalStatus)
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
		if err := s.state.UpdateUserOpStatus(ticker.OpId, pool.PendingStatus, dbTx); err != nil {
			_ = dbTx.Rollback(s.ctx)
			log.Errorf("update deposit signedUserOperation status error: %+v", err)
			return
		}

		// TODO add history
		/*uoHis := stateTypes.ToUserOperationHis(dp.TxHash, ticker.SignedUserOp.UserOperation)
		err := s.state.GetHisIns().SaveAccountHis(ticker.SignedUserOp.Owner, ticker.SignedUserOp.Sender, &uoHis)
		if err != nil {
			log.Errorf("cache account(%s, %s) history error: %v",
				ticker.SignedUserOp.Owner, ticker.SignedUserOp.Sender, err)
		}*/

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
		if eo.Phase != pool.PhaseSecond || !eo.InnerExec || !eo.Success {
			return
		}
		log.Info("sync to a new exec op evnet，data: %+v", eo)
		uop, err := s.state.GetSignedUserOp(eo.Owner, eo.Account, eo.ID, pool.SuccessStatus)
		if err != nil {
			log.Errorf("get signedUserOperation by id, error: %+v", err)
			return
		}
		if uop.InnerExec.ChainId < 1 || uop.InnerExec.CallData.String() == "0x" {
			err = s.state.UpdateUserOpStatus(uop.OpId, pool.FailedStatus, nil)
			log.Warnf("no inner exec data, %+v, err: %+v", uop, err)
			return
		}
		dbTx, err := s.db.BeginDBTransaction(s.ctx)
		if err != nil {
			log.Errorf("use db transaction err: %+v", err)
			return
		}
		err = s.state.UpdateUserOpStatus(uop.OpId, pool.PendingStatus, dbTx)
		if err != nil {
			_ = dbTx.Rollback(s.ctx)
			log.Errorf("update op status err: %+v", err)
			return
		}
		err = s.state.UpdateUserOpPhase(uop.OpId, pool.PhaseSecond, dbTx)
		if err != nil {
			_ = dbTx.Rollback(s.ctx)
			log.Errorf("update signedUserOperation phase error: %+v", err)
			return
		}
		_ = dbTx.Commit(s.ctx)
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
