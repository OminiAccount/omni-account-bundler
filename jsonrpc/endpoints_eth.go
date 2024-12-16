package jsonrpc

import (
	"fmt"
	"github.com/OAB/jsonrpc/types"
	"github.com/OAB/pool"
	"github.com/OAB/state"
	state_types "github.com/OAB/state/types"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/common"
)

type EthEndpoints struct {
	state types.StateInterface
	his   types.HisInterface
}

func NewEthEndpoints(s types.StateInterface, his types.HisInterface) *EthEndpoints {
	return &EthEndpoints{s, his}
}

func (e *EthEndpoints) SendUserOperation(suop *pool.SignedUserOperation) error {
	log.Infof("[receive op] %+v", suop.UserOperation)
	if len(suop.Signature) == 0 {
		return fmt.Errorf("invalid signature")
	}
	if suop.Owner.Hex() == "0x" {
		return fmt.Errorf("invalid owner")
	}
	ai, err := e.state.GetAccountInfoByAA(suop.Owner, suop.Sender)
	if err != nil {
		return err
	}
	if _, ok := ai.Nonce[suop.Exec.ChainId.Uint64()]; suop.Exec.ChainId.Uint64() > 0 && !ok {
		return fmt.Errorf("chain:%d does not exist this account:%s", suop.Exec.ChainId, suop.Sender)
	}
	if _, ok := ai.Nonce[suop.InnerExec.ChainId.Uint64()]; suop.InnerExec.ChainId.Uint64() > 0 && !ok {
		return fmt.Errorf("chain:%d does not exist this account:%s", suop.InnerExec.ChainId, suop.Sender)
	}
	return e.state.AddSignedUserOperation(suop)
}

func (e *EthEndpoints) GetBatchProof() (interface{}, error) {
	return e.state.GetBatchProof()
}

func (e *EthEndpoints) SetBatchProofResult(result state.ProofResult) error {
	return e.state.SetBatchProofResult(&result)
}

func (e *EthEndpoints) CreateUserAccount(user common.Address) interface{} {
	adr, _ := e.state.GetAccountInfo(user)
	if adr != nil {
		return fmt.Errorf("account already exist")
	}
	return e.state.CreateAccount(user)
}

func (e *EthEndpoints) GetUserAccount(user common.Address) interface{} {
	return e.state.GetAccountAdrs(user)
}

func (e *EthEndpoints) GetAccountInfo(user, account common.Address, chainId uint64) (interface{}, error) {
	accInfo, err := e.state.GetAccountInfoByAA(user, account)
	if err != nil {
		return nil, err
	}
	return types.AccountInfo{
		Balance:        accInfo.Gas.String(),
		Nonce:          accInfo.Nonce[chainId] + 1,
		UserOperations: accInfo.UserOperations,
		LatestPage:     accInfo.HistoryPage,
	}, nil
}

func (e *EthEndpoints) ReportHis(user, account common.Address, data state_types.AccountHistory) error {
	data.Owner = user.Hex()
	data.Account = account.Hex()
	data.ID = data.UniqueID()
	return e.his.SaveAccountHis(user, account, &data)
}

func (e *EthEndpoints) GetUserHistory(user, account common.Address, page uint64) (interface{}, error) {
	_, err := e.state.GetAccountInfoByAA(user, account)
	if err != nil {
		return nil, err
	}
	return e.his.GetAccountHis(user, account, page)
}
