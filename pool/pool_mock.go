package pool

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

import (
	"github.com/OAB/lib/common/hexutil"
)

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

func CreateUserOps() []*SignedUserOperation {
	var sus []*SignedUserOperation

	// userOp for deposit 0.2 ether
	userOpDeposit := UserOperation{
		OperationType:  1,
		OperationValue: (*hexutil.Big)(big.NewInt(100)),
		Sender:         common.HexToAddress("0xd09d22e15b8c387a023811e5c1021b441b8f0e5a"),
		Exec: ExecData{
			Nonce:                  1,
			ChainId:                hexutil.Uint64(hexTo("0xaa36a7").Uint64()),
			CallData:               common.FromHex("0x"),
			MainChainGasLimit:      0x30d40,
			DestChainGasLimit:      0,
			ZkVerificationGasLimit: 0x898,
			MainChainGasPrice:      hexTo("0x9502f900"),
			DestChainGasPrice:      hexTo("0x0"),
		},
	}

	{
		dataHash := userOpDeposit.CalculateEIP712TypeDataHash()

		// Sign
		signedUserOperation := SignedUserOperation{
			UserOperation: &userOpDeposit,
			Signature:     signMessage(dataHash),
		}
		// verify signature
		recoveredAddr := signedUserOperation.RecoverAddress()
		fmt.Println(recoveredAddr.Hex())

		sus = append(sus, &signedUserOperation)
	}

	// userOp for withdraw 0.05 ether
	userOpWithdraw := UserOperation{
		OperationType:  2,
		OperationValue: (*hexutil.Big)(big.NewInt(100)),
		Sender:         common.HexToAddress("0xd09d22e15b8c387a023811e5c1021b441b8f0e5a"),
		Exec: ExecData{
			Nonce:                  2,
			ChainId:                hexutil.Uint64(hexTo("0xaa36a7").Uint64()),
			CallData:               common.FromHex("0x"),
			MainChainGasLimit:      0x30d40,
			DestChainGasLimit:      0,
			ZkVerificationGasLimit: 0x898,
			MainChainGasPrice:      hexTo("0x9502f900"),
			DestChainGasPrice:      hexTo("0x0"),
		},
	}

	{
		dataHash := userOpWithdraw.CalculateEIP712TypeDataHash()

		// Sign
		signedUserOperation := SignedUserOperation{
			UserOperation: &userOpWithdraw,
			Signature:     signMessage(dataHash),
		}
		// verify signature
		recoveredAddr := signedUserOperation.RecoverAddress()
		fmt.Println(recoveredAddr.Hex())

		sus = append(sus, &signedUserOperation)
	}

	for i := 0; i < 62; i++ {
		// userOp for counter contract
		userOpCounter := UserOperation{
			OperationType:  0,
			OperationValue: (*hexutil.Big)(big.NewInt(0)),
			Sender:         common.HexToAddress("0xd09d22e15b8c387a023811e5c1021b441b8f0e5a"),
			Exec: ExecData{
				Nonce:                  hexutil.Uint64(3 + i),
				ChainId:                hexutil.Uint64(hexTo("0xaa36a7").Uint64()),
				CallData:               common.FromHex("0xb61d27f6000000000000000000000000c97e73b2770a0eb767407242fb3d35524fe229de000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000004d09de08a00000000000000000000000000000000000000000000000000000000"),
				MainChainGasLimit:      0x30d40,
				DestChainGasLimit:      0,
				ZkVerificationGasLimit: 0x898,
				MainChainGasPrice:      hexTo("0x9502f900"),
				DestChainGasPrice:      hexTo("0x0"),
			},
		}

		{
			dataHashCounter := userOpCounter.CalculateEIP712TypeDataHash()

			// Sign
			signedUserOperationCounter := SignedUserOperation{
				UserOperation: &userOpCounter,
				Signature:     signMessage(dataHashCounter),
			}
			// verify signature
			signedUserOperationCounter.RecoverAddress()
			//fmt.Println(recoveredAddr.Hex())

			sus = append(sus, &signedUserOperationCounter)
		}

	}

	return sus
}
