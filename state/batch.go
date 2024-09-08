package state

import (
	"github.com/OAB/pool"
	"github.com/OAB/utils/smt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type UserOperationInput struct {
	UserOperation     *pool.UserOperation   `json:"user_operation"`
	Signature         hexutil.Bytes         `json:"sig_bytes"`
	EthRecoveryId     uint8                 `json:"eth_reconvery_id"`
	BalanceDeltaProof *smt.DeltaMerkleProof `json:"balance_delta_proof"`
	NonceDeltaProof   *smt.DeltaMerkleProof `json:"nonce_delta_proof"`
}

type TicketInput struct {
	Ticket      *pool.Ticket          `json:"ticket"`
	TicketProof *smt.DeltaMerkleProof `json:"delta_proof"`
}

type Batch struct {
	Number              uint64 `json:"number"`
	timestamp           uint64
	UserOperationProofs *[]UserOperationInput `json:"userop_inputs"`
	DepositTickets      *[]TicketInput        `json:"d_ticket_inputs"`
	WithdrawTickets     *[]TicketInput        `json:"w_ticket_inputs"`
	OldSMTRoot          common.Hash           `json:"old_smt_root"`
}

func NewBatch(number uint64) (*Batch, error) {
	batch := &Batch{
		Number:    number,
		timestamp: 0,
	}
	return batch, nil
}

func (b *Batch) SetUserOperationProofs(value []UserOperationInput) {
	b.UserOperationProofs = &value
}

func (b *Batch) SetDepositTickets(value []TicketInput) {
	b.DepositTickets = &value
}

func (b *Batch) SetWithdrawTickets(value []TicketInput) {
	b.WithdrawTickets = &value
}

func (b *Batch) SetOldSMTRoot(value common.Hash) {
	b.OldSMTRoot = value
}

type ProofResult struct {
	Number       uint64        `json:"number"`
	Proof        hexutil.Bytes `json:"proof"`
	PublicValues hexutil.Bytes `json:"public_values"`
}
