package state

import (
	"context"
	"github.com/OAB/utils/db"
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/merkletree"
	"github.com/OAB/utils/queue"
	"time"
)

type State struct {
	cfg    Config
	ctx    context.Context
	cancel func()

	pool     PoolInterface
	ethereum EthereumInterface
	storage  *Storage
	tree     *merkletree.SMT
	his      *HistoryManager

	proofQueue  *queue.ConcurrentQueue[Batch]
	provenQueue *queue.ConcurrentQueue[ProofResult]
}

func NewState(cfg Config, pool PoolInterface, ethereum EthereumInterface) (*State, error) {
	ctx, cancel := context.WithCancel(cfg.Context)
	state := &State{
		cfg:         cfg,
		ctx:         ctx,
		cancel:      cancel,
		pool:        pool,
		ethereum:    ethereum,
		tree:        merkletree.NewSMT(db.NewMemDb(cfg.LevelDB), false),
		proofQueue:  queue.NewConcurrentQueue[Batch](),
		provenQueue: queue.NewConcurrentQueue[ProofResult](),
		storage:     NewStorage(cfg.LevelDB),
	}
	state.his = NewHistoryManager(ctx, cfg, state)

	state.LoadCache()

	return state, nil
}

func (s *State) Start() {
	log.Info("state start")
	go s.ProcessBatch()
	go s.ExecuteBatch()
	go s.his.Start()
	go func() {
		for {
			time.Sleep(time.Second*5)
			s.Persistence()
		}
	}()
	/*go func() {
		time.Sleep(time.Second*20)
		// create uop
		for i := 1; i < 3; i++ {
			err := s.AddSignedUserOperation(&pool.SignedUserOperation{
				Owner:     common.HexToAddress("0x24F4F8e83eA08b7DCb7A579B1236fc3155300854"),
				Did:       strconv.FormatInt(int64(i), 10),
				Signature: []byte{},
				UserOperation: &pool.UserOperation{
					OperationType:          pool.DepositAction,
					OperationValue:         2000,
					Sender:                 common.HexToAddress("0x2c6a4142c56b848e01e95fc35c4ad7b6483bc60c"),
					Nonce:                  hexutil.Uint64(i),
					ChainId:                (*hexutil.Big)(big.NewInt(28516)),
					CallData:               hexutil.Bytes{0},
					MainChainGasLimit:      hexutil.Uint64(10),
					DestChainGasLimit:      hexutil.Uint64(10),
					ZkVerificationGasLimit: hexutil.Uint64(10),
					MainChainGasPrice:      (*hexutil.Big)(big.NewInt(2)),
					DestChainGasPrice:      (*hexutil.Big)(big.NewInt(2)),
				},
			})
			log.Infof("%+v", err)
		}

		// create proof
		time.Sleep(time.Minute)
		_ = s.SetBatchProofResult(&ProofResult{
			Number: 1,
		})
	}()*/
}

func (s *State) Stop() {
	log.Info("state stop")
	s.cancel()
	time.Sleep(time.Second * 2)
	s.Persistence()
}

func (s *State) ProcessBatch() {
	log.Info("ProcessBatch start")
	for {
		select {
		case batchContext := <-s.pool.BatchContext():
			log.Debugf("Processing batch, SignedUserOperations count: %d", len(batchContext.SignedUserOperations()))
			err := s.processBatch(batchContext)
			if err != nil {
				log.Error("failed process a batch", "error", err)
			}
		case <-s.ctx.Done():
			if len(s.pool.BatchContext()) > 0 {
				continue
			}
			log.Warn("Stopping PreProcessBatch due to context cancellation")
			return
		}
	}
}

func (s *State) ExecuteBatch() {
	log.Info("ExecuteBatch start")
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			res, err := s.executeBatch()
			if err != nil {
				log.Errorf("send batch to verify error: %s", err.Error())
			}
			if res != nil {
				s.provenQueue.Lpush(*res)
			}
		case <-s.ctx.Done():
			if !s.provenQueue.IsEmpty() {
				continue
			}
			log.Warn("Stopping ExecuteBatch due to context cancellation")
			return
		}
	}
}

func (s *State) GetHisIns() *HistoryManager {
	return s.his
}

func (s *State) Persistence() {
	//log.Info("cache state data into disk")
	if err := s.storage.cache(); err != nil {
		log.Error("cache state data into disk error", "error", err)
	}
	if err := s.tree.Db.Cache(); err != nil {
		log.Error("cache smt into disk error", "error", err)
	}
	if err := s.pool.Cache(); err != nil {
		log.Error("cache pool into disk error", "error", err)
	}
}

func (s *State) LoadCache() {
	log.Info("load state data from disk")
	if err := s.storage.loadCache(); err != nil {
		log.Fatalf("load state data from disk error: %v", err)
	}
	err := s.tree.Db.LoadCache()
	if err != nil {
		log.Fatalf("load smt data from disk error: %v", err)
	}
	for i := s.storage.VerifyBatchNumber + 1; i <= s.storage.BatchNumber; i++ {
		batch, err := s.storage.loadBatch(i)
		if err != nil && err != NotBatch {
			log.Fatalf("load batch info from disk error: %v", err)
		}
		if batch != nil {
			log.Infof("from cache load batch: %d", batch.NewNumBatch)
			s.proofQueue.Enqueue(*batch)
		}
	}
}
