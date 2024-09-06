package types

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
)

const UserHaveAccountsMaxNumber = 1

type UserAccount map[common.Address]map[common.Address]AccountInfo

type AccountInfo struct {
	Nonce uint64
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
		Nonce: 0,
	}
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
