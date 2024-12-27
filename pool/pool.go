package pool

import (
	"context"
	"errors"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/state"
	"github.com/OAB/utils/log"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
	"time"
)

type ProofResult struct {
	Number       uint64        `json:"number"`
	FinalNumber  uint64        `json:"finalNumber"`
	Proof        hexutil.Bytes `json:"proof"`
	PublicValues hexutil.Bytes `json:"public_values"`
}

type Pool struct {
	mu            sync.Mutex
	cfg           Config
	lastFlushTime time.Time
	ctx           context.Context
	cancelCtx     context.CancelFunc
	state         StateInterface
	ethereum      EthereumInterface
	db            *PostgresStorage
}

func NewMemoryPool(cfg Config, ether EthereumInterface, s StateInterface, pg *pgxpool.Pool) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	pool := &Pool{
		ctx:           ctx,
		cancelCtx:     cancel,
		cfg:           cfg,
		lastFlushTime: time.Now(),
		db:            NewPostgresStorage(pg),
		state:         s,
		ethereum:      ether,
	}
	return pool
}

func (p *Pool) StartAutoFlush() {
	log.Info("pool start")
	go func() {
		ticker := time.NewTicker(time.Minute * 3)
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
	go func() {
		for {
			if p.ctx.Err() != nil {
				return
			}
			err := p.verifyProof()
			if err != nil {
				if !errors.Is(err, pgx.ErrNoRows) {
					log.Errorf("send batch to verify error: %s", err.Error())
				}
				time.Sleep(time.Minute * 10) // 1min
				continue
			}
		}
	}()
}

func (p *Pool) StopAutoFlush() {
	log.Info("pool stop auto flush")
	p.cancelCtx()
}

func (p *Pool) CheckFlush(flag bool) {
	count, err := p.db.GetPendingOpCount(p.ctx, nil)
	if err != nil {
		log.Errorf("GetPendingOpCount err: %+v", err)
		return
	}
	if count < 1 {
		return
	}
	if flag {
		p.processBatch()
	} else if count >= p.cfg.MaxOps {
		p.processBatch()
		p.CheckFlush(false)
	} else if p.cfg.FlushInterval > 0 && time.Since(p.lastFlushTime).Seconds() >= float64(p.cfg.FlushInterval) {
		p.processBatch()
	}
}

func (p *Pool) GetBatchProof() ([]*Batch, error) {
	vb, err := p.db.GetBatchNoOp(p.ctx, state.SuccessStatus, nil)
	if err != nil {
		return nil, err
	}
	return p.db.GetBatch(p.ctx, state.PendingStatus, vb.NewNumBatch+1, vb.NewNumBatch+p.cfg.MaxBatches, nil)
}

func (p *Pool) SetBatchProofResult(res *ProofResult) error {
	log.Infof("SetBatchProofResult: %+v", res)
	dbTx, err := p.db.BeginDBTransaction(p.ctx)
	if err != nil {
		log.Errorf("use db transaction err: %+v", err)
		return err
	}
	err = p.db.AddProof(p.ctx, res, dbTx)
	if err != nil {
		_ = dbTx.Rollback(p.ctx)
		log.Errorf("add proof err: %+v", err)
		return err
	}
	for i := res.Number; i <= res.FinalNumber; i++ {
		err = p.db.UpdateBatch(p.ctx, i, "status", state.BatchStatus, dbTx)
		if err != nil {
			_ = dbTx.Rollback(p.ctx)
			log.Errorf("update batch status err: %+v", err)
			return err
		}
	}
	err = dbTx.Commit(p.ctx)
	return err
}
