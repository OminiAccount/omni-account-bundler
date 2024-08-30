package pool

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type UserOperation struct {
	Sender               common.Address `json:"sender"`
	Nonce                hexutil.Uint64 `json:"nonce"`
	ChainId              hexutil.Big    `json:"chainId"`
	InitCode             hexutil.Bytes  `json:"initCode"`
	CallData             hexutil.Bytes  `json:"callData"`
	CallGasLimit         hexutil.Uint64 `json:"callGasLimit"`
	VerificationGasLimit hexutil.Uint64 `json:"verificationGasLimit"`
	PreVerificationGas   hexutil.Big    `json:"preVerificationGasLimit"`
	MaxFeePerGas         hexutil.Big    `json:"maxFeePerGas"`
	MaxPriorityFeePerGas hexutil.Big    `json:"maxPriorityFeePerGas"`
	PaymasterAndData     hexutil.Bytes  `json:"paymasterAndData"`
}

type SignedUserOperation struct {
	UserOperation
	Signature hexutil.Bytes `json:"signature"`
}
