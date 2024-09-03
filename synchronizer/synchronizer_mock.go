package synchronizer

import (
	"github.com/OAAC/pool"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func insertTicket(ticketChannel chan<- pool.TicketFull) {
	ticket := pool.TicketFull{
		Ticket: pool.Ticket{
			User:      common.HexToAddress("0xa54753229AD35abC403B53E629A28820C8041eaA"),
			Amount:    (*hexutil.Big)(big.NewInt(10)),
			TimeStamp: uint64(10),
		},
		Type: pool.Deposit,
	}
	ticketChannel <- ticket
}
