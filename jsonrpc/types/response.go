package types

import (
	"github.com/OAB/state"
)

const (
	SignUserOPType = iota
	SwapType
	ExploreType
)

type AccountInfo struct {
	Balance        string
	Nonce          uint64
	UserOperations []*state.UserOperation
}

type UserHistory struct {
	OrderType int
	Data      interface{}
}
