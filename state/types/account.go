package types

import (
	"errors"
	"fmt"
	"github.com/OAAC/pool"
	"github.com/ethereum/go-ethereum/common"
)

const UserHaveAccountsMaxNumber = 1

type UserAccount map[common.Address]map[common.Address]AccountInfo

type AccountInfo struct {
	Nonce          uint64
	UserOperations []pool.UserOperation
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
		Nonce:          0,
		UserOperations: []pool.UserOperation{},
	}
	return nil
}

func (u *UserAccount) AddUserOperation(userOp pool.UserOperation) error {
	user, exists := u.FindUserByAccount(userOp.Sender)
	if !exists {
		return fmt.Errorf("exists error %s", userOp.Sender)
	}
	accountInfo, exists := (*u)[user][userOp.Sender]
	if !exists {
		return fmt.Errorf("exists error %s", userOp.Sender)
	}
	accountInfo.UserOperations = append(accountInfo.UserOperations, userOp)

	(*u)[user][userOp.Sender] = accountInfo
	return nil
}

func (u *UserAccount) FindUserByAccount(account common.Address) (common.Address, bool) {
	for user, accounts := range *u {
		if _, exists := accounts[account]; exists {
			return user, true
		}
	}
	return common.Address{}, false
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

func (u *UserAccount) GetUserOpsForAccount(user, account common.Address) (*[]pool.UserOperation, error) {
	accountInfo, exists := (*u)[user][account]
	if !exists {
		return nil, fmt.Errorf("exists error %s %s", user, account)
	}
	return &accountInfo.UserOperations, nil
}
