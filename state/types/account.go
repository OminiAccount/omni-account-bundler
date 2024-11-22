package types

import (
	"errors"
	"fmt"
	"github.com/OAB/pool"
	"github.com/ethereum/go-ethereum/common"
	"maps"
	"math/big"
)

const UserHaveAccountsMaxNumber = 1

type Nonce map[uint64]uint64

func (n Nonce) Copy() Nonce {
	return maps.Clone(n)
}

type UserAccount map[common.Address]map[common.Address]AccountInfo

type AccountInfo struct {
	Nonce          Nonce
	Gas            *big.Int
	UserOperations []*pool.UserOperation
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
	value := new(big.Int).SetUint64(operation.OperationValue.Uint64())
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
}

func NewUserAccount() *UserAccount {
	ua := make(UserAccount)
	return &ua
}

func (u *UserAccount) AddNewMapping(mapping AccountMapping) error {
	if (*u)[mapping.User] == nil {
		(*u)[mapping.User] = make(map[common.Address]AccountInfo)
	}

	// Check if the user already has the maximum number of accounts
	if len((*u)[mapping.User]) >= UserHaveAccountsMaxNumber {
		return errors.New("user has reached the maximum number of accounts")
	}

	(*u)[mapping.User][mapping.Account] = AccountInfo{
		Nonce:          make(map[uint64]uint64),
		Gas:            big.NewInt(0),
		UserOperations: []*pool.UserOperation{},
	}
	return nil
}

func (u *UserAccount) GetAccount(user, account common.Address) (*AccountInfo, error) {
	accountInfo, exists := (*u)[user][account]
	if !exists {
		return nil, fmt.Errorf("user:%s does not exist in this account:%s", user, account)
	}
	return &accountInfo, nil
}

func (u *UserAccount) AddSignedUserOperation(signedUserOp *pool.SignedUserOperation) error {

	accountInfo, err := u.GetAccount(signedUserOp.Owner, signedUserOp.Sender)
	if err != nil {
		return err
	}

	// Create a snapshot, save the original state
	originalAccountInfo := accountInfo.deepCopy()

	var opErr error

	// deposit/withdraw
	if signedUserOp.UserOperation.IsGasOperation() {
		if success := accountInfo.gasOperation(signedUserOp.UserOperation); success == false {
			opErr = fmt.Errorf("deposit or withdraw operation failed for user:%s", signedUserOp.Owner)
		}
	}

	// check gas
	if opErr == nil && !accountInfo.remainingGas(signedUserOp.UserOperation) {
		opErr = fmt.Errorf("insufficient gas balance")
	}

	// nonce ++
	if opErr == nil {
		accountInfo.increaseNonce(signedUserOp.ChainId.Uint64())
		expectNonce := accountInfo.Nonce[signedUserOp.ChainId.Uint64()]
		if expectNonce != signedUserOp.Nonce.Uint64() {
			opErr = fmt.Errorf("account:%s nonce mismatch, want:%d, get:%d", signedUserOp.Owner, expectNonce, signedUserOp.Nonce.Uint64())
		}
	}

	if opErr == nil {
		accountInfo.addUserOperations(signedUserOp.UserOperation)
	} else {
		// If an error occurs, restore the original state
		(*u)[signedUserOp.Owner][signedUserOp.Sender] = *originalAccountInfo
		return opErr
	}

	// recopy
	(*u)[signedUserOp.Owner][signedUserOp.Sender] = *accountInfo
	return nil
}

// GetAccountsForUser Iterate and print all account information for a given user
func (u *UserAccount) GetAccountsForUser(user common.Address) *[]common.Address {
	accountsInfo, exists := (*u)[user]
	if !exists {
		return nil
	}

	var accounts []common.Address
	for account := range accountsInfo {
		accounts = append(accounts, account)
	}

	return &accounts
}
