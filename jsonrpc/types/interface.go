package types

import (
	"github.com/OAB/pool"
	"github.com/OAB/state"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type PoolInterface interface {
	AddUserOp(op pool.SignedUserOperation)
}

type StateInterface interface {
	GetBatchProof() (*state.Batch, error)
	SetBatchProofResult(result *state.ProofResult) error
	GetAccountsForUser(user common.Address) *[]common.Address
	GetBalanceAndNonceForAccount(account common.Address, chainId uint64) (*big.Int, uint64)
	GetUserOpsForAccount(user, account common.Address) (*[]pool.UserOperation, error)
	AddUserOperation(userOp pool.UserOperation) error
}
