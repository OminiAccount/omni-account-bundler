package types

import (
	"github.com/OAB/pool"
	"github.com/OAB/state"
	"github.com/OAB/state/types"
	"github.com/ethereum/go-ethereum/common"
)

type StateInterface interface {
	GetBatchProof() ([]*state.Batch, error)
	SetBatchProofResult(*state.ProofResult) error
	GetAccountAdrs(common.Address) *[]common.Address
	GetAccountInfoByAA(common.Address, common.Address) (*types.AccountInfo, error)
	GetAccountInfo(common.Address) (*common.Address, *types.AccountInfo)
	AddSignedUserOperation(*pool.SignedUserOperation) error
	CreateAccount(common.Address) *common.Address
}

type HisInterface interface {
	GetAccountHis(user, account common.Address, page uint64) ([]types.AccountHistory, error)
	SaveAccountHis(user, account common.Address, data *types.AccountHistory) error
}
