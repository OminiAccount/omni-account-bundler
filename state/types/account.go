package types

import (
	"errors"
	"fmt"
	"github.com/OAB/pool"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const UserHaveAccountsMaxNumber = 1

type UserAccount map[common.Address]map[common.Address]AccountInfo

type AccountInfo struct {
	Nonce          map[uint64]uint64
	Gas            *big.Int
	UserOperations []pool.UserOperation
}

func (info *AccountInfo) IncreaseNonce(chainId uint64) {
	info.Nonce[chainId]++
}

func (info *AccountInfo) RemainingGas(operation pool.UserOperation) bool {
	usedGas := operation.CalculateGasUsed()
	if info.Gas.Cmp(usedGas) >= 0 {
		info.Gas.Sub(info.Gas, usedGas)
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
		UserOperations: []pool.UserOperation{},
	}
	return nil
}

func (u *UserAccount) AddTicket(ticket pool.TicketFull) error {
	// ticket.User is Account address
	user, exists := u.FindUserByAccount(ticket.User)
	if !exists {
		return fmt.Errorf("exists error %s", ticket.User)
	}
	accountInfo, exists := (*u)[user][ticket.User]
	if !exists {
		return fmt.Errorf("exists error %s", ticket.User)
	}
	actionAmount := ticket.Amount.ToInt()

	balance := accountInfo.Gas

	operation := func() error {
		switch ticket.Type {
		case pool.Deposit:
			balance.Add(balance, actionAmount)
		case pool.Withdraw:
			if balance.Cmp(actionAmount) < 0 {
				return fmt.Errorf("current account balance %s is insufficient to withdraw %s", &balance, actionAmount)
			}
			balance.Sub(balance, actionAmount)
		default:
			return fmt.Errorf("unsupported ticket type: %v", ticket.Type)
		}
		return nil
	}

	accountInfo.Gas = balance

	if err := operation(); err != nil {
		return err
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
	//// check gas
	//if accountInfo.RemainingGas(userOp) {
	//	return fmt.Errorf("insufficient gas balance")
	//}
	//
	//// nonce ++
	//accountInfo.IncreaseNonce(uint64(userOp.Nonce))
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

func (u *UserAccount) GetAccountInfoByAccountAndChainId(account common.Address, chainId uint64) (*big.Int, uint64) {
	user, exists := u.FindUserByAccount(account)
	if !exists {
		return nil, 0
	}
	accountInfo, exists := (*u)[user][account]
	if !exists {
		return nil, 0
	}
	return accountInfo.Gas, accountInfo.Nonce[chainId]
}
