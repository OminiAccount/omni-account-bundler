package types

import (
	"errors"
	"fmt"
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

// Iterate and print all account information for a given user
func (u *UserAccount) getAccountsForUser(user common.Address) (*[]common.Address, error) {
	accountsInfo, exists := (*u)[user]
	if !exists {
		return nil, fmt.Errorf("no account mappings exist for this user")
	}

	fmt.Printf("user %s account:\n", user.Hex())
	var accounts []common.Address
	for account, info := range accountsInfo {
		fmt.Printf("Account: %s, Nonce: %d\n", account.Hex(), info.Nonce)
		accounts = append(accounts, account)
	}

	return &accounts, nil
}

func (u *UserAccount) accountSaveInDisk() {

}
