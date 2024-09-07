package synchronizer

import (
	"github.com/OAAC/pool"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func createAndSendTicket(ticketChannel chan<- pool.TicketFull, userAddress string, amountHex, timestamp string, ticketType pool.TicketType) {
	bigInt := new(big.Int)
	bigInt, _ = bigInt.SetString(amountHex[2:], 16)

	timestampBigInt := new(big.Int)
	timestampBigInt, _ = new(big.Int).SetString(timestamp[2:], 16)

	ticket := pool.TicketFull{
		Ticket: pool.Ticket{
			User:      common.HexToAddress(userAddress),
			Amount:    (*hexutil.Big)(bigInt),
			TimeStamp: (*hexutil.Big)(timestampBigInt),
		},
		Type: ticketType,
	}

	ticketChannel <- ticket
}

func insertTicket(ticketChannel chan<- pool.TicketFull) {
	createAndSendTicket(
		ticketChannel,
		"0xd09d22e15B8c387a023811E5C1021b441B8F0E5a",
		"0x16345785d8a0000",
		"0x66dc6530",
		pool.Deposit,
	)

	createAndSendTicket(
		ticketChannel,
		"0xd09d22e15B8c387a023811E5C1021b441B8F0E5a",
		"0x38d7ea4c68000",
		"0x66dc6560",
		pool.Withdraw,
	)

	createAndSendTicket(
		ticketChannel,
		"0xd09d22e15B8c387a023811E5C1021b441B8F0E5a",
		"0x16345785d8a0000",
		"0x66dc686c",
		pool.Deposit,
	)

}
