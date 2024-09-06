package types

import (
	"github.com/OAAC/pool"
	"github.com/OAAC/state"
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
}
