package pool

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func hexTo(hex string) *hexutil.Big {
	bigInt := new(big.Int)
	bigInt, _ = bigInt.SetString(hex[2:], 16)
	return (*hexutil.Big)(bigInt)
}
func (p *Pool) mockPool() {
	fmt.Println("Mock user op")

	mockOp := SignedUserOperation{
		UserOperation: &UserOperation{
			Sender:               common.HexToAddress("0x38bdb2abd66c00cbf05584a9717c1094181a8780"),
			Nonce:                1,
			ChainId:              (*hexutil.Big)(big.NewInt(11155111)),
			InitCode:             common.FromHex(""),
			CallData:             common.FromHex("0xb61d27f600000000000000000000000001b7ca9d6b8ac943185e107e4be7430e5d90b5a5000000000000000000000000000000000000000000000000002386f26fc1000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
			CallGasLimit:         0x30d40,
			VerificationGasLimit: 0x41eb0,
			PreVerificationGas:   hexTo("0x29810"),
			MaxFeePerGas:         hexTo("0x6fc23ac00"),
			MaxPriorityFeePerGas: hexTo("0x77359400"),
			PaymasterAndData:     common.FromHex(""),
		},
		Signature: common.FromHex("0x2a33138b20e431c074197c721ca41d78634484a3c1717cc181d334074c98d44946930a3851026adbe647490524d6c369d2eb5879c2e297b6dd5befe2828a18071c"),
	}

	p.AddUserOp(mockOp)

	mockOp1 := SignedUserOperation{
		UserOperation: &UserOperation{
			Sender:               common.HexToAddress("0x38bdb2abd66c00cbf05584a9717c1094181a8780"),
			Nonce:                2,
			ChainId:              (*hexutil.Big)(big.NewInt(11155111)),
			InitCode:             common.FromHex(""),
			CallData:             common.FromHex("0xb61d27f6000000000000000000000000c97e73b2770a0eb767407242fb3d35524fe229de000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000004d09de08a00000000000000000000000000000000000000000000000000000000"),
			CallGasLimit:         0x30d40,
			VerificationGasLimit: 0x41eb0,
			PreVerificationGas:   hexTo("0x29810"),
			MaxFeePerGas:         hexTo("0x6fc23ac00"),
			MaxPriorityFeePerGas: hexTo("0x77359400"),
			PaymasterAndData:     common.FromHex(""),
		},
		Signature: common.FromHex("0x5df3d9f5bb5565321f674adcb0fd3a93dd940d4d21e059563c08b28d950398154a8c67733aef5da89a289722c4937c446a516d74d1796020907e9423e83e9fe81b"),
	}

	p.AddUserOp(mockOp1)
}
