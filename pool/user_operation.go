package pool

import (
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
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

// Hash for the EIP712 Domain
func (domain EIP712Domain) hashEIP712Domain() common.Hash {
	return crypto.Keccak256Hash(
		crypto.Keccak256([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)")),
		crypto.Keccak256([]byte(domain.Name)),
		crypto.Keccak256([]byte(domain.Version)),
		big.NewInt(int64(domain.ChainId)).Bytes(),
		domain.VerifyingContract.Bytes(),
	)
}

// Hash the UserOperation
//func (userOp UserOperation) hashUserOperation() common.Hash {
//	return crypto.Keccak256Hash(
//		crypto.Keccak256([]byte("UserOperation(address sender,uint256 nonce,uint64 chainId,bytes initCode,bytes callData,uint256 callGasLimit,uint256 verificationGasLimit,uint256 preVerificationGas,uint256 maxFeePerGas,uint256 maxPriorityFeePerGas,bytes paymasterAndData)")),
//		userOp.Bytes(),
//	)
//}
//
//func (userOp UserOperation) Bytes() []byte {
//	return bytes.Join([][]byte{
//		userOp.Sender.Bytes(),
//		U64ToBytes(userOp.Nonce),
//		userOp.ChainId.ToInt().Bytes(),
//		crypto.Keccak256(userOp.InitCode),
//		crypto.Keccak256(userOp.CallData),
//		U64ToBytes(userOp.CallGasLimit),
//		U64ToBytes(userOp.VerificationGasLimit),
//		userOp.PreVerificationGas.ToInt().Bytes(),
//		userOp.MaxFeePerGas.ToInt().Bytes(),
//		userOp.MaxPriorityFeePerGas.ToInt().Bytes(),
//		crypto.Keccak256(userOp.PaymasterAndData),
//	}, nil)
//}

func U64ToBytes(u64 hexutil.Uint64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(u64))
	return bytes
}

func (su *SignedUserOperation) RecoverId() uint8 {
	v := su.Signature[64]
	return v
}

func MockSignedUserOperation(address string, privateKeyHex string, chainID uint64, nonce uint64) (*UserOperation, []byte, uint8, *ecdsa.PublicKey, error) {
	ethAddress := common.HexToAddress(address)
	userOperation := &UserOperation{
		Sender:               ethAddress,
		Nonce:                hexutil.Uint64(nonce),
		ChainId:              (*hexutil.Big)(big.NewInt(int64(chainID))),
		InitCode:             []byte{},
		CallData:             []byte{},
		CallGasLimit:         20000,
		VerificationGasLimit: 20000,
		PreVerificationGas:   (*hexutil.Big)(big.NewInt(10000)),
		MaxFeePerGas:         (*hexutil.Big)(big.NewInt(20000000000)),
		MaxPriorityFeePerGas: (*hexutil.Big)(big.NewInt(0)),
		PaymasterAndData:     []byte{},
	}

	domainSeparator := sha3.NewLegacyKeccak256()
	domainSeparator.Write([]byte{0x19, 0x01})

	domain := EIP712Domain{
		Name:              "ZK-AA",
		Version:           "1.0",
		ChainId:           chainID,
		VerifyingContract: ethAddress,
	}
	domainHash := hashStruct(domain)
	domainSeparator.Write(domainHash)

	userOpHash := hashStruct(userOperation)
	digest := sha3.NewLegacyKeccak256()
	digest.Write(domainSeparator.Sum(nil))
	digest.Write(userOpHash)

	digestBytes := digest.Sum(nil)

	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, nil, 0, nil, fmt.Errorf("invalid hex string: %v", err)
	}

	_, err = crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, nil, 0, nil, fmt.Errorf("invalid private key: %v", err)
	}

	sig, err := secp256k1.Sign(digestBytes, privateKeyBytes)
	if err != nil {
		return nil, nil, 0, nil, fmt.Errorf("failed to sign digest: %v", err)
	}

	recoveryID := sig[64]

	publicKey, err := crypto.SigToPub(digestBytes, sig)
	if err != nil {
		return nil, nil, 0, nil, fmt.Errorf("failed to recover public key: %v", err)
	}

	return userOperation, sig, recoveryID, publicKey, nil
}

func hashStruct(data interface{}) []byte {
	encoded, err := rlp.EncodeToBytes(data)
	if err != nil {
		fmt.Printf("failed to encode struct: %v\n", err)
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(encoded)
	return hash.Sum(nil)
}
