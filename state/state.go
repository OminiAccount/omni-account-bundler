package state

import (
	"context"
	"fmt"
	"github.com/OAB/pool"
	"github.com/OAB/state/types"
	"github.com/OAB/utils/queue"
	"github.com/OAB/utils/smt"
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

	tree        *smt.ZeroMerkleTree
	proofQueue  *queue.ConcurrentQueue[Batch]
	provenQueue *queue.ConcurrentQueue[ProofResult]

	storage *Storage

	logger log.Logger
}

func NewState(cfg Config, pool PoolInterface, ethereum EthereumInterface) (*State, error) {
	ctx, cancel := context.WithCancel(cfg.Context)
	state := &State{
		cfg:         cfg,
		ctx:         ctx,
		cancel:      cancel,
		pool:        pool,
		ethereum:    ethereum,
		tree:        smt.NewZeroMerkleTree(256),
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
			s.logger.Debug("Processing batch", "batchContext", batchContext)
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

func (s *State) processBatch(batchContext pool.BatchContext) error {
	batch, err := NewBatch(s.storage.BatchNumber)
	if err != nil {
		return err
	}

	// Set old smt root
	batch.SetOldSMTRoot(s.tree.GetRoot())
	var stateDiff *smt.DeltaMerkleProof

	updateDiff := func(diff *smt.DeltaMerkleProof) {
		if stateDiff == nil {
			stateDiff = diff
		}
	}

	// Set tickets
	var depositTickets, withdrawTickets []TicketInput

	for _, ticketFull := range batchContext.SortedTickets() {
		fmt.Println("old root", s.tree.GetRoot())
		var balance big.Int
		balanceKey := smt.ComputeBalanceKey(ticketFull.User.Bytes())
		balanceKeyIndex := smt.KeyToIndex(balanceKey)
		balanceMerkleProof := s.tree.GetLeaf(balanceKeyIndex)
		balance.SetString(string(balanceMerkleProof.Value), 16)
		fmt.Println("Ticket oldBalance", balance)

		actionAmount := ticketFull.Amount.ToInt()

		operation := func() error {
			switch ticketFull.Type {
			case pool.Deposit:
				balance.Add(&balance, actionAmount)
			case pool.Withdraw:
				if balance.Cmp(actionAmount) < 0 {
					return fmt.Errorf("current account balance %s is insufficient to withdraw %s", &balance, actionAmount)
				}
				balance.Sub(&balance, actionAmount)
			default:
				return fmt.Errorf("unsupported ticket type: %v", ticketFull.Type)
			}
			return nil
		}

		if err := operation(); err != nil {
			return err
		}

		// Set newBalance
		delta := s.tree.SetLeaf(balanceKeyIndex, smt.MerkleNodeValue(common.Bytes2Hex(common.LeftPadBytes(balance.Bytes(), 32))))
		fmt.Println("new root", s.tree.GetRoot())
		ticketProof := TicketInput{
			Ticket:      &ticketFull.Ticket,
			TicketProof: &delta,
		}

		switch ticketFull.Type {
		case pool.Deposit:
			depositTickets = append(depositTickets, ticketProof)
		case pool.Withdraw:
			withdrawTickets = append(withdrawTickets, ticketProof)
		}

		fmt.Println("Ticket newBalance", balance)
		updateDiff(&delta)
	}

	batch.SetDepositTickets(depositTickets)
	batch.SetWithdrawTickets(withdrawTickets)

	// Set user operations
	var userOperationProof []UserOperationInput

	for _, userOperation := range batchContext.SignedUserOperations() {
		err = s.storage.Account.AddUserOperation(*userOperation.UserOperation)
		if err != nil {
			return err
		}
		fmt.Println("Nonce========", userOperation.Nonce)
		fmt.Println("old root", s.tree.GetRoot())
		// balance
		var balance big.Int
		balanceKey := smt.ComputeBalanceKey(userOperation.Sender.Bytes())
		balanceKeyIndex := smt.KeyToIndex(balanceKey)
		balanceMerkleProof := s.tree.GetLeaf(balanceKeyIndex)
		balance.SetString(string(balanceMerkleProof.Value), 16)
		fmt.Println("oldBalance", balance)

		balance.Sub(&balance, userOperation.CalculateGasUsed())
		fmt.Println("newBalance", balance)

		// nonce
		var nonce big.Int
		nonceKey := smt.ComputeNonceKey(userOperation.Sender.Bytes(), userOperation.ChainId.ToInt().Uint64())
		nonceKeyIndex := smt.KeyToIndex(nonceKey)
		nonceMerkleProof := s.tree.GetLeaf(nonceKeyIndex)
		nonce.SetString(string(nonceMerkleProof.Value), 16)
		fmt.Println("oldNonce", nonce.Uint64())
		nonce.Add(&nonce, big.NewInt(1))
		fmt.Println("newNonce", nonce.Uint64())

		// merkle proof
		newBalanceMerkleProof := s.tree.SetLeaf(balanceKeyIndex, smt.MerkleNodeValue(common.Bytes2Hex(common.LeftPadBytes(balance.Bytes(), 32))))
		updateDiff(&newBalanceMerkleProof)

		newNonceMerkleProof := s.tree.SetLeaf(nonceKeyIndex, smt.MerkleNodeValue(common.Bytes2Hex(common.LeftPadBytes(nonce.Bytes(), 32))))
		fmt.Println("new root", s.tree.GetRoot())

		userOperationProof = append(userOperationProof, UserOperationInput{
			UserOperation:     userOperation.UserOperation,
			Signature:         userOperation.Signature[0:64],
			EthRecoveryId:     userOperation.RecoverId(),
			BalanceDeltaProof: &newBalanceMerkleProof,
			NonceDeltaProof:   &newNonceMerkleProof,
		})
		//expectGasBalance, expectNonce := s.storage.Account.GetAccountInfoByAccountAndChainId(userOperation.Sender, userOperation.ChainId.ToInt().Uint64())
		//fmt.Println("expectGasBalance", expectGasBalance)
		//fmt.Println("smtGasBalance", balance)
		//fmt.Println("expectNonce", expectNonce)
		//fmt.Println("smtNonce", nonce)
	}

	batch.SetUserOperationProofs(userOperationProof)

	// mock
	s.proofQueue.Enqueue(*batch)

	if stateDiff != nil {
		s.storage.updateStateDiff(stateDiff)
	} else {
		s.logger.Info("This batch has not stateDiff", "number", s.storage.BatchNumber)
	}
	s.logger.Info("Successfully sealing a batch", "number", s.storage.BatchNumber, "stateRoot", s.tree.GetRoot().String())
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

func (s *State) AddTicket(ticket pool.TicketFull) error {
	return s.storage.Account.AddTicket(ticket)
}

func (s *State) AddUserOperation(userOp pool.UserOperation) error {
	return s.storage.Account.AddUserOperation(userOp)
}

func (s *State) GetAccountsForUser(user common.Address) *[]common.Address {
	return s.storage.Account.GetAccountsForUser(user)
}

func (s *State) GetUserOpsForAccount(user, account common.Address) (*[]pool.UserOperation, error) {
	return s.storage.Account.GetUserOpsForAccount(user, account)
}

func (s *State) GetBalanceAndNonceForAccount(account common.Address, chainId uint64) (*big.Int, uint64) {
	//return s.storage.Account.GetAccountInfoByAccountAndChainId(account, chainId)
	var balance big.Int
	balanceKeyIndex := smt.ComputeBalanceIndex(account.Bytes())
	balanceMerkleProof := s.tree.GetLeaf(balanceKeyIndex)
	balance.SetString(string(balanceMerkleProof.Value), 16)

	var nonce big.Int
	nonceKeyIndex := smt.ComputeNonceIndex(account.Bytes(), chainId)
	nonceMerkleProof := s.tree.GetLeaf(nonceKeyIndex)
	nonce.SetString(string(nonceMerkleProof.Value), 16)
	return &balance, nonce.Uint64()
}

func (s *State) Persistence() {
	s.logger.Info("Cache StateContext info into disk")
	if err := s.storage.cache(); err != nil {
		s.logger.Error("Cache StateContext info into disk error", "error", err)
	}
	if err := s.storage.cacheSmt(s.tree); err != nil {
		s.logger.Error("Cache Smt into disk error", "error", err)
	}
}

func (s *State) LoadCache() {
	s.logger.Info("Load StateContext info from disk")
	if err := s.storage.loadCache(); err != nil {
		s.logger.Error("Load StateContext info from disk error", "error", err)
	}
	tree, err := s.storage.loadCacheSmt()
	if err != nil {
		s.logger.Error("Load Smt info from disk error", "error", err)
	}
	if tree != nil {
		s.tree = tree
		if true {
			diff := s.storage.StateProof[0]
			if err = s.tree.RollbackMerkleTree(*diff); err != nil {
				s.logger.Error("Failed to rollback smt", "error", err)
			} else {
				if s.storage.BatchNumber != 0 {
					s.storage.BatchNumber--
					s.logger.Info("Rollback smt success")
				} else {
					s.storage.BatchNumber = 0
				}

			}
		}
		s.logger.Info("Load smt success", "number", s.storage.BatchNumber, "root", s.tree.GetRoot().String())
	}

}
