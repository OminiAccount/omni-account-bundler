package pool

import (
	"fmt"
	msgpack2 "github.com/OAAC/utils/msgpack"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"testing"
)

func TestMsgPack(t *testing.T) {
	amountHex := "0x27147114878000"
	// Create an instance of TicketFull
	bigInt := new(big.Int)
	bigInt, _ = bigInt.SetString(amountHex[2:], 16)
	timestampBigInt := new(big.Int).SetUint64(1632748392)
	ticket := &TicketFull{
		Ticket: Ticket{
			User:      common.HexToAddress("0xfD63ed0566A782Ef57F559C6F5f9AfeCE4866423"),
			Amount:    (*hexutil.Big)(bigInt),
			TimeStamp: (*hexutil.Big)(timestampBigInt),
		},
		Type: Deposit,
	}

	data, err := msgpack2.MarshalStruct(ticket)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	decodeTicket, err := msgpack2.UnmarshalStruct[TicketFull](data)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	fmt.Printf("Decoded ticket: %+v\n", decodeTicket)

	fmt.Println("Decoded big.Int:", decodeTicket)
}
