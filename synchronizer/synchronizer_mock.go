package synchronizer

import (
	"github.com/OAB/pool"
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
		"0xD31959035048676fc27d84C8Bc120997204b16B6",
		"0x6f05b59d3b20000",
		"0x66e2aacc",
		pool.Deposit,
	)

	createAndSendTicket(
		ticketChannel,
		"0xD31959035048676fc27d84C8Bc120997204b16B6",
		"0x16345785d8a0000",
		"0x66e2ab08",
		pool.Withdraw,
	)

}
