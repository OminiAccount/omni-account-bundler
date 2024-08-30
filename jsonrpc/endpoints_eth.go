package jsonrpc

import (
	"fmt"
	"github.com/OAAC/jsonrpc/types"
	"github.com/OAAC/pool"
)

type EthEndpoints struct {
	pool types.PoolInterface
}

func NewEthEndpoints(p types.PoolInterface) *EthEndpoints {
	return &EthEndpoints{p}
}

func (e *EthEndpoints) SendUserOperation(signedUserOp pool.SignedUserOperation) error {
	// ctx := context.Background()
	if len(signedUserOp.Signature) == 0 {
		return fmt.Errorf("invalid signature")
	}
	e.pool.AddUserOp(signedUserOp)
	return nil
}
