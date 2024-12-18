package synchronizer

import (
	"context"
	"github.com/OAB/etherman"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/pool"
	stateTypes "github.com/OAB/state/types"
	"github.com/OAB/synchronizer/types"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/ethdb"
)

type Synchronizer struct {
	pool      types.PoolInterface
	ether     types.EthereumInterface
	state     types.StateInterface
	ctx       context.Context
	cancelCtx context.CancelFunc
	storage   ethdb.Database
}

func NewSynchronizer(p types.PoolInterface, ethereum types.EthereumInterface, state types.StateInterface, db ethdb.Database) (*Synchronizer, error) {
	ctx, cancel := context.WithCancel(context.Background())
	return &Synchronizer{
		pool:      p,
		ether:     ethereum,
		state:     state,
		ctx:       ctx,
		cancelCtx: cancel,
		storage:   db,
	}, nil
}

func (s *Synchronizer) Start() {
	log.Info("Sync start......")
	acFunc := func(acc etherman.AccountCreateData) {
		log.Infof("chainID: %d, sync to a new account, user: %s, account: %s",
			acc.ChainID, acc.Owner, acc.Account)
		s.state.InitAccountNonce(acc.ChainID, stateTypes.AccountMapping{
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
		uoHis := stateTypes.ToUserOperationHis(dp.TxHash, &pool.UserOperation{
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
		}
	}
	vdpFunc := func(dp etherman.DepositData) {
		log.Infof("sync to a new vizing deposit ticket, data: %+v", dp)
		ticker := s.pool.GetTicket(dp.Did)
		if ticker == nil {
			log.Errorf("invalid deposit operation, data: %+v", dp)
			return
		}
		if ticker.User != dp.Account {
			log.Errorf("invalid deposit signedUserOperation, account mismatch, data: %+v", dp)
			return
		}
		if ticker.Amount.Uint64() != dp.Amount.Uint64() {
			log.Errorf("invalid deposit signedUserOperation, amount mismatch, data: %+v", dp)
			return
		}
		if err := s.state.AddAccountGas(ticker.SignedUserOp); err != nil {
			log.Errorf("deposit signedUserOperation, add account gas error: %+v", err)
			return
		}
		s.pool.AddSignedUserOperation(ticker.SignedUserOp)
		uoHis := stateTypes.ToUserOperationHis(dp.TxHash, ticker.SignedUserOp.UserOperation)
		err := s.state.GetHisIns().SaveAccountHis(ticker.SignedUserOp.Owner, ticker.SignedUserOp.Sender, &uoHis)
		if err != nil {
			log.Errorf("cache account(%s, %s) history error: %v",
				ticker.SignedUserOp.Owner, ticker.SignedUserOp.Sender, err)
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
		uop, err := s.state.GetSignedUserOp(eo.Owner, eo.Account, eo.ID)
		if err != nil {
			log.Errorf("get signedUserOperation by id, error: %+v", err)
			return
		}
		if uop.InnerExec.ChainId < 1 || uop.InnerExec.CallData.String() == "0x" {
			log.Warnf("no inner exec data, %+v", uop)
			return
		}
		uop.Phase = pool.PhaseSecond
		s.pool.AddStepUserOperation(uop)
	}
	for _, cli := range s.ether.GetChains() {
		if !cli.IsNeedSync() {
			continue
		}
		chainSync, err := etherman.NewSynchronizer(s.ctx, s.storage, cli, acFunc, vdpFunc, dpFunc, wtFunc, eoFunc)
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
