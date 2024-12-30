package types

import (
	"github.com/OAB/pool"
	"github.com/OAB/state"
	"github.com/ethereum/go-ethereum/common"
)

type PoolInterface interface {
	GetBatchProof() ([]*pool.Batch, error)
	SetBatchProofResult(*pool.ProofResult) error
}

type StateInterface interface {
	IsSupportChain(uint64) bool
	GetAccountInfo(common.Address, common.Address) (*state.AccountInfo, error)
	GetAccountInfoByChain(common.Address, uint64) (*state.AccountInfo, error)
	GetAccountOps(uid uint64) ([]*state.UserOperation, error)
	GetUser(user, account common.Address) *state.AccountInfo
	AddSignedUserOp(*state.AccountInfo, *state.SignedUserOperation) error
	CreateVizingAccount(common.Address) *common.Address
	CreateOtherAccount(common.Address, uint64, uint64)
	AddFailedCreateAA(uint64, uint64)
}

type HisInterface interface {
	GetAccountHis(uint64, uint64, uint64) ([]state.AccountHistory, error)
	SaveAccountHis(*state.AccountHistory) error
}
