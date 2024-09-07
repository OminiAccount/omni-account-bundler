package pool

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"sync"
	"time"
)

type Pool struct {
	mu            sync.Mutex
	cfg           Config
	lastFlushTime time.Time
	context       chan BatchContext
	stopChan      chan struct{}

	storage *Storage

	logger log.Logger
}

func NewMemoryPool(cfg Config, db ethdb.Database) *Pool {
	pool := &Pool{
		cfg:           cfg,
		lastFlushTime: time.Now(),
		context:       make(chan BatchContext, 100),
		storage:       NewStorage(db),
		logger:        log.New("service", "pool"),
	}

	pool.LoadCache()

	// mock
	go func() {
		time.Sleep(10 * time.Second)
		pool.mockPool()
	}()

	go pool.StartAutoFlush()

	return pool
}

func (p *Pool) AddUserOp(op SignedUserOperation) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.logger.Info("Add a new userOp", "userOp", op)

	p.storage.addUserOp(op)
	p.CheckFlush()
}

func (p *Pool) AddTicket(ticket TicketFull) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.logger.Info("Add a new ticket", "ticket", ticket)

	p.storage.addTicket(ticket)
	p.CheckFlush()
}

func (p *Pool) CheckFlush() {
	// If the flushInterval is not 0, check both maxOps and time interval
	if p.cfg.flushInterval != 0 {
		if uint64(len(p.storage.userOps)) >= p.cfg.maxOps || time.Since(p.lastFlushTime).Seconds() >= float64(p.cfg.flushInterval) {
			p.Flush()
		}
	} else {
		if uint64(len(p.storage.userOps)) >= p.cfg.maxOps {
			p.Flush()
		}
	}
}

func (p *Pool) Flush() {
	fmt.Println("Flushing memory pool...")

	context := BatchContext{
		userOps: p.storage.getUserOps(),
		tickets: p.storage.getTickets(),
	}

	// Empty the memory pool
	p.storage.empty()
	p.lastFlushTime = time.Now()

	// Execute specific processing logic
	fmt.Println("UserOps:", len(context.userOps))
	fmt.Println("Tickets:", len(context.tickets))
	p.context <- context
}

func (p *Pool) StartAutoFlush() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			p.CheckFlush()
		case <-p.stopChan:
			ticker.Stop()
			return
		}
	}
}

func (p *Pool) StopAutoFlush() {
	close(p.stopChan)
}

func (p *Pool) Context() chan BatchContext {
	return p.context
}

func (p *Pool) Cache() {
	if err := p.storage.cacheTickets(); err != nil {
		p.logger.Error("pool cache tickets error", "error", err)
	}
	if err := p.storage.cacheUserOps(); err != nil {
		p.logger.Error("pool cache userOps error", "error", err)
	}
}

func (p *Pool) LoadCache() {
	p.logger.Info("Load cache")
	if err := p.storage.loadTickets(); err != nil {
		p.logger.Error("pool load cache tickets error", "error", err)
	}
	if err := p.storage.loadUserOps(); err != nil {
		p.logger.Error("pool load cache userOps error", "error", err)
	}
}
