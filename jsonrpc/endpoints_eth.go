package jsonrpc

import (
	"fmt"
	"github.com/OAB/jsonrpc/types"
	"github.com/OAB/pool"
	"github.com/OAB/state"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type EthEndpoints struct {
	state types.StateInterface
}

func NewEthEndpoints(s types.StateInterface) *EthEndpoints {
	return &EthEndpoints{s}
}

func (e *EthEndpoints) SendUserOperation(signedUserOp *pool.SignedUserOperation) error {
	if len(signedUserOp.Signature) == 0 {
		return fmt.Errorf("invalid signature")
	}

	return e.state.AddSignedUserOperation(signedUserOp)
}

func (e *EthEndpoints) GetBatchProof() (interface{}, error) {
	return e.state.GetBatchProof()
}

func (e *EthEndpoints) SetBatchProofResult(result state.ProofResult) error {
	return e.state.SetBatchProofResult(&result)
}

func (e *EthEndpoints) GetUserAccount(user common.Address) interface{} {
	return e.state.GetAccountsForUser(user)
}

type AccountInfo struct {
	Balance        *big.Int
	Nonce          uint64
	UserOperations []*pool.UserOperation
}

func (e *EthEndpoints) GetAccountInfo(user, account common.Address, chainId uint64) (interface{}, error) {
	balance, nonce, userOps, err := e.state.GetAccountInfo(user, account, chainId)
	if err != nil {
		return nil, err
	}
	return AccountInfo{
		Balance:        balance,
		Nonce:          nonce,
		UserOperations: userOps,
	}, nil
}
