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

func TestSignedUserOperation(t *testing.T) {
	testPrk := "82693fc767eb00e3288a01b7516f3a98269882e951be74403e4061d898ea0929"
	testUser := common.HexToAddress("0x7c38C1646213255f62dB509688B8fA062e0Ed8e4")
	testAcc := common.HexToAddress("0x65278ef7227697f934d03ef3dcb8f0f6ee822230")
	suo := &SignedUserOperation{
		UserOperation: &UserOperation{
			Uid:            1,
			Owner:          testUser,
			OperationType:  DepositAction,
			OperationValue: (*hexutil.Big)(big.NewInt(20000000000000000)),
			Sender:         testAcc,
			Exec: ExecData{
				ChainId:                28516,
				Nonce:                  hexutil.Uint64(1),
				CallData:               hexutil.Bytes(""),
				MainChainGasPrice:      (*hexutil.Big)(big.NewInt(20)),
				MainChainGasLimit:      hexutil.Uint64(10),
				DestChainGasPrice:      (*hexutil.Big)(big.NewInt(20)),
				DestChainGasLimit:      hexutil.Uint64(10),
				ZkVerificationGasLimit: hexutil.Uint64(10),
			},
			InnerExec: ExecData{},
			Status:    PendingStatus,
		},
	}
	privateKey, err := crypto.HexToECDSA(testPrk)
	if err != nil {
		t.Errorf("Error generating private key: %v", err)
		return
	}
	signature, err := crypto.Sign(accounts.TextHash(suo.CalculateEIP712TypeDataHash()), privateKey)
	if err != nil {
		t.Errorf("Failed to sign hash typed data: %v", err)
		return
	}
	signature[crypto.RecoveryIDOffset] += 27
	suo.Signature = signature
	suo.Did = getDepositID(suo)

	sigStr := suo.Signature.String()
	//t.Log(sigStr)
	sigBy := common.Hex2Bytes(strings.TrimPrefix(sigStr, "0x"))
	sigStr2 := "0x" + common.Bytes2Hex(sigBy)
	t.Log(sigStr == sigStr2)

	owner := suo.RecoverAddress()
	t.Log(owner.Hex())
	t.Log(testUser.Hex())
}
