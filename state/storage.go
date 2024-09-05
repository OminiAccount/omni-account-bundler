package state

import "github.com/OAAC/state/types"

type Storage struct {
	Account *types.UserAccount
}

func NewStorage() *Storage {
	return &Storage{Account: types.NewUserAccount()}
}
