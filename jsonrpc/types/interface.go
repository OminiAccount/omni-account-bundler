package types

import (
	"github.com/OAB/pool"
	"github.com/OAB/state"
	"github.com/OAB/state/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type StateInterface interface {
	GetBatchProof() ([]*state.Batch, error)
	SetBatchProofResult(result *state.ProofResult) error
	GetAccountsForUser(user common.Address) *[]common.Address
	GetAccountInfo(user, account common.Address, chainId uint64) (*big.Int, uint64, uint64, []*pool.UserOperation, error)
	AddSignedUserOperation(userOp *pool.SignedUserOperation) error
	CreateAccount(common.Address) *common.Address
}

type HisInterface interface {
	GetAccountHis(user, account common.Address, page uint64) ([]types.AccountHistory, error)
	SaveAccountHis(user, account common.Address, data *types.AccountHistory) error
}
