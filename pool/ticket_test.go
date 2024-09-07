package pool

import (
	"bytes"
	"encoding/gob"
	"fmt"
	msgpack2 "github.com/OAAC/utils/msgpack"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"testing"
)

func TestGob(t *testing.T) {
	amountHex := "0x27147114878000"
	// Create an instance of TicketFull
	bigInt := new(big.Int)
	bigInt, _ = bigInt.SetString(amountHex[2:], 16)
	ticket := &TicketFull{
		Ticket: Ticket{
			User:      common.HexToAddress("0xfD63ed0566A782Ef57F559C6F5f9AfeCE4866423"),
			Amount:    (*hexutil.Big)(bigInt),
			TimeStamp: 1632748392,
		},
		Type: Deposit,
	}

	// Serialize it
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(ticket)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}

	// Deserialize it
	var decodedTicket TicketFull
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(&decodedTicket)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	fmt.Printf("Decoded TicketFull: %+v\n", decodedTicket)
}

func TestMsgPack(t *testing.T) {
	amountHex := "0x27147114878000"
	// Create an instance of TicketFull
	bigInt := new(big.Int)
	bigInt, _ = bigInt.SetString(amountHex[2:], 16)
	ticket := &TicketFull{
		Ticket: Ticket{
			User:      common.HexToAddress("0xfD63ed0566A782Ef57F559C6F5f9AfeCE4866423"),
			Amount:    (*hexutil.Big)(bigInt),
			TimeStamp: 1632748392,
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
