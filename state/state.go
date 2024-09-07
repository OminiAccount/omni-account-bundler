package state

import (
	"context"
	"fmt"
	"github.com/OAAC/pool"
	"github.com/OAAC/state/types"
	"github.com/OAAC/utils/queue"
	"github.com/OAAC/utils/smt"
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
		tree:        smt.NewZeroMerkleTree(50),
		proofQueue:  queue.NewConcurrentQueue[Batch](),
		provenQueue: queue.NewConcurrentQueue[ProofResult](),
		storage:     NewStorage(cfg.LevelDB),

		logger: log.New("service", "state"),
	}

	state.LoadCache()

	// mock
	//state.mockTree()

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
			s.processBatch(batchContext)
		case <-ticker.C:
		// Todo: Here you can add processing logic, such as periodically checking the status of the pool or other tasks.
		case <-s.ctx.Done():
			s.logger.Warn("Stopping PreProcessBatch due to context cancellation")
			return
		default:
		}
	}
}

// save to storage
var number uint64

func (s *State) processBatch(batchContext pool.BatchContext) error {
	batch, err := NewBatch(number)
	if err != nil {
		return err
	}

	// Set old smt root
	batch.SetOldSMTRoot(s.tree.GetRoot())

	// Set tickets
	var depositTickets, withdrawTickets []TicketProof

	for _, ticketFull := range batchContext.SortedTickets() {
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

		ticketProof := TicketProof{
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
	}

	batch.SetDepositTickets(depositTickets)
	batch.SetWithdrawTickets(withdrawTickets)

	// Set user operations
	var userOperationProof []UserOperationProof

	for _, userOperation := range batchContext.SignedUserOperations() {
		// balance
		var balance big.Int
		balanceKey := smt.ComputeBalanceKey(userOperation.Sender.Bytes())
		balanceKeyIndex := smt.KeyToIndex(balanceKey)
		balanceMerkleProof := s.tree.GetLeaf(balanceKeyIndex)
		balance.SetString(string(balanceMerkleProof.Value), 16)
		fmt.Println("oldBalance", balance)

		var totalGas big.Int
		totalGas.Add(big.NewInt(int64(userOperation.CallGasLimit)), big.NewInt(int64(userOperation.VerificationGasLimit)))
		totalGas.Add(&totalGas, userOperation.PreVerificationGas.ToInt())
		var totalGasCoeff big.Int
		totalGasCoeff.Add(userOperation.MaxFeePerGas.ToInt(), userOperation.MaxPriorityFeePerGas.ToInt())
		var gas big.Int
		gas.Mul(&totalGas, &totalGasCoeff)
		balance.Sub(&balance, &gas)
		fmt.Println("newBalance", balance)

		// nonce
		var nonce big.Int
		nonceKey := smt.ComputeNonceKey(userOperation.Sender.Bytes(), userOperation.ChainId.ToInt().Uint64())
		nonceKeyIndex := smt.KeyToIndex(nonceKey)
		nonceMerkleProof := s.tree.GetLeaf(nonceKeyIndex)
		nonce.SetString(string(nonceMerkleProof.Value), 16)
		fmt.Println("oldNonce", nonce)
		nonce.Add(&nonce, big.NewInt(1))
		fmt.Println("newNonce", nonce)

		// merkle proof
		newBalanceMerkleProof := s.tree.SetLeaf(balanceKeyIndex, smt.MerkleNodeValue(common.Bytes2Hex(common.LeftPadBytes(balance.Bytes(), 32))))

		newNonceMerkleProof := s.tree.SetLeaf(nonceKeyIndex, smt.MerkleNodeValue(common.Bytes2Hex(common.LeftPadBytes(nonce.Bytes(), 32))))
		domainInfo := pool.EIP712Domain{
			Name:              "OMNI-ACCOUNT",
			Version:           "1.0",
			ChainId:           11155111,
			VerifyingContract: common.HexToAddress("0x5F2464f924b7D9166a870cCe9201AFBC2a2f151D"),
		}
		userOperationProof = append(userOperationProof, UserOperationProof{
			UserOperation:     userOperation.UserOperation,
			Signature:         userOperation.Signature,
			EthRecoveryId:     userOperation.RecoverId(),
			DomainInfo:        domainInfo,
			BalanceDeltaProof: &newBalanceMerkleProof,
			NonceDeltaProof:   &newNonceMerkleProof,
		})
	}

	batch.SetUserOperationProofs(userOperationProof)

	// mock
	number++
	s.proofQueue.Enqueue(*batch)

	s.logger.Info("Successfully sealing a batch", "number", number)
	return nil
}

func (s *State) ExecuteBatch() {
	s.logger.Info("ExecuteBatch start")
	ticker := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-ticker.C:
			s.executeBatch()
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
	} else {
		s.logger.Debug("the proven queue is empty")
	}

	return nil
}

func (s *State) AddNewMapping(mapping types.AccountMapping) error {
	return s.storage.Account.AddNewMapping(mapping)
}

func (s *State) GetAccountsForUser(user common.Address) *[]common.Address {
	return s.storage.Account.GetAccountsForUser(user)
}

func (s *State) GetBalanceAndNonceForAccount(account common.Address, chainId uint64) (*big.Int, uint64) {
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
	s.logger.Info("Storage persistence")
	if err := s.storage.persistence(); err != nil {
		s.logger.Error("Storage persistence error", "error", err)
	}
}

func (s *State) LoadCache() {
	s.logger.Info("Load persistence")
	if err := s.storage.loadDisk(); err != nil {
		s.logger.Error("Load persistence error", "error", err)
	}
}
