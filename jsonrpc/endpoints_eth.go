package jsonrpc

import (
	"fmt"
	"github.com/OAAC/jsonrpc/types"
	"github.com/OAAC/pool"
	"github.com/OAAC/state"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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
	if len(signedUserOp.InitCode) != 0 {
		return fmt.Errorf("creation of AA contracts is not supported")
	}
	e.pool.AddUserOp(signedUserOp)
	return nil
}

func (e *EthEndpoints) GetBatchProof() (interface{}, error) {
	fmt.Println("getBatchProof")
	return e.state.GetBatchProof()
}

func (e *EthEndpoints) SetBatchProofResult(result state.ProofResult) error {
	return e.state.SetBatchProofResult(&result)
}

func (e *EthEndpoints) GetUserAccount(user common.Address) interface{} {
	return e.state.GetAccountsForUser(user)
}

type AccountInfo struct {
	Balance *big.Int
	Nonce   uint64
}

func (e *EthEndpoints) GetAccountInfo(account common.Address, chainId uint64) interface{} {
	balance, nonce := e.state.GetBalanceAndNonceForAccount(account, chainId)
	return AccountInfo{
		Balance: balance,
		Nonce:   nonce,
	}
}

func (e *EthEndpoints) GetUserOpsForAccount(user, account common.Address) (interface{}, error) {
	return e.state.GetUserOpsForAccount(user, account)
}
