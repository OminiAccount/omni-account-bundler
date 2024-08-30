package pool

import (
	"fmt"
	"sync"
	"time"
)

type Pool struct {
	mu            sync.Mutex
	userOps       []SignedUserOperation
	tickets       []Ticket
	cfg           Config
	lastFlushTime time.Time
}

func NewMemoryPool(cfg Config) *Pool {
	return &Pool{
		userOps:       []SignedUserOperation{},
		tickets:       []Ticket{},
		cfg:           cfg,
		lastFlushTime: time.Now(),
	}
}

func (mp *Pool) AddUserOp(op SignedUserOperation) {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	mp.userOps = append(mp.userOps, op)
	mp.checkFlush()
}

func (mp *Pool) AddTicket(ticket Ticket) {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	mp.tickets = append(mp.tickets, ticket)
	mp.checkFlush()
}

func (mp *Pool) checkFlush() {
	if uint64(len(mp.userOps)) >= mp.cfg.maxOps || time.Since(mp.lastFlushTime) >= mp.cfg.flushInterval {
		mp.Flush()
	}
}

func (mp *Pool) Flush() {
	fmt.Println("Flushing memory pool...")

	userOps := mp.userOps
	tickets := mp.tickets

	// Empty the memory pool
	mp.userOps = []SignedUserOperation{}
	mp.tickets = []Ticket{}
	mp.lastFlushTime = time.Now()

	// Execute specific processing logic
	fmt.Println("UserOps:", userOps)
	fmt.Println("Tickets:", tickets)
}
func (mp *Pool) StartAutoFlush() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			mp.mu.Lock()
			mp.checkFlush()
			mp.mu.Unlock()
		}
	}()
}
