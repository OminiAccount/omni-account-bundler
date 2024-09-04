package state

import (
	"context"
	"fmt"
	"github.com/OAAC/pool"
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
	logger      log.Logger
}

func NewState(sc Config, pool PoolInterface, ethereum EthereumInterface) (*State, error) {
	ctx, cancel := context.WithCancel(sc.Context)
	state := &State{
		cfg:         sc,
		ctx:         ctx,
		cancel:      cancel,
		pool:        pool,
		ethereum:    ethereum,
		tree:        smt.NewZeroMerkleTree(50),
		proofQueue:  queue.NewConcurrentQueue[Batch](),
		provenQueue: queue.NewConcurrentQueue[ProofResult](),

		logger: log.New("service", "state"),
	}

	// mock
	state.mockTree()

	return state, nil
}

func (s *State) Start() {
	go s.ProcessBatch()
	go s.ExecuteBatch()
}

func (s *State) Stop() {
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

		userOperationProof = append(userOperationProof, UserOperationProof{
			UserOperation:     userOperation.UserOperation,
			Signature:         userOperation.Signature,
			EthRecoveryId:     userOperation.RecoverId(),
			DomainInfo:        nil,
			BalanceDeltaProof: &newBalanceMerkleProof,
			NonceDeltaProof:   &newNonceMerkleProof,
		})
	}

	batch.SetUserOperationProofs(userOperationProof)

	// Set tickets
	var depositTickets, withdrawTickets []pool.Ticket

	for _, ticketFull := range batchContext.Tickets() {
		var balance big.Int
		balanceKey := smt.ComputeBalanceKey(ticketFull.User.Bytes())
		balanceKeyIndex := smt.KeyToIndex(balanceKey)
		balanceMerkleProof := s.tree.GetLeaf(balanceKeyIndex)
		balance.SetString(string(balanceMerkleProof.Value), 16)
		fmt.Println("Ticket oldBalance", balance)

		actionAmount := ticketFull.Amount.ToInt()
		switch ticketFull.Type {
		// balance + actionAmount
		case pool.Deposit:
			depositTickets = append(depositTickets, ticketFull.Ticket)
			balance.Add(&balance, actionAmount)
		// balance - actionAmount (if balance >= actionAmount)
		case pool.Withdraw:
			withdrawTickets = append(withdrawTickets, ticketFull.Ticket)
			remainBalance := balance.Cmp(actionAmount)
			if remainBalance < 0 {
				return fmt.Errorf("current account balance %s is insufficient to withdraw %s", &balance, actionAmount)
			}
			balance.Sub(&balance, actionAmount)
		}
		fmt.Println("Ticket newBalance", balance)

		// Set newBalance
		s.tree.SetLeaf(balanceKeyIndex, smt.MerkleNodeValue(common.Bytes2Hex(common.LeftPadBytes(balance.Bytes(), 32))))
	}

	batch.SetDepositTickets(depositTickets)
	batch.SetWithdrawTickets(withdrawTickets)

	// mock
	number++
	s.proofQueue.Enqueue(*batch)

	s.logger.Info("Successfully sealing a batch", "number", number)
	return nil
}

func (s *State) ExecuteBatch() {
	s.logger.Info("ExecuteBatch start")
	ticker := time.NewTicker(5 * time.Second)
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
		err := s.ethereum.UpdateEntryPointRoot(res.Proof, res.PublicValues)
		if err != nil {
			return err
		}
	} else {
		s.logger.Debug("the proven queue is empty")
	}

	return nil
}
