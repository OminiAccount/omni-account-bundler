package pool

import (
	"context"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
	"time"
)

type Pool struct {
	mu            sync.Mutex
	cfg           Config
	lastFlushTime time.Time
	batchCtx      chan *BatchContext
	ctx           context.Context
	cancelCtx     context.CancelFunc
	db            *PostgresStorage
}

func NewMemoryPool(cfg Config, db ethdb.Database, pg *pgxpool.Pool) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	pool := &Pool{
		ctx:           ctx,
		cancelCtx:     cancel,
		cfg:           cfg,
		lastFlushTime: time.Now(),
		batchCtx:      make(chan *BatchContext, 100),
		db:            NewPostgresStorage(pg),
	}
	return pool
}

func (p *Pool) CheckFlush(flag bool) {
	// If the flushInterval is not 0, check both maxOps and time interval
	/*if flag {
		p.Flush()
	} else if p.cfg.FlushInterval != 0 {
		if uint64(len(p.storage.UserOps)) >= p.cfg.MaxOps || time.Since(p.lastFlushTime).Seconds() >= float64(p.cfg.FlushInterval) {
			p.Flush()
		}
	} else if uint64(len(p.storage.UserOps)) >= p.cfg.MaxOps {
		p.Flush()
	}*/
}

func (p *Pool) Flush() {
	/*p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.storage.getUserOps()) < 1 {
		return
	}
	log.Debug("flushing memory pool...")
	batchCtx := NewBatchContext(p.storage.getUserOps())

	// Empty the memory pool
	p.storage.empty()
	p.lastFlushTime = time.Now()

	// Execute specific processing logic
	p.batchCtx <- batchCtx*/
}

func (p *Pool) StartAutoFlush() {
	log.Info("pool start")
	go func() {
		ticker := time.NewTicker(time.Minute)
		for {
			select {
			case <-ticker.C:
				p.CheckFlush(false)
			case <-p.ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}

func (p *Pool) StopAutoFlush() {
	log.Info("pool stop auto flush")
	p.cancelCtx()
}

func (p *Pool) BatchContext() chan *BatchContext {
	return p.batchCtx
}
