package pool

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type TicketType uint8

const (
	Deposit TicketType = iota
	Withdraw
)

type Ticket struct {
	Did    string         `json:"did"`
	User   common.Address `json:"user"`
	Amount *big.Int       `json:"amount"`
}

type TicketFull struct {
	Ticket
	SignedUserOp *SignedUserOperation
	Type         TicketType
	TimeAt       int64
}
