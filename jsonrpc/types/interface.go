package types

import (
	"github.com/OAB/pool"
	"github.com/OAB/state"
	"github.com/OAB/state/types"
	"github.com/ethereum/go-ethereum/common"
)

type StateInterface interface {
	GetBatchProof() ([]state.Batch, error)
	SetBatchProofResult(*state.ProofResult) error
	GetAccountAdrs(common.Address) *[]common.Address
	GetAccountInfo(common.Address, common.Address) (*state.AccountInfo, error)
	GetAccountOps(uid uint64) ([]*pool.UserOperation, error)
	GetAccountAdr(common.Address) *common.Address
	AddSignedUserOp(*state.AccountInfo, *pool.SignedUserOperation) error
	CreateAccount(common.Address) *common.Address
}

type HisInterface interface {
	GetAccountHis(user, account common.Address, page uint64) ([]types.AccountHistory, error)
	SaveAccountHis(user, account common.Address, data *types.AccountHistory) error
}
