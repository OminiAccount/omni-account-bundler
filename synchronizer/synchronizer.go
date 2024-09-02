package synchronizer

import (
	"context"
	"github.com/OAAC/pool"
	"github.com/OAAC/synchronizer/types"
	"github.com/OAAC/utils/chains"
	"github.com/ethereum/go-ethereum/log"
	"sync"
)

type Synchronizer struct {
	pool  types.PoolInterface
	ether types.EthereumInterface
	cfg   Config

	logger    log.Logger
	ctx       context.Context
	cancelCtx context.CancelFunc
}

func NewSynchronizer(pool types.PoolInterface,
	ethereum types.EthereumInterface,
	cfg Config) (*Synchronizer, error) {
	ctx, cancel := context.WithCancel(context.Background())
	return &Synchronizer{
		pool:      pool,
		ether:     ethereum,
		cfg:       cfg,
		logger:    log.New("service", "synchronizer"),
		ctx:       ctx,
		cancelCtx: cancel,
	}, nil
}

func (s *Synchronizer) Sync() {
	s.logger.Info("Sync start")
	// start all chains tickets sync
	var chans []<-chan pool.Ticket

	for _, network := range s.cfg.EthereumCfg.Networks {
		ch := make(chan pool.Ticket)
		chans = append(chans, ch)
		go func(chainId chains.ChainId, ch chan pool.Ticket) {
			if err := s.ether.WatchEntryPointEvent(s.ctx, chainId, 0, ch); err != nil {
				s.logger.Error("Failed to start event listener", "chainId", chainId, "error", err)
			}
		}(chains.ChainId(network.ChainId), ch)
	}

	//go func() {
	//	ch := make(chan pool.Ticket)
	//	chans = append(chans, ch)
	//	time.Sleep(3 * time.Second)
	//	insertTicket(ch)
	//}()
	//time.Sleep(1 * time.Second)

	ticketChannel := mergeChannels(s.ctx, chans...)

	for {
		select {
		case ticket, ok := <-ticketChannel:
			if !ok {
				return
			}
			s.logger.Info("Synchronize to a new ticket")
			s.pool.AddTicket(ticket)
		case <-s.ctx.Done():
			return
		}
	}
}

func (s *Synchronizer) Stop() {
	s.logger.Info("Sync stop")
	s.cancelCtx()
}

func mergeChannels(ctx context.Context, chans ...<-chan pool.Ticket) <-chan pool.Ticket {
	out := make(chan pool.Ticket)
	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, ch := range chans {
		go func(c <-chan pool.Ticket) {
			defer wg.Done()
			for {
				select {
				case ticket, ok := <-c:
					if !ok {
						return
					}
					select {
					case out <- ticket:
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
