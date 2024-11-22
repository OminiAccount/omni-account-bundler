package synchronizer

import (
	"context"
	stateTypes "github.com/OAB/state/types"
	"github.com/OAB/synchronizer/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"time"
)

type Synchronizer struct {
	ether types.EthereumInterface
	state types.StateInterface
	cfg   Config

	logger    log.Logger
	ctx       context.Context
	cancelCtx context.CancelFunc
}

func NewSynchronizer(
	ethereum types.EthereumInterface,
	state types.StateInterface,
	cfg Config) (*Synchronizer, error) {
	ctx, cancel := context.WithCancel(context.Background())
	return &Synchronizer{
		ether:     ethereum,
		state:     state,
		cfg:       cfg,
		logger:    log.New("service", "synchronizer"),
		ctx:       ctx,
		cancelCtx: cancel,
	}, nil
}

func (s *Synchronizer) Start() {
	s.sync()
}

func (s *Synchronizer) Stop() {
	s.logger.Info("Sync stop")
	s.cancelCtx()
}

func (s *Synchronizer) sync() {
	s.logger.Info("Sync start")
	//go s.syncTickets()
	go s.syncAccountCreated()
}

//func (s *Synchronizer) syncTickets() {
//	s.logger.Info("Components 1/2", "component", "tickets")
//	// start all chains tickets sync
//	var chans []<-chan pool.TicketFull
//
//	for _, network := range s.cfg.EthereumCfg.Networks {
//		ch := make(chan pool.TicketFull)
//		chans = append(chans, ch)
//		go func(chainId chains.ChainId, ch chan pool.TicketFull) {
//			if err := s.ether.WatchEntryPointEvent(s.ctx, chainId, 0, ch); err != nil {
//				s.logger.Error("Failed to start event listener", "chainId", chainId, "error", err)
//			}
//		}(chains.ChainId(network.ChainId), ch)
//	}
//
//	// mock
//	go func() {
//		ch := make(chan pool.TicketFull)
//		chans = append(chans, ch)
//		time.Sleep(3 * time.Second)
//		insertTicket(ch)
//	}()
//	time.Sleep(1 * time.Second)
//
//	ticketChannel := mergeChannels(s.ctx, chans...)
//
//	for {
//		select {
//		case ticket, ok := <-ticketChannel:
//			if !ok {
//				return
//			}
//			s.logger.Info("Synchronize to a new ticket")
//			// check
//			//err := s.state.AddTicket(ticket)
//			//if err != nil {
//			//	s.logger.Warn("Failed to synchronize to a new ticket", "error", err)
//			//}
//			s.pool.AddTicket(ticket)
//		case <-s.ctx.Done():
//			s.logger.Warn("Stopping Sync due to context cancellation")
//			return
//		default:
//		}
//	}
//}

func (s *Synchronizer) syncAccountCreated() {
	s.logger.Info("Components 2/2", "component", "accountCreated")
	ch := make(chan stateTypes.AccountMapping)

	go func(ch chan stateTypes.AccountMapping) {
		if err := s.ether.WatchAAFactoryEvent(s.ctx, 0, ch); err != nil {
			s.logger.Error("Failed to start aa factory listener", "error", err)
		}
	}(ch)

	go func() {
		// get all events
		time.Sleep(3 * time.Second)
		mappingInsert := stateTypes.AccountMapping{ //01b7cA9d6B8Ac943185E107e4BE7430e5D90B5A5
			User:    common.HexToAddress("a0Ee7A142d267C1f36714E4a8F75612F20a79720"),
			Account: common.HexToAddress("BF9bb0ED00C1d6a1D95A4cbcac9897eb33247580"),
		}
		ch <- mappingInsert
	}()

	for {
		select {
		case mapping, ok := <-ch:
			if !ok {
				return
			}
			s.logger.Info("Synchronize to a new account mapping", "user", mapping.User, "account", mapping.Account)
			err := s.state.AddNewMapping(mapping)
			if err != nil {
				s.logger.Warn("Add a new account mapping error", "error", err)
			}
		case <-s.ctx.Done():
			s.logger.Warn("Stopping Sync due to context cancellation")
			return
		default:
		}
	}
}

//func mergeChannels(ctx context.Context, chans ...<-chan pool.TicketFull) <-chan pool.TicketFull {
//	out := make(chan pool.TicketFull)
//	var wg sync.WaitGroup
//	wg.Add(len(chans))
//
//	for _, ch := range chans {
//		go func(c <-chan pool.TicketFull) {
//			defer wg.Done()
//			for {
//				select {
//				case ticket, ok := <-c:
//					if !ok {
//						return
//					}
//					select {
//					case out <- ticket:
//					case <-ctx.Done():
//						return
//					}
//				case <-ctx.Done():
//					return
//				}
//			}
//		}(ch)
//	}
//
//	go func() {
//		wg.Wait()
//		close(out)
//	}()
//
//	return out
//}
