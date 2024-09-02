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
	User      common.Address
	Amount    *hexutil.Big
	TimeStamp uint64
	Type      TicketType
}
