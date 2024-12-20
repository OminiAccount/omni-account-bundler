package types

import (
	"encoding/json"
	"fmt"
	"github.com/OAB/pool"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"maps"
	"math/big"
	"sync"
	"time"
)

const UserHaveAccountsMaxNumber = 1
const HistoryPageSize = 10

var Lock sync.RWMutex

type Nonce map[uint64]uint64

func (n Nonce) Copy() Nonce {
	return maps.Clone(n)
}

type UserAccount map[common.Address]map[common.Address]AccountInfo

type Token struct {
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Logo     string `json:"logo"`
}

type HistoryDetail struct {
	Token   Token  `json:"token"`
	Address string `json:"address"`
	Value   string `json:"value"`
}

type AccountHistory struct {
	ID          string        `json:"id"`
	Owner       string        `json:"owner"`
	Account     string        `json:"account"`
	OrderType   uint8         `json:"orderType"`
	From        HistoryDetail `json:"from"`
	To          HistoryDetail `json:"to"`
	SourceChain uint64        `json:"sourceChain"`
	TargetChain uint64        `json:"targetChain"`
	SourceHash  string        `json:"sourceHash"`
	TargetHash  string        `json:"targetHash"`
	SwapHash    string        `json:"swapHash"`
	Time        int64         `json:"time"`
}

func (s AccountHistory) UniqueID() string {
	data, _ := json.Marshal(s)
	hashBytes := crypto.Keccak256Hash(data)
	return hashBytes.Hex()
}

func ToUserOperationHis(tx string, op *pool.UserOperation) AccountHistory {
	uoHis := AccountHistory{
		ID:        op.Did,
		Owner:     op.Owner.Hex(),
		Account:   op.Sender.Hex(),
		OrderType: op.OperationType,
		From: HistoryDetail{
			Address: op.Sender.Hex(),
			Value:   op.OperationValue.String(),
		},
		SourceChain: op.Exec.ChainId.Uint64(),
		SourceHash:  tx,
		To: HistoryDetail{
			Address: op.Sender.Hex(),
			Value:   op.OperationValue.String(),
		},
		TargetChain: op.Exec.ChainId.Uint64(),
		TargetHash:  tx,
		Time:        time.Now().Unix(),
	}
	return uoHis
}

type AccountInfo struct {
	Salt           uint64
	Nonce          Nonce
	Gas            *big.Int
	UserOperations []*pool.UserOperation
	HistoryPage    uint64
}

type AccountMapping struct {
	User    common.Address
	Account common.Address
	Salt    uint64
}

func NewUserAccount() *UserAccount {
	ua := make(UserAccount)
	return &ua
}
