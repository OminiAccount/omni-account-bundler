package state

import (
	"context"
	"encoding/json"
	"github.com/OAB/lib/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
)

func TestBlock(t *testing.T) {
	cli, err := ethclient.Dial("https://sepolia-rollup.arbitrum.io/rpc")
	if err != nil {
		t.Fatal(err)
	}
	//block, err := cli.BlockByHash(context.Background(), common.HexToHash("0x37f01502c50a5cd9a95febd3bfb612e3026859589e645b5b3600cd7dd670fe9f"))
	/*block, err := cli.BlockByNumber(context.Background(), big.NewInt(106514310))
	if err != nil {
		t.Fatalf("Failed to get block: %v", err)
	}
	for _, tx := range block.Transactions() {
		switch tx.Type() {
		case 0x0: // Legacy transaction
			t.Log("LegacyTx:", tx)
		case 0x1: // AccessList transaction (EIP-2930)
			t.Log("AccessListTx:", tx)
		case 0x2: // DynamicFee transaction (EIP-1559)
			t.Log("DynamicFeeTx:", tx)
		default:
			t.Fatalf("Unsupported transaction type: %d", tx.Type())
		}
	}
	*/

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

/*
func TestEncodeCircuitInput(t *testing.T) {
	userOpDeposit := pool.UserOperation{
		OperationType:          1,
		OperationValue:         100,
		Sender:                 common.HexToAddress("0xd09d22e15b8c387a023811e5c1021b441b8f0e5a"),
		Nonce:                  1,
		ChainId:                hexutil.Uint64(hexTo("0xaa36a7").Uint64()),
		CallData:               common.FromHex("0x"),
		MainChainGasLimit:      0x30d40,
		DestChainGasLimit:      0,
		ZkVerificationGasLimit: 0x898,
		MainChainGasPrice:      hexTo("0x9502f900"),
		DestChainGasPrice:      hexTo("0x0"),
	}
	t.Logf("%+v", userOpDeposit)
	dataHash := userOpDeposit.CalculateEIP712TypeDataHash()
	signedUserOperation := pool.SignedUserOperation{
		UserOperation: &userOpDeposit,
		Signature:     signMessage(dataHash),
	}
	t.Logf("%+v", signedUserOperation)
	var sus []*pool.SignedUserOperation
	sus = append(sus, &signedUserOperation)
	by := encodeCircuitInput(sus)
	t.Log(by)
}

func TestDecodeCircuitInput(t *testing.T) {
	userOpDeposit := pool.UserOperation{
		OperationType:          1,
		OperationValue:         100,
		Sender:                 common.HexToAddress("0xd09d22e15b8c387a023811e5c1021b441b8f0e5a"),
		Nonce:                  1200,
		ChainId:                hexutil.Uint64(hexTo("0xaa36a7").Uint64()),
		CallData:               common.FromHex("0x"),
		MainChainGasLimit:      0x30d40,
		DestChainGasLimit:      0x30,
		ZkVerificationGasLimit: 0x898,
		MainChainGasPrice:      hexTo("0x9502f900"),
		DestChainGasPrice:      hexTo("0x90"),
	}
	dataHash := userOpDeposit.CalculateEIP712TypeDataHash()
	signedUserOperation := pool.SignedUserOperation{
		UserOperation: &userOpDeposit,
		Signature:     signMessage(dataHash),
	}
	t.Log(common.Bytes2Hex(signedUserOperation.Signature))
	t.Logf("%+v", signedUserOperation.UserOperation)
	var sus []*pool.SignedUserOperation
	sus = append(sus, &signedUserOperation)

	userOpDeposit2 := pool.UserOperation{
		OperationType:          2,
		OperationValue:         222,
		Sender:                 common.HexToAddress("0xd09d22e15b8c387a023811e5c1021b441b8f0e5a"),
		Nonce:                  1500,
		ChainId:                hexutil.Uint64(hexTo("0xaa36a7").Uint64()),
		CallData:               common.FromHex("0x"),
		MainChainGasLimit:      0x20d4,
		DestChainGasLimit:      0x130,
		ZkVerificationGasLimit: 0x89,
		MainChainGasPrice:      hexTo("0x9502f900"),
		DestChainGasPrice:      hexTo("0x190"),
	}
	dataHash2 := userOpDeposit2.CalculateEIP712TypeDataHash()
	signedUserOperation2 := pool.SignedUserOperation{
		UserOperation: &userOpDeposit2,
		Signature:     signMessage(dataHash2),
	}
	t.Log(common.Bytes2Hex(signedUserOperation2.Signature))
	t.Logf("%+v", signedUserOperation2.UserOperation)
	sus = append(sus, &signedUserOperation2)

	by := encodeCircuitInput(sus)
	sus = decodeCircuitInput(by)
	newUo := sus[0]
	t.Log(common.Bytes2Hex(newUo.Signature))
	t.Logf("%+v", newUo.UserOperation)

	newUo = sus[1]
	t.Log(common.Bytes2Hex(newUo.Signature))
	t.Logf("%+v", newUo.UserOperation)
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

func CreateUserOps() []*pool.SignedUserOperation {
	var sus []*pool.SignedUserOperation

	// userOp for deposit 0.2 ether
	userOpDeposit := pool.UserOperation{
		OperationType:          1,
		OperationValue:         0x2c68af0bb140000,
		Sender:                 common.HexToAddress("0xd09d22e15b8c387a023811e5c1021b441b8f0e5a"),
		Nonce:                  1,
		ChainId:                hexutil.Uint64(hexTo("0xaa36a7").Uint64()),
		CallData:               common.FromHex("0x"),
		MainChainGasLimit:      0x30d40,
		DestChainGasLimit:      0,
		ZkVerificationGasLimit: 0x898,
		MainChainGasPrice:      hexTo("0x9502f900"),
		DestChainGasPrice:      hexTo("0x0"),
	}

	{
		dataHash := userOpDeposit.CalculateEIP712TypeDataHash()

		// Sign
		signedUserOperation := pool.SignedUserOperation{
			UserOperation: &userOpDeposit,
			Signature:     signMessage(dataHash),
		}
		// verify signature
		recoveredAddr := signedUserOperation.RecoverAddress()
		fmt.Println(recoveredAddr.Hex())

		sus = append(sus, &signedUserOperation)
	}

	// userOp for withdraw 0.05 ether
	userOpWithdraw := pool.UserOperation{
		OperationType:          2,
		OperationValue:         0xb1a2bc2ec50000,
		Sender:                 common.HexToAddress("0xd09d22e15b8c387a023811e5c1021b441b8f0e5a"),
		Nonce:                  2,
		ChainId:                hexutil.Uint64(hexTo("0xaa36a7").Uint64()),
		CallData:               common.FromHex("0x"),
		MainChainGasLimit:      0x30d40,
		DestChainGasLimit:      0,
		ZkVerificationGasLimit: 0x898,
		MainChainGasPrice:      hexTo("0x9502f900"),
		DestChainGasPrice:      hexTo("0x0"),
	}

	{
		dataHash := userOpWithdraw.CalculateEIP712TypeDataHash()

		// Sign
		signedUserOperation := pool.SignedUserOperation{
			UserOperation: &userOpWithdraw,
			Signature:     signMessage(dataHash),
		}
		// verify signature
		recoveredAddr := signedUserOperation.RecoverAddress()
		fmt.Println(recoveredAddr.Hex())

		sus = append(sus, &signedUserOperation)
	}

	//for i := 0; i < 62; i++ {
	//	// userOp for counter contract
	//	userOpCounter := pool.UserOperation{
	//		OperationType:          0,
	//		Sender:                 common.HexToAddress("0xd09d22e15b8c387a023811e5c1021b441b8f0e5a"),
	//		Nonce:                  hexutil.Uint64(3 + i),
	//		ChainId:                hexTo("0xaa36a7"),
	//		CallData:               common.FromHex("0xb61d27f6000000000000000000000000c97e73b2770a0eb767407242fb3d35524fe229de000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000004d09de08a00000000000000000000000000000000000000000000000000000000"),
	//		MainChainGasLimit:      0x30d40,
	//		DestChainGasLimit:      0,
	//		ZkVerificationGasLimit: 0x41eb0,
	//		MainChainGasPrice:      hexTo("0x29810"),
	//		DestChainGasPrice:      hexTo("0x0"),
	//	}

	//	{
	//		dataHashCounter := userOpCounter.CalculateEIP712TypeDataHash()
	//
	//		// Sign
	//		signedUserOperationCounter := pool.SignedUserOperation{
	//			UserOperation: &userOpCounter,
	//			Signature:     signMessage(dataHashCounter),
	//		}
	//		// verify signature
	//		signedUserOperationCounter.RecoverAddress()
	//		//fmt.Println(recoveredAddr.Hex())
	//
	//		sus = append(sus, &signedUserOperationCounter)
	//	}
	//
	//}

	return sus
}*/
