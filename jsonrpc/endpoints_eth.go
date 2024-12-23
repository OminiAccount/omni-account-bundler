package jsonrpc

import (
	"fmt"
	"github.com/OAB/jsonrpc/types"
	"github.com/OAB/pool"
	"github.com/OAB/state"
	"github.com/OAB/utils/hex"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

type EthEndpoints struct {
	pool  types.PoolInterface
	state types.StateInterface
	his   types.HisInterface
}

func NewEthEndpoints(p types.PoolInterface, s types.StateInterface, his types.HisInterface) *EthEndpoints {
	return &EthEndpoints{p, s, his}
}

func (e *EthEndpoints) SendUserOperation(suop *state.SignedUserOperation) error {
	log.Infof("[receive op] %+v %+v", suop.Signature, suop.UserOperation)
	if len(suop.Signature) == 0 {
		return fmt.Errorf("invalid signature")
	}
	if suop.Owner.Hex() == "0x" {
		return fmt.Errorf("invalid owner")
	}
	if !e.state.IsSupportChain(suop.Exec.ChainId) {
		return fmt.Errorf("chain:%d no support", suop.Exec.ChainId)
	}
	if suop.InnerExec.ChainId.Uint64() > 0 && !e.state.IsSupportChain(suop.InnerExec.ChainId) {
		return fmt.Errorf("chain:%d no support", suop.InnerExec.ChainId)
	}
	ai, err := e.state.GetAccountInfo(suop.Owner, suop.Sender)
	if err != nil {
		return err
	}
	if _, ok := ai.Chain[suop.Exec.ChainId.Uint64()]; suop.Exec.ChainId.Uint64() > 0 && !ok {
		go e.state.AddFailedCreateAA(ai.Uid, suop.Exec.ChainId.Uint64())
		return fmt.Errorf("chain:%d does not exist this account:%s", suop.Exec.ChainId, suop.Sender)
	}
	if _, ok := ai.Chain[suop.InnerExec.ChainId.Uint64()]; suop.InnerExec.ChainId.Uint64() > 0 && !ok {
		go e.state.AddFailedCreateAA(ai.Uid, suop.InnerExec.ChainId.Uint64())
		return fmt.Errorf("chain:%d does not exist this account:%s", suop.InnerExec.ChainId, suop.Sender)
	}
	if suop.OperationType == state.DepositAction ||
		suop.OperationType == state.WithdrawAction {
		if suop.OperationValue.Uint64() <= 0 {
			return fmt.Errorf("operation value param error")
		}
	} else if suop.OperationValue.Uint64() > 0 {
		return fmt.Errorf("operation value param error")
	}
	return e.state.AddSignedUserOp(ai, suop)
}

func (e *EthEndpoints) GetBatchProof() (interface{}, error) {
	return e.pool.GetBatchProof()
}

func (e *EthEndpoints) SetBatchProofResult(result pool.ProofResult) error {
	return e.pool.SetBatchProofResult(&result)
}

func (e *EthEndpoints) CreateUserAccount(user common.Address) interface{} {
	adr := e.state.GetAccountAdr(user)
	if adr != nil {
		return adr
	}
	return e.state.CreateAccount(user)
}

func (e *EthEndpoints) GetUserAccount(user common.Address) interface{} {
	return e.state.GetAccountAdr(user)
}

func (e *EthEndpoints) GetAccountInfo(user, account common.Address, chainId uint64) (interface{}, error) {
	accInfo, err := e.state.GetAccountInfo(user, account)
	if err != nil {
		return nil, err
	}
	//uos, err := e.state.GetAccountOps(accInfo.Uid)
	//if err != nil {
	//	return nil, err
	//}
	if accInfo.Uid < 1 {
		return nil, fmt.Errorf("account not exist")
	}
	if _, ok := accInfo.Chain[chainId]; !ok {
		return nil, fmt.Errorf("chain:%d does not exist this account", chainId)
	}
	return types.AccountInfo{
		Balance: hex.EncodeBig(accInfo.Chain[chainId].Gas),
		Nonce:   accInfo.Chain[chainId].Nonce + 1,
		//UserOperations: uos,
	}, nil
}

func (e *EthEndpoints) ReportHis(user, account common.Address, data state.AccountHistory) error {
	accInfo, err := e.state.GetAccountInfo(user, account)
	if err != nil {
		return err
	}
	data.Uid = accInfo.Uid
	if data.HisTime < 1 {
		data.HisTime = time.Now().Unix()
	} else if data.HisTime > 10000000000 {
		data.HisTime /= 1000
	}
	data.TimeAt = time.Unix(data.HisTime, 0)
	data.Did = data.UniqueID()
	return e.his.SaveAccountHis(&data)
}

func (e *EthEndpoints) GetUserHistory(user, account common.Address, timestamp uint64) (interface{}, error) {
	ai, err := e.state.GetAccountInfo(user, account)
	if err != nil {
		return nil, err
	}
	if ai.Uid < 1 {
		return nil, fmt.Errorf("account not exist")
	}
	if timestamp < 1 {
		timestamp = uint64(time.Now().Unix())
	}
	return e.his.GetAccountHis(ai.Uid, timestamp, 10)
}
