package pool

import (
	"fmt"
	"sync"
	"time"
)

type Pool struct {
	mu            sync.Mutex
	userOps       []SignedUserOperation
	tickets       []TicketFull
	cfg           Config
	lastFlushTime time.Time
	context       chan BatchContext
}

func NewMemoryPool(cfg Config) *Pool {
	pool := &Pool{
		userOps:       []SignedUserOperation{},
		tickets:       []TicketFull{},
		cfg:           cfg,
		lastFlushTime: time.Now(),
		context:       make(chan BatchContext, 100),
	}

	// mock
	pool.mockPool()

	return pool
}

func (p *Pool) AddUserOp(op SignedUserOperation) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.userOps = append(p.userOps, op)
	p.CheckFlush()
}

func (p *Pool) AddTicket(ticket TicketFull) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.tickets = append(p.tickets, ticket)
	p.CheckFlush()
}

func (p *Pool) CheckFlush() {
	if uint64(len(p.userOps)) >= p.cfg.maxOps || time.Since(p.lastFlushTime).Seconds() >= float64(p.cfg.flushInterval) {
		p.Flush()
	}
}

func (p *Pool) Flush() {
	fmt.Println("Flushing memory pool...")

	context := BatchContext{
		userOps: p.userOps,
		tickets: p.tickets,
	}

	// Empty the memory pool
	p.userOps = []SignedUserOperation{}
	p.tickets = []TicketFull{}
	p.lastFlushTime = time.Now()

	// Execute specific processing logic
	fmt.Println("UserOps:", context.userOps)
	fmt.Println("Tickets:", context.tickets)
	p.context <- context
}

func (p *Pool) Context() chan BatchContext {
	return p.context
}
