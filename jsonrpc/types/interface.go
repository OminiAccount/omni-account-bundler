package types

import (
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/pool"
	"github.com/OAB/state"
	"github.com/ethereum/go-ethereum/common"
)

type PoolInterface interface {
	GetBatchProof() ([]*pool.Batch, error)
	SetBatchProofResult(*pool.ProofResult) error
}

type StateInterface interface {
	IsSupportChain(hexutil.Uint64) bool
	GetAccountInfo(common.Address, common.Address) (*state.AccountInfo, error)
	GetAccountOps(uid uint64) ([]*state.UserOperation, error)
	GetAccountAdr(common.Address) *common.Address
	AddSignedUserOp(*state.AccountInfo, *state.SignedUserOperation) error
	CreateAccount(common.Address) *common.Address
	AddFailedCreateAA(uint64, uint64)
}

type HisInterface interface {
	GetAccountHis(uint64, uint64, uint64) ([]state.AccountHistory, error)
	SaveAccountHis(*state.AccountHistory) error
}
