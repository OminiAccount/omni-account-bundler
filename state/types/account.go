package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OAB/pool"
	"github.com/OAB/utils/chains"
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

func (info *AccountInfo) deepCopy() *AccountInfo {
	obj := &AccountInfo{
		Nonce:          info.Nonce.Copy(),
		Gas:            new(big.Int).Set(info.Gas),
		UserOperations: info.UserOperations,
	}

	return obj
}

func (info *AccountInfo) increaseNonce(chainId uint64) {
	info.Nonce[chainId]++
}

func (info *AccountInfo) addUserOperations(operation *pool.UserOperation) {
	info.UserOperations = append(info.UserOperations, operation)
}

func (info *AccountInfo) remainingGas(operation *pool.UserOperation) bool {
	usedGas := operation.CalculateGasUsed()
	if info.Gas.Cmp(usedGas) >= 0 {
		info.Gas.Sub(info.Gas, usedGas)
		return true
	}
	return false
}

func (info *AccountInfo) gasOperation(operation *pool.UserOperation) bool {
	value := operation.OperationValue.ToInt()
	switch operation.OperationType {
	case pool.DepositAction:
		info.Gas.Add(info.Gas, value)
		return true
	case pool.WithdrawAction:
		if info.Gas.Cmp(value) < 0 {
			return false
		}
		info.Gas.Sub(info.Gas, value)
		return true
	}
	return false
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

func (u *UserAccount) InitNonce(chainID chains.ChainId, am AccountMapping) {
	ai, err := u.GetAccount(am.User, am.Account)
	if err != nil {
		return
	}

	Lock.Lock()
	defer Lock.Unlock()
	ai.Nonce[uint64(chainID)] = 0
	(*u)[am.User][am.Account] = *ai
}

func (u *UserAccount) AddNewMapping(am AccountMapping) error {
	Lock.Lock()
	defer Lock.Unlock()

	if (*u)[am.User] == nil {
		(*u)[am.User] = make(map[common.Address]AccountInfo)
	}

	// Check if the user already has the maximum number of accounts
	if len((*u)[am.User]) >= UserHaveAccountsMaxNumber {
		return errors.New("user has reached the maximum number of accounts")
	}

	(*u)[am.User][am.Account] = AccountInfo{
		Nonce:          make(map[uint64]uint64),
		Gas:            big.NewInt(0),
		UserOperations: []*pool.UserOperation{},
		Salt:           am.Salt,
	}
	return nil
}

func (u *UserAccount) GetAccount(user, account common.Address) (*AccountInfo, error) {
	Lock.RLock()
	defer Lock.RUnlock()

	accountInfo, exists := (*u)[user][account]
	if !exists {
		return nil, fmt.Errorf("user:%s does not exist in this account:%s", user, account)
	}
	return &accountInfo, nil
}

func (u *UserAccount) AddSignedUserOperation(suo *pool.SignedUserOperation) error {
	accountInfo, err := u.GetAccount(suo.Owner, suo.Sender)
	if err != nil {
		return err
	}

	Lock.Lock()
	defer Lock.Unlock()
	// Create a snapshot, save the original state
	originalAccountInfo := accountInfo.deepCopy()

	var opErr error

	// deposit/withdraw
	if suo.UserOperation.IsGasOperation() {
		if success := accountInfo.gasOperation(suo.UserOperation); success == false {
			opErr = fmt.Errorf("deposit or withdraw operation failed for user:%s", suo.Owner)
		}
	}

	// check gas
	if opErr == nil && !accountInfo.remainingGas(suo.UserOperation) {
		opErr = fmt.Errorf("insufficient gas balance")
	}

	// nonce ++
	if opErr == nil {
		accountInfo.increaseNonce(suo.Exec.ChainId.Uint64())
		expectNonce := accountInfo.Nonce[suo.Exec.ChainId.Uint64()]
		if expectNonce != suo.Exec.Nonce.Uint64() {
			opErr = fmt.Errorf("account:%s nonce mismatch, want:%d, get:%d", suo.Owner, expectNonce, suo.Exec.Nonce.Uint64())
		}
		accountInfo.increaseNonce(suo.InnerExec.ChainId.Uint64())
		expectNonce = accountInfo.Nonce[suo.InnerExec.ChainId.Uint64()]
		if expectNonce != suo.InnerExec.Nonce.Uint64() {
			opErr = fmt.Errorf("account:%s nonce mismatch, want:%d, get:%d", suo.Owner, expectNonce, suo.InnerExec.Nonce.Uint64())
		}
	}

	if opErr == nil {
		accountInfo.addUserOperations(suo.UserOperation)
	} else {
		// If an error occurs, restore the original state
		(*u)[suo.Owner][suo.Sender] = *originalAccountInfo
		return opErr
	}

	// recopy
	(*u)[suo.Owner][suo.Sender] = *accountInfo
	return nil
}

// GetAccountsForUser Iterate and print all account information for a given user
func (u *UserAccount) GetAccountsForUser(user common.Address) map[common.Address]AccountInfo {
	Lock.RLock()
	defer Lock.RUnlock()

	accountsInfo, exists := (*u)[user]
	if !exists {
		return nil
	}

	return accountsInfo
}
