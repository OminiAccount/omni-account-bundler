package jsonrpc

import (
	"fmt"
	"github.com/OAB/jsonrpc/types"
	"github.com/OAB/pool"
	"github.com/OAB/state"
	state_types "github.com/OAB/state/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EthEndpoints struct {
	state types.StateInterface
	his   types.HisInterface
}

func NewEthEndpoints(s types.StateInterface, his types.HisInterface) *EthEndpoints {
	return &EthEndpoints{s, his}
}

func (e *EthEndpoints) SendUserOperation(signedUserOp *pool.SignedUserOperation) error {
	if len(signedUserOp.Signature) == 0 {
		return fmt.Errorf("invalid signature")
	}
	if signedUserOp.Did == "" {
		hashBytes := crypto.Keccak256Hash(signedUserOp.Encode())
		signedUserOp.Did = hashBytes.Hex()
	}
	return e.state.AddSignedUserOperation(signedUserOp)
}

func (e *EthEndpoints) GetBatchProof() (interface{}, error) {
	return e.state.GetBatchProof()
}

func (e *EthEndpoints) SetBatchProofResult(result state.ProofResult) error {
	return e.state.SetBatchProofResult(&result)
}

func (e *EthEndpoints) CreateUserAccount(user common.Address) interface{} {
	return e.state.CreateAccount(user)
}

func (e *EthEndpoints) GetUserAccount(user common.Address) interface{} {
	return e.state.GetAccountsForUser(user)
}

func (e *EthEndpoints) GetAccountInfo(user, account common.Address, chainId uint64) (interface{}, error) {
	balance, nonce, page, userOps, err := e.state.GetAccountInfo(user, account, chainId)
	if err != nil {
		return nil, err
	}
	return types.AccountInfo{
		Balance:        balance.String(),
		Nonce:          nonce + 1,
		UserOperations: userOps,
		LatestPage:     page,
	}, nil
}

func (e *EthEndpoints) ReportHis(user, account common.Address, data state_types.AccountHistory) error {
	data.Owner = user.Hex()
	data.Account = account.Hex()
	data.ID = data.UniqueID()
	return e.his.SaveAccountHis(user, account, &data)
}

func (e *EthEndpoints) GetUserHistory(user, account common.Address, page uint64) (interface{}, error) {
	_, _, _, _, err := e.state.GetAccountInfo(user, account, 0)
	if err != nil {
		return nil, err
	}
	return e.his.GetAccountHis(user, account, page)
}
