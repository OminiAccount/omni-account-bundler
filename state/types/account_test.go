package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserAccount_AddNewMapping(t *testing.T) {
	userAccounts := NewUserAccount()

	user := common.HexToAddress("0x123")
	account1 := common.HexToAddress("0x456")
	account2 := common.HexToAddress("0x789")

	err := userAccounts.AddNewMapping(AccountMapping{user, account1, 0})
	assert.NoError(t, err, "There should be no errors when adding user accounts")
	err = userAccounts.AddNewMapping(AccountMapping{user, account2, 0})
	assert.Error(t, err, "max number")

	accounts := userAccounts.GetAccountsForUser(user)
	assert.NoError(t, err, "There should be no error when getting user accounts")
	assert.Len(t, len(accounts), 1, "User should have 1 account mappings")
}
