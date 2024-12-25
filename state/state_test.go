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
	"testing"
)

func TestDepositDid(t *testing.T) {
	sender := common.HexToAddress("0x3307c5023677a4c085855117fa8bbce765e9246f")
	chainID := hexutil.Uint64(28516)
	nonce := hexutil.Uint64(3)
	val := big.NewInt(20000000000000000)

	var encodeBytes []byte
	encodeBytes = append(encodeBytes, sender.Bytes()...)
	encodeBytes = append(encodeBytes, common.LeftPadBytes(big.NewInt(int64(chainID)).Bytes(), 32)...)
	encodeBytes = append(encodeBytes, common.LeftPadBytes(big.NewInt(int64(nonce)).Bytes(), 32)...)
	encodeBytes = append(encodeBytes, common.LeftPadBytes(val.Bytes(), 32)...)
	hash := crypto.Keccak256Hash(encodeBytes).Hex()
	t.Log(hash)
}

func TestBlock(t *testing.T) {
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
