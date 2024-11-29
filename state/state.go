package state

import (
	"context"
	"fmt"
	"github.com/OAB/pool"
	"github.com/OAB/state/types"
	"github.com/OAB/utils/db"
	"github.com/OAB/utils/merkletree"
	"github.com/OAB/utils/queue"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"time"
)

type State struct {
	cfg    Config
	ctx    context.Context
	cancel func()

	pool     PoolInterface
	ethereum EthereumInterface
	storage  *Storage
	logger   log.Logger
	tree     *merkletree.SMT

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
		tree:        merkletree.NewSMT(db.NewMemDb(cfg.LevelDB),false),
		proofQueue:  queue.NewConcurrentQueue[Batch](),
		provenQueue: queue.NewConcurrentQueue[ProofResult](),
		storage:     NewStorage(cfg.LevelDB),

		logger: log.New("service", "state"),
	}

	state.LoadCache()

	return state, nil
}

func (s *State) Start() {
	go s.ProcessBatch()
	go s.ExecuteBatch()
	go func() {
		for {
			time.Sleep(time.Minute)
			s.Persistence()
		}
	}()
}

func (s *State) Stop() {
	s.Persistence()
	s.pool.Cache()
	s.cancel()
}

func (s *State) ProcessBatch() {
	s.logger.Info("ProcessBatch start")
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case batchContext := <-s.pool.Context():
			s.logger.Debug("Processing batch", "SignedUserOperations number", len(batchContext.SignedUserOperations()))
			err := s.processBatch(batchContext)
			if err != nil {
				s.logger.Warn("failed process a batch", "error", err)
			}
		case <-ticker.C:
		// Todo: Here you can add processing logic, such as periodically checking the status of the pool or other tasks.
		case <-s.ctx.Done():
			s.logger.Warn("Stopping PreProcessBatch due to context cancellation")
			return
		default:
		}
	}
}

func (s *State) processBatch(batchContext *pool.BatchContext) error {
	batch, err := NewBatch(s.storage.BatchNumber+1)
	if err != nil {
		return err
	}

	// Set old smt root
	batch.SetOldStateRoot(common.BigToHash(s.tree.LastRoot()))

	for _, userOperation := range batchContext.SignedUserOperations() {
		fmt.Printf("Address: %s ,Nonce: %d ,ChainId %s ,OperationType: %d ,operationValue: %d ===============\n", userOperation.Owner, userOperation.Nonce, userOperation.ChainId.String(), userOperation.OperationType, userOperation.OperationValue)

		// balance
		var balance big.Int
		balanceU256, err := s.tree.GetAccountBalance(userOperation.Sender)
		if err != nil {
			return err
		}
		balance.SetBytes(balanceU256.Bytes())
		fmt.Println("oldBalance", &balance)
		operationValue := new(big.Int).SetUint64(userOperation.OperationValue.Uint64())

		if userOperation.OperationType == 1 {
			fmt.Printf("balance %d + operationValue %d\n", &balance, operationValue)
			balance.Add(&balance, operationValue)
		} else if userOperation.OperationType == 2 {
			if balance.Cmp(operationValue) < 0 {
				fmt.Printf("current account balance %s is insufficient to withdraw %s", &balance, operationValue)
			}
			fmt.Printf("balance %d - operationValue %d\n", &balance, operationValue)
			balance.Sub(&balance, operationValue)
		}

		fmt.Printf("balance %d - CalculateGasUsed %d\n", &balance, userOperation.CalculateGasUsed())
		balance.Sub(&balance, userOperation.CalculateGasUsed())

		if _, err = s.tree.SetAccountBalance(userOperation.Sender.String(), &balance); err != nil {
			return err
		}
		fmt.Println("newBalance", &balance)

		// nonce
		var nonce big.Int
		nonceU256, err := s.tree.GetAccountNonce(userOperation.Sender, userOperation.ChainId.String())
		if err != nil {
			return err
		}
		nonce.SetBytes(nonceU256.Bytes())
		fmt.Println("oldNonce", nonce.Uint64())
		nonce.Add(&nonce, big.NewInt(1))
		fmt.Println("newNonce", nonce.Uint64())
		_, err = s.tree.SetAccountNonce(userOperation.Sender.String(), &nonce, userOperation.ChainId.String())
		if err != nil {
			return err
		}

		//expectGasBalance, expectNonce, _, err := s.GetAccountInfo(userOperation.Owner, userOperation.Sender, userOperation.ChainId.ToInt().Uint64())
		//if err != nil {
		//	return err
		//}
		//fmt.Println("expectGasBalance", expectGasBalance)
		//fmt.Println("smtGasBalance", balance)
		//fmt.Println("expectNonce", expectNonce)
		//fmt.Println("smtNonce", nonce)
	}

	// Set BatchL2Data and BatchHashData
	batch.SetBatchL2Data(batchContext.SignedUserOperations())

	// Set NewAccInputHash
	// Todo: Now just copy the batchHashData
	batch.SetNewAccInputHash(batch.BatchHashData)

	// Set new state root
	batch.SetNewStateRoot(common.BigToHash(s.tree.LastRoot()))
	s.proofQueue.Enqueue(*batch)
	s.logger.Info("Successfully sealing a batch", "number", batch.NewNumBatch, "stateRoot", batch.NewStateRoot.Hex())

	return nil
}

func (s *State) ExecuteBatch() {
	s.logger.Info("ExecuteBatch start")
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			err := s.executeBatch()
			if err != nil {
				s.logger.Warn("send batch to verify error", "error", err)
			}
		case <-s.ctx.Done():
			s.logger.Warn("Stopping ExecuteBatch due to context cancellation")
			return
		default:
		}
	}
}

func (s *State) executeBatch() error {
	if !s.provenQueue.IsEmpty() {
		s.logger.Info("Executing batch")
		res, _ := s.provenQueue.Dequeue()
		_, err := s.ethereum.UpdateEntryPointRoot(res.Proof, res.PublicValues)
		if err != nil {
			return err
		}
		// Todo: The target chain should be listening for synchronous events
	}

	return nil
}

func (s *State) AddNewMapping(mapping types.AccountMapping) error {
	return s.storage.Account.AddNewMapping(mapping)
}

func (s *State) AddSignedUserOperation(signedUserOp *pool.SignedUserOperation) error {
	// check
	s.logger.Debug("Check SignedUserOperation")
	if err := s.storage.Account.AddSignedUserOperation(signedUserOp); err != nil {
		return fmt.Errorf("check SignedUserOperation failed, error: %s", err.Error())
	}

	s.pool.AddSignedUserOperation(signedUserOp)

	return nil
}

func (s *State) GetAccountsForUser(user common.Address) *[]common.Address {
	return s.storage.Account.GetAccountsForUser(user)
}

func (s *State) GetAccountInfo(user, account common.Address, chainId uint64) (*big.Int, uint64, []*pool.UserOperation, error) {
	accountInfo, err := s.storage.Account.GetAccount(user, account)
	if err != nil {
		return nil, 0, nil, err
	}
	return accountInfo.Gas, accountInfo.Nonce[chainId], accountInfo.UserOperations, nil
}

func (s *State) Persistence() {
	s.logger.Info("Cache StateContext info into disk")
	if err := s.storage.cache(); err != nil {
		s.logger.Error("Cache StateContext info into disk error", "error", err)
	}
	if err := s.tree.Db.Cache(); err != nil {
		s.logger.Error("Cache Smt into disk error", "error", err)
	}
}

func (s *State) LoadCache() {
	s.logger.Info("Load StateContext info from disk")
	if err := s.storage.loadCache(); err != nil {
		s.logger.Error("Load StateContext info from disk error", "error", err)
	}
	err := s.tree.Db.LoadCache()
	if err != nil {
		s.logger.Error("Load Smt info from disk error", "error", err)
	}
}
