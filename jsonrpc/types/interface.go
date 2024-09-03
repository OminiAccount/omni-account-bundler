package types

import (
	"github.com/OAAC/pool"
	"github.com/OAAC/state"
)

type PoolInterface interface {
	AddUserOp(op pool.SignedUserOperation)
}

type StateInterface interface {
	GetBatchProof() (*state.Batch, error)
	SetBatchProofResult(result *state.ProofResult) error
}
