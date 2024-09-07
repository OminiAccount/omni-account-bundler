package pool

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type TicketType uint8

const (
	Deposit TicketType = iota
	Withdraw
)

type Ticket struct {
	User      common.Address `json:"user"`
	Amount    *hexutil.Big   `json:"amount"`
	TimeStamp *hexutil.Big   `json:"timestamp"`
}

type TicketFull struct {
	Ticket
	Type TicketType
}
