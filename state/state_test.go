package state

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/pool"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
)

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

func TestEncodeCircuitInput(t *testing.T) {
	userOpDeposit := pool.UserOperation{
		OperationType:  1,
		OperationValue: (*hexutil.Big)(big.NewInt(100)),
		Sender:         common.HexToAddress("0xd09d22e15b8c387a023811e5c1021b441b8f0e5a"),
		Exec: pool.ExecData{
			Nonce:                  1,
			ChainId:                hexutil.Uint64(28516),
			CallData:               common.FromHex("0x"),
			MainChainGasLimit:      0x30d40,
			DestChainGasLimit:      0,
			ZkVerificationGasLimit: 0x898,
			MainChainGasPrice:      (*hexutil.Big)(big.NewInt(100)),
			DestChainGasPrice:      (*hexutil.Big)(big.NewInt(100)),
		},
	}
	t.Logf("%+v", userOpDeposit)
	dataHash := userOpDeposit.CalculateEIP712TypeDataHash()
	signedUserOperation := pool.SignedUserOperation{
		UserOperation: &userOpDeposit,
		Signature:     signMessage(dataHash),
	}
	t.Logf("%+v", signedUserOperation.Signature)
	var sus []*pool.SignedUserOperation
	sus = append(sus, &signedUserOperation)
	by := encodeCircuitInput(sus)
	t.Log(by)
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
