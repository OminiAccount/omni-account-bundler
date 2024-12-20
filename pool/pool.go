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

	storage *Storage
	db      *PostgresStorage
}

func NewMemoryPool(cfg Config, db ethdb.Database, pg *pgxpool.Pool) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	pool := &Pool{
		ctx:           ctx,
		cancelCtx:     cancel,
		cfg:           cfg,
		lastFlushTime: time.Now(),
		batchCtx:      make(chan *BatchContext, 100),
		storage:       NewStorage(db),
		db:            NewPostgresStorage(pg),
	}

	pool.LoadCache()

	return pool
}

func (p *Pool) CheckFlush(flag bool) {
	// If the flushInterval is not 0, check both maxOps and time interval
	if flag {
		p.Flush()
	} else if p.cfg.FlushInterval != 0 {
		if uint64(len(p.storage.UserOps)) >= p.cfg.MaxOps || time.Since(p.lastFlushTime).Seconds() >= float64(p.cfg.FlushInterval) {
			p.Flush()
		}
	} else if uint64(len(p.storage.UserOps)) >= p.cfg.MaxOps {
		p.Flush()
	}
}

func (p *Pool) Flush() {
	p.mu.Lock()
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
	p.batchCtx <- batchCtx
}

func (p *Pool) CleanTicker() {
	log.Infof("clean ticket...")
	p.mu.Lock()
	defer p.mu.Unlock()
	nowAt := time.Now().Unix()
	for id, t := range p.storage.Tickets {
		if (nowAt - t.TimeAt) > 3600 {
			delete(p.storage.Tickets, id)
		}
	}
}

func (p *Pool) StartAutoFlush() {
	log.Info("pool start")
	go func() {
		ticker := time.NewTicker(time.Minute)
		for {
			select {
			case <-ticker.C:
				p.CheckFlush(false)
				p.CleanTicker()
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

func (p *Pool) Cache() error {
	if err := p.storage.cacheUserOps(); err != nil {
		log.Errorf("pool cache userOps error: %+v", err)
		return err
	}
	return nil
}

func (p *Pool) LoadCache() {
	log.Info("load pool cache")
	if err := p.storage.loadUserOps(); err != nil {
		log.Error("pool load cache userOps error", "error", err)
	}
}
