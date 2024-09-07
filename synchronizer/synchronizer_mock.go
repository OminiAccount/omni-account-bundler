package synchronizer

import (
	"github.com/OAAC/pool"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func createAndSendTicket(ticketChannel chan<- pool.TicketFull, userAddress string, amountHex string, timestamp uint64, ticketType pool.TicketType) {
	bigInt := new(big.Int)
	bigInt, _ = bigInt.SetString(amountHex[2:], 16)

	timestampBigInt := new(big.Int).SetUint64(timestamp)

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
		"0xfD63ed0566A782Ef57F559C6F5f9AfeCE4866423",
		"0x27147114878000",
		1725609828,
		pool.Deposit,
	)

	createAndSendTicket(
		ticketChannel,
		"0xfD63ed0566A782Ef57F559C6F5f9AfeCE4866423",
		"0xb1a2bc2ec500000",
		1725616764,
		pool.Deposit,
	)

	createAndSendTicket(
		ticketChannel,
		"0x93d53D2d8f0d623C5cbE46daa818177a450bd9f7",
		"0x470de4df820000",
		1725617724,
		pool.Withdraw,
	)

	createAndSendTicket(
		ticketChannel,
		"0x93d53D2d8f0d623C5cbE46daa818177a450bd9f7",
		"0x6f05b59d3b20000",
		1725617724,
		pool.Deposit,
	)

}
