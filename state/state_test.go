package state

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/OAB/lib/common/hexutil"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"testing"
)

func TestBlock(t *testing.T) {
	encodeStr := "0x0101000000000000000000000000000000000000000000000000470de4df820000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000006f64c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4700000000000000000000000000000000a0000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000001400000000000000000000000000000014307866623130313134663630663636666463386630373764313631653765646163323430313134396262323438663362626330326437343439393730326535643535343564613831333333383937306139333732396363306235313733356462313463613066393631336434393466303464306666643462383433313866366136663162"
	l2data := common.Hex2Bytes(strings.TrimPrefix(encodeStr, "0x"))
	encode := hexutil.Encode(l2data)
	t.Log(encode)
	if encode == encodeStr {
		t.Log(true)
	} else {
		t.Log(false)
	}
	return

	cli, err := ethclient.Dial("https://sepolia-rollup.arbitrum.io/rpc")
	if err != nil {
		t.Fatal(err)
	}

	var raw json.RawMessage
	err = cli.Client().Call(&raw, "eth_getBlockByNumber", hexutil.EncodeBig(big.NewInt(106514310)), true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(raw)

	header, err := cli.HeaderByNumber(context.Background(), big.NewInt(106514310))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(header.Number)
	t.Log(header.Time)
	t.Log(header.ParentHash)
	t.Log(header.Hash())
}

func hexTo(hex string) *hexutil.Big {
	bigInt := new(big.Int)
	bigInt, _ = bigInt.SetString(hex[2:], 16)
	return (*hexutil.Big)(bigInt)
}

func signMessage(dataHash []byte) (signature []byte) {
	privateKey, err := crypto.HexToECDSA("1cb564eaae3cc8d2f4bfbc2b352c114a4de315288155d6f12cfa418aa8929864")
	if err != nil {
		fmt.Println("Error generating private key:", err)
	}
	signature, err = crypto.Sign(accounts.TextHash(dataHash), privateKey)
	if err != nil {
		fmt.Printf("Failed to sign hash typed data: %v\n", err)
	}

	signature[crypto.RecoveryIDOffset] += 27 // Transform yellow paper V from 0/1 to 27/28

	return signature
}
