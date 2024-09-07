package pool

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
	"testing"
)

//func TestPoolNew(t *testing.T) {
//	pool := NewMemoryPool(10, 5*time.Second)
//	pool.StartAutoFlush()
//
//	// Simulate adding operations
//	for i := 0; i < 25; i++ {
//		fmt.Println("i", i)
//		op := SignedUserOperation{
//			UserOperation: UserOperation{
//				Sender: common.HexToAddress(fmt.Sprintf("0x%040x", i)),
//			},
//			Signature: []byte(fmt.Sprintf("signature-%d", i)),
//		}
//		pool.AddUserOp(op)
//		time.Sleep(1 * time.Second)
//	}
//}

func TestSignedUserOperation_Recover(t *testing.T) {
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

	domainInfo := EIP712Domain{
		Name:              "OMNI-ACCOUNT",
		Version:           "1.0",
		ChainId:           11155111,
		VerifyingContract: common.HexToAddress("0x5F2464f924b7D9166a870cCe9201AFBC2a2f151D"),
	}

	chainId := math.HexOrDecimal256(*big.NewInt(int64(domainInfo.ChainId)))

	data := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"UserOperation": []apitypes.Type{
				{Name: "sender", Type: "address"},
				{Name: "nonce", Type: "uint256"},
				{Name: "chainId", Type: "uint64"},
				{Name: "initCode", Type: "bytes"},
				{Name: "callData", Type: "bytes"},
				{Name: "callGasLimit", Type: "uint256"},
				{Name: "verificationGasLimit", Type: "uint256"},
				{Name: "preVerificationGas", Type: "uint256"},
				{Name: "maxFeePerGas", Type: "uint256"},
				{Name: "maxPriorityFeePerGas", Type: "uint256"},
				{Name: "paymasterAndData", Type: "bytes"},
			},
		},
		Domain: apitypes.TypedDataDomain{
			Name:              domainInfo.Name,
			Version:           domainInfo.Version,
			ChainId:           &chainId,
			VerifyingContract: domainInfo.VerifyingContract.String(),
		},
		PrimaryType: "UserOperation",
		Message: apitypes.TypedDataMessage{
			"sender":               mockOp.Sender.Hex(),
			"nonce":                mockOp.Nonce.String(),
			"chainId":              mockOp.ChainId.String(),
			"initCode":             mockOp.InitCode.String(),
			"callData":             mockOp.CallData.String(),
			"callGasLimit":         mockOp.CallGasLimit.String(),
			"verificationGasLimit": mockOp.VerificationGasLimit.String(),
			"preVerificationGas":   mockOp.PreVerificationGas.String(),
			"maxFeePerGas":         mockOp.MaxFeePerGas.String(),
			"maxPriorityFeePerGas": mockOp.MaxPriorityFeePerGas.String(),
			"paymasterAndData":     mockOp.PaymasterAndData.String(),
		},
	}
	hash, _, err := apitypes.TypedDataAndHash(data)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("hash", common.BytesToHash(hash))
	address, _ := recoverAddressFromSignature(mockOp.Signature, common.BytesToHash(hash))
	fmt.Println("address", address)

}
