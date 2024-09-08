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
		"0x5C7feffd7955E5fCA77e64f01cC911C255Ee6d55",
		"0x16345785d8a0000",
		"0x66dd58f0",
		pool.Deposit,
	)

	createAndSendTicket(
		ticketChannel,
		"0x5C7feffd7955E5fCA77e64f01cC911C255Ee6d55",
		"0x2c68af0bb140000",
		"0x66dd58fc",
		pool.Deposit,
	)

	createAndSendTicket(
		ticketChannel,
		"0x5C7feffd7955E5fCA77e64f01cC911C255Ee6d55",
		"0x2386f26fc10000",
		"0x66dd59a4",
		pool.Withdraw,
	)

}
