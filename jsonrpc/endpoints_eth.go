package jsonrpc

import (
	"fmt"
	"github.com/OAAC/jsonrpc/types"
	"github.com/OAAC/pool"
	"github.com/OAAC/state"
)

type EthEndpoints struct {
	pool  types.PoolInterface
	state types.StateInterface
}

func NewEthEndpoints(p types.PoolInterface, s types.StateInterface) *EthEndpoints {
	return &EthEndpoints{p, s}
}

func (e *EthEndpoints) SendUserOperation(signedUserOp pool.SignedUserOperation) error {
	if len(signedUserOp.Signature) == 0 {
		return fmt.Errorf("invalid signature")
	}
	e.pool.AddUserOp(signedUserOp)
	return nil
}

func (e *EthEndpoints) GetBatchProof() (interface{}, error) {
	return e.state.GetBatchProof()
}

func (e *EthEndpoints) SetBatchProofResult(result state.ProofResult) error {
	return e.state.SetBatchProofResult(&result)
}
