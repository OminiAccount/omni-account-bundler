package types

import (
	"github.com/OAB/pool"
	"github.com/OAB/state"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type StateInterface interface {
	GetBatchProof() (*state.Batch, error)
	SetBatchProofResult(result *state.ProofResult) error
	GetAccountsForUser(user common.Address) *[]common.Address
	GetAccountInfo(user, account common.Address, chainId uint64) (*big.Int, uint64, []*pool.UserOperation, error)
	AddSignedUserOperation(userOp *pool.SignedUserOperation) error
}
