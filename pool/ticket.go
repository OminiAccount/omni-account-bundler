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
	User   common.Address `json:"user"`
	Amount *big.Int       `json:"amount"`
}

type TicketFull struct {
	Ticket
	Type TicketType
}
