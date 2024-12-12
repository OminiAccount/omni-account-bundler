package types

import (
	"github.com/OAB/pool"
)

const (
	SignUserOPType = iota
	SwapType
	ExploreType
)

type AccountInfo struct {
	Balance        string
	Nonce          uint64
	LatestPage     uint64
	UserOperations []*pool.UserOperation
}

type UserHistory struct {
	OrderType int
	Data      interface{}
}
