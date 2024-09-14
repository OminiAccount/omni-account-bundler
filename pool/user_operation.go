package pool

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

// EIP712Domain represents the domain object for EIP-712
type EIP712Domain struct {
	Name              string
	Version           string
	ChainId           uint64
	VerifyingContract common.Address
}

type UserOperation struct {
	Sender               common.Address `json:"sender"`
	Nonce                hexutil.Uint64 `json:"nonce"`
	ChainId              *hexutil.Big   `json:"chainId"`
	InitCode             hexutil.Bytes  `json:"initCode"`
	CallData             hexutil.Bytes  `json:"callData"`
	CallGasLimit         hexutil.Uint64 `json:"callGasLimit"`
	VerificationGasLimit hexutil.Uint64 `json:"verificationGasLimit"`
	PreVerificationGas   *hexutil.Big   `json:"preVerificationGasLimit"`
	MaxFeePerGas         *hexutil.Big   `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *hexutil.Big   `json:"maxPriorityFeePerGas"`
	PaymasterAndData     hexutil.Bytes  `json:"paymasterAndData"`
}

type SignedUserOperation struct {
	*UserOperation
	Signature hexutil.Bytes `json:"signature"`
}

// Recover the address from the signature
func recoverAddressFromSignature(signature []byte, messageHash common.Hash) (common.Address, error) {
	if len(signature) != 65 {
		return common.Address{}, fmt.Errorf("invalid signature length: expected 65 bytes, got %d", len(signature))
	}

	r := signature[:32]
	s := signature[32:64]
	v := signature[64]

	// Adjust v value for EIP-155 compliance
	if v < 27 {
		v += 27
	}

	// Recover the public key from the signature
	pubKey, err := crypto.Ecrecover(messageHash.Bytes(), append(r, append(s, v)...))
	if err != nil {
		return common.Address{}, err
	}

	// Get the address from the recovered public key
	pubKeyECDSA, err := crypto.UnmarshalPubkey(pubKey)
	if err != nil {
		return common.Address{}, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKeyECDSA)
	return recoveredAddr, nil
}

func (su *SignedUserOperation) RecoverId() uint8 {
	v := su.Signature[64]
	return v
}

// CalculateGasUsed Todo: Update to contract calculateGasUsed
func (u *UserOperation) CalculateGasUsed() *big.Int {
	var totalGas big.Int
	totalGas.Add(big.NewInt(int64(u.CallGasLimit)), big.NewInt(int64(u.VerificationGasLimit)))
	totalGas.Add(&totalGas, u.PreVerificationGas.ToInt())
	var totalGasCoeff big.Int
	totalGasCoeff.Add(u.MaxFeePerGas.ToInt(), u.MaxPriorityFeePerGas.ToInt())
	var gasUsed big.Int
	gasUsed.Mul(&totalGas, &totalGasCoeff)
	return &gasUsed
}
