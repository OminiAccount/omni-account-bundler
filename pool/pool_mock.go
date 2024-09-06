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
			Sender:               common.HexToAddress("0x93d53d2d8f0d623c5cbe46daa818177a450bd9f7"),
			Nonce:                1,
			ChainId:              (*hexutil.Big)(big.NewInt(11155111)),
			InitCode:             common.FromHex(""),
			CallData:             common.FromHex("0xb61d27f600000000000000000000000027916984c665f15041929b68451303136fa16653000000000000000000000000000000000000000000000000002386f26fc1000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
			CallGasLimit:         0x30d40,
			VerificationGasLimit: 0x41eb0,
			PreVerificationGas:   hexTo("0x29810"),
			MaxFeePerGas:         hexTo("0x6fc23ac00"),
			MaxPriorityFeePerGas: hexTo("0x77359400"),
			PaymasterAndData:     common.FromHex(""),
		},
		Signature: common.FromHex("0xc5ffeb27a1232cd0f25d69294de67a975a08ca7a34d8ed1fb8769c64f423645e7e6a07e1ad4bd91f5b8b9512aa016bb40deca600de01db8aadda3401c03548aa1c"),
	}

	fmt.Println("op", mockOp.Signature)

	p.AddUserOp(mockOp)

	mockOp1 := SignedUserOperation{
		UserOperation: &UserOperation{
			Sender:               common.HexToAddress("0x93d53d2d8f0d623c5cbe46daa818177a450bd9f7"),
			Nonce:                2,
			ChainId:              (*hexutil.Big)(big.NewInt(11155111)),
			InitCode:             common.FromHex(""),
			CallData:             common.FromHex("0xb61d27f600000000000000000000000027916984c665f15041929b68451303136fa16653000000000000000000000000000000000000000000000000002386f26fc1000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
			CallGasLimit:         0x30d40,
			VerificationGasLimit: 0x41eb0,
			PreVerificationGas:   hexTo("0x29810"),
			MaxFeePerGas:         hexTo("0x6fc23ac00"),
			MaxPriorityFeePerGas: hexTo("0x77359400"),
			PaymasterAndData:     common.FromHex(""),
		},
		Signature: common.FromHex("0xeadb0145d599f3f88df32cf3deed7e6ac039d94f327a3b4c630e4808d9074b9527af31056bf0df3bcb4b013606af0aa094725e3b0e5fd032b4a1fa49a20e60421b"),
	}

	p.AddUserOp(mockOp1)
}
