package state

import (
	"github.com/OAAC/pool"
	"github.com/OAAC/utils/smt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type UserOperationProof struct {
	UserOperation     *pool.UserOperation   `json:"user_operation"`
	Signature         hexutil.Bytes         `json:"sig_bytes"`
	EthRecoveryId     uint8                 `json:"eth_reconvery_id"`
	DomainInfo        interface{}           `json:"domain_info"`
	BalanceDeltaProof *smt.DeltaMerkleProof `json:"balance_delta_proof"`
	NonceDeltaProof   *smt.DeltaMerkleProof `json:"nonce_delta_proof"`
}

type TicketProof struct {
	Ticket      *pool.Ticket          `json:"ticket"`
	TicketProof *smt.DeltaMerkleProof `json:"delta_proof"`
}

type Batch struct {
	Number              uint64 `json:"number"`
	timestamp           uint64
	UserOperationProofs *[]UserOperationProof `json:"userop_inputs"`
	DepositTickets      *[]TicketProof        `json:"d_ticket_inputs"`
	WithdrawTickets     *[]TicketProof        `json:"w_ticket_inputs"`
	OldSMTRoot          common.Hash           `json:"old_smt_root"`
}

func NewBatch(number uint64) (*Batch, error) {
	batch := &Batch{
		Number:    number,
		timestamp: 0,
	}
	return batch, nil
}

func (b *Batch) SetUserOperationProofs(value []UserOperationProof) {
	b.UserOperationProofs = &value
}

func (b *Batch) SetDepositTickets(value []TicketProof) {
	b.DepositTickets = &value
}

func (b *Batch) SetWithdrawTickets(value []TicketProof) {
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
