package synchronizer

import (
	"context"
	"github.com/OAB/etherman"
	"github.com/OAB/pool"
	stateTypes "github.com/OAB/state/types"
	"github.com/OAB/synchronizer/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
)

type Synchronizer struct {
	pool      types.PoolInterface
	ether     *etherman.EtherMan
	state     types.StateInterface
	logger    log.Logger
	ctx       context.Context
	cancelCtx context.CancelFunc
	storage   ethdb.Database
}

func NewSynchronizer(p types.PoolInterface, ethereum *etherman.EtherMan, state types.StateInterface, db ethdb.Database) (*Synchronizer, error) {
	ctx, cancel := context.WithCancel(context.Background())
	return &Synchronizer{
		pool:      p,
		ether:     ethereum,
		state:     state,
		logger:    log.New("service", "synchronizer"),
		ctx:       ctx,
		cancelCtx: cancel,
		storage:   db,
	}, nil
}

func (s *Synchronizer) Start() {
	s.logger.Info("Sync start")
	acFunc := func(acc etherman.AccountCreateData) {
		s.logger.Info("Synchronize to a new account mapping", "user", acc.Owner, "account", acc.Account)
		err := s.state.AddNewMapping(stateTypes.AccountMapping{
			User:    acc.Owner,
			Account: acc.Account,
		})
		if err != nil {
			s.logger.Warn("Add a new account mapping error", "error", err)
		}
	}
	dpFunc := func(dp etherman.DepositData) {
		s.logger.Info("Synchronize to a new deposit ticket")
		s.pool.AddTicket(&pool.TicketFull{
			Ticket: pool.Ticket{
				User:   dp.Account,
				Amount: dp.Amount,
			},
			Type: pool.Deposit,
		})
	}
	wtFunc := func(wt etherman.WithdrawData) {
		s.logger.Info("Synchronize to a new withdraw ticket")
		// TODO check withdraw
		/*s.pool.AddTicket(&pool.TicketFull{
			Ticket: pool.Ticket{
				User:   wt.Account,
				Amount: wt.Amount,
			},
			Type: pool.Withdraw,
		})*/
	}
	for _, cli := range s.ether.ChainsClient {
		chainSync, err := etherman.NewSynchronizer(s.storage, cli, 0, acFunc, dpFunc, wtFunc)
		if err != nil {
			panic(err.Error())
		}
		go chainSync.Sync()
	}
}

func (s *Synchronizer) Stop() {
	s.logger.Info("Sync stop")
	s.cancelCtx()
}
