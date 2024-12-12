package pool

import (
	"fmt"
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/etherman/contracts/SyncRouter"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/utils/apitypes"
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/packutils"
	poseidon2 "github.com/OAB/utils/poseidon"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

const (
	EIP712DomainName    = "OMNI-ACCOUNT"
	EIP712DomainVersion = "1.0"

	UserAction      = 0
	DepositAction   = 1
	WithdrawAction  = 2
	SwapCrossAction = 3
	SwapAction      = 4
	TrueMarketBuy   = 5
	TrueMarketSell  = 6

	PhaseFirst  = 0
	PhaseSecond = 1
)

type ExecData struct {
	Nonce                  hexutil.Uint64 `json:"nonce"`
	ChainId                hexutil.Uint64 `json:"chainId"`
	CallData               hexutil.Bytes  `json:"callData"`
	MainChainGasLimit      hexutil.Uint64 `json:"mainChainGasLimit"`
	DestChainGasLimit      hexutil.Uint64 `json:"destChainGasLimit"`
	ZkVerificationGasLimit hexutil.Uint64 `json:"zkVerificationGasLimit"`
	MainChainGasPrice      *hexutil.Big   `json:"mainChainGasPrice"`
	DestChainGasPrice      *hexutil.Big   `json:"destChainGasPrice"`
}

type UserOperation struct {
	Did            string         `json:"did"`
	OperationType  uint8          `json:"operationType"`
	OperationValue *hexutil.Big   `json:"operationValue"`
	Owner          common.Address `json:"owner"`
	Sender         common.Address `json:"sender"`
	Exec           ExecData       `json:"exec"`
	InnerExec      ExecData       `json:"inner_exec"`
	Phase          uint8          `json:"destChainGasPrice"`
}

type SignedUserOperation struct {
	*UserOperation
	Signature hexutil.Bytes `json:"signature"`
}

func (u *UserOperation) CalculateGasUsed() *big.Int {
	var totalGas big.Int
	totalGas.Add(big.NewInt(int64(u.Exec.MainChainGasLimit)), big.NewInt(int64(u.Exec.ZkVerificationGasLimit)))
	totalGas.Mul(&totalGas, u.Exec.MainChainGasPrice.ToInt())
	var destGas big.Int
	destGas.Mul(big.NewInt(int64(u.Exec.DestChainGasLimit)), u.Exec.DestChainGasPrice.ToInt())
	totalGas.Add(&totalGas, &destGas)
	var sSrcGas big.Int
	sSrcGas.Add(big.NewInt(int64(u.InnerExec.MainChainGasLimit)), big.NewInt(int64(u.InnerExec.ZkVerificationGasLimit)))
	sSrcGas.Mul(&sSrcGas, u.InnerExec.MainChainGasPrice.ToInt())
	totalGas.Add(&totalGas, &sSrcGas)
	var sDestGas big.Int
	sDestGas.Mul(big.NewInt(int64(u.InnerExec.DestChainGasLimit)), u.InnerExec.DestChainGasPrice.ToInt())
	totalGas.Add(&totalGas, &sDestGas)

	return &totalGas
}

func (u *UserOperation) IsGasOperation() bool {
	return u.OperationType == DepositAction || u.OperationType == WithdrawAction
}

// PackOpInfo the `Nonce` and `ChainId`
func (u *UserOperation) PackOpInfo(d ExecData) []byte {
	value, _ := packutils.PackUints(big.NewInt(int64(d.Nonce)), big.NewInt(int64(d.ChainId)))
	return value
}

func (u *UserOperation) PackChainGasLimit(d ExecData) []byte {
	value, _ := packutils.PackUints(big.NewInt(int64(d.MainChainGasLimit)), big.NewInt(int64(d.DestChainGasLimit)))
	return value
}

func (u *UserOperation) PackChainGasPrice(d ExecData) []byte {
	value, _ := packutils.PackUints(d.MainChainGasPrice.ToInt(), d.DestChainGasPrice.ToInt())
	return value
}

func (u *UserOperation) PackOperation() []byte {
	var encodeBytes []byte
	encodeBytes = append(encodeBytes, u.OperationType)
	encodeBytes = append(encodeBytes, common.LeftPadBytes(u.OperationValue.ToInt().Bytes(), 248)...)
	return encodeBytes
}

func (u *UserOperation) CalculateEIP712TypeDataHash() []byte {
	domain := apitypes.TypedDataDomain{
		Name:    EIP712DomainName,
		Version: EIP712DomainVersion,
	}

	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
			},
			"UserOperation": []apitypes.Type{
				{Name: "operation", Type: "bytes32"},
				{Name: "sender", Type: "address"},
				{Name: "opInfo", Type: "bytes32"},
				{Name: "callData", Type: "bytes32"},
				{Name: "chainGasLimit", Type: "bytes32"},
				{Name: "zkVerificationGasLimit", Type: "uint256"},
				{Name: "chainGasPrice", Type: "bytes32"},
			},
		},
		PrimaryType: "UserOperation",
		Domain:      domain,
		Message: apitypes.TypedDataMessage{
			"operation": hexutil.Encode(u.PackOperation()),
			"sender":    u.Sender.String(),
			//"opInfo":                 hexutil.Encode(u.PackOpInfo()),
			//"callData":               crypto.Keccak256Hash(u.CallData).Hex(),
			//"chainGasLimit":          hexutil.Encode(u.PackChainGasLimit()),
			//"zkVerificationGasLimit": u.ZkVerificationGasLimit.String(),
			//"chainGasPrice":          hexutil.Encode(u.PackChainGasPrice()),
		},
	}

	if false {
		// Domain hash
		domainSeparator, _ := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
		log.Info("domainSeparator:", domainSeparator)

		// Primary type hash
		primaryHash := typedData.TypeHash(typedData.PrimaryType)
		log.Info("primaryHash:", primaryHash)

		// Message Data
		encodedMessageData, _ := typedData.EncodeData(typedData.PrimaryType, typedData.Message, 1)
		log.Info("encodedMessage:", encodedMessageData)
		log.Info("encodedMessageHash:", crypto.Keccak256Hash(encodedMessageData))

	}

	// Data hash = poseidon(0x1901 + DomainHash + MessageHash(poseidon(PrimaryHash+ContextHash))
	dataHash, _, err := apitypes.TypedDataAndHash(typedData)
	if err != nil {
		log.Errorf("Failed to hash typed data: %v", err)
	}
	return dataHash
}

// RecoverAddress Returns address from the signature
func (s *SignedUserOperation) RecoverAddress() common.Address {
	dataHash := s.CalculateEIP712TypeDataHash()

	messageHash := accounts.TextHash(dataHash)

	signature := common.CopyBytes(s.Signature)
	if len(signature) != crypto.SignatureLength {
		fmt.Printf("signature must be %d bytes long", crypto.SignatureLength)
	}
	if signature[crypto.RecoveryIDOffset] != 27 && signature[crypto.RecoveryIDOffset] != 28 {
		fmt.Printf("invalid Ethereum signature (V is not 27 or 28)")
	}
	signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	sigPublicKey, err := crypto.SigToPub(messageHash, signature)
	if err != nil {
		fmt.Printf("Failed to recover public key from signature: %v", err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*sigPublicKey)

	s.Owner = recoveredAddr

	return recoveredAddr
}

// Encode Returns the bytes value of the EIP712 value + recover address
func (s *SignedUserOperation) Encode() []byte {
	var encodeBytes []byte
	encodeBytes = append(encodeBytes, s.PackOperation()...)
	encodeBytes = append(encodeBytes, common.LeftPadBytes(s.Sender.Bytes(), 32)...)
	encodeBytes = append(encodeBytes, common.LeftPadBytes(s.Owner.Bytes(), 32)...)
	if s.Exec.ChainId > 0 {
		encodeBytes = append(encodeBytes, s.PackOpInfo(s.Exec)...)
		encodeBytes = append(encodeBytes, crypto.Keccak256(s.Exec.CallData)...)
		encodeBytes = append(encodeBytes, s.PackChainGasLimit(s.Exec)...)
		encodeBytes = append(encodeBytes, common.LeftPadBytes(packutils.Uint64ToBytes(uint64(s.Exec.ZkVerificationGasLimit)), 32)...)
		encodeBytes = append(encodeBytes, s.PackChainGasPrice(s.Exec)...)
	}
	if s.InnerExec.ChainId > 0 {
		encodeBytes = append(encodeBytes, s.PackOpInfo(s.InnerExec)...)
		encodeBytes = append(encodeBytes, crypto.Keccak256(s.InnerExec.CallData)...)
		encodeBytes = append(encodeBytes, s.PackChainGasLimit(s.InnerExec)...)
		encodeBytes = append(encodeBytes, common.LeftPadBytes(packutils.Uint64ToBytes(uint64(s.InnerExec.ZkVerificationGasLimit)), 32)...)
		encodeBytes = append(encodeBytes, s.PackChainGasPrice(s.InnerExec)...)
	}
	return encodeBytes
}

// EncodeEip712Context The first byte represents the number of UserOp (ff)
// [UserOpsLen,UserOps..]
// UserOp = [operationType,sender,chainId,callDataLen,callData,(mainChainGasLimit,destChainGasLimit),zkVerificationGasLimit,
//
//			 (mainChainGasPrice,destChainGasPrice),signature
//	]
//	callData and paymasterAndData are both dynamic bytes,so the field indicating their length is two bytes(ffff) ,
//	and the other fields are fixed 32 bytes.
func EncodeEip712Context(sus []*SignedUserOperation) []byte {
	var encodeBytes []byte
	encodeBytes = append(encodeBytes, byte(len(sus)))
	/*for _, su := range sus {
		encodeBytes = append(encodeBytes, su.PackOperation()...)
		encodeBytes = append(encodeBytes, common.LeftPadBytes(su.Sender.Bytes(), 32)...)
		encodeBytes = append(encodeBytes, su.PackOpInfo()...)
		encodeBytes = append(encodeBytes, crypto.Keccak256(su.CallData)...)
		encodeBytes = append(encodeBytes, su.PackChainGasLimit()...)
		encodeBytes = append(encodeBytes, common.LeftPadBytes(packutils.Uint64ToBytes(uint64(su.ZkVerificationGasLimit)), 32)...)
		encodeBytes = append(encodeBytes, su.PackChainGasPrice()...)
		encodeBytes = append(encodeBytes, su.Signature...)
	}*/

	return encodeBytes
}

func HashSignedUserOperationV1s(sus []*SignedUserOperation) common.Hash {
	var encodeBytes []byte
	for _, su := range sus {
		encodeBytes = append(encodeBytes, su.Encode()...)
	}

	hashBytes, _ := poseidon2.HashMessage(encodeBytes)

	return common.BytesToHash(poseidon2.H4ToBytes(hashBytes))
}

func (s *SignedUserOperation) ToEntryPointOp() EntryPoint.BaseStructPackedUserOperation {
	return EntryPoint.BaseStructPackedUserOperation{
		Phase:          s.Phase,
		OperationType:  s.OperationType,
		OperationValue: s.OperationValue.ToInt(),
		Sender:         s.Sender,
		Owner:          s.Owner,
		Exec: EntryPoint.BaseStructExecData{
			Nonce:                  uint64(s.Exec.Nonce),
			ChainId:                s.Exec.ChainId.Uint64(),
			CallData:               s.Exec.CallData,
			MainChainGasLimit:      uint64(s.Exec.MainChainGasLimit),
			MainChainGasPrice:      s.Exec.MainChainGasPrice.Uint64(),
			DestChainGasLimit:      uint64(s.Exec.DestChainGasLimit),
			DestChainGasPrice:      s.Exec.DestChainGasPrice.Uint64(),
			ZkVerificationGasLimit: uint64(s.Exec.ZkVerificationGasLimit),
		},
		InnerExec: EntryPoint.BaseStructExecData{
			Nonce:                  uint64(s.InnerExec.Nonce),
			ChainId:                s.InnerExec.ChainId.Uint64(),
			CallData:               s.InnerExec.CallData,
			MainChainGasLimit:      uint64(s.InnerExec.MainChainGasLimit),
			MainChainGasPrice:      s.InnerExec.MainChainGasPrice.Uint64(),
			DestChainGasLimit:      uint64(s.InnerExec.DestChainGasLimit),
			DestChainGasPrice:      s.InnerExec.DestChainGasPrice.Uint64(),
			ZkVerificationGasLimit: uint64(s.InnerExec.ZkVerificationGasLimit),
		},
	}
}

func (s *SignedUserOperation) ToSyncRouterOp() SyncRouter.BaseStructPackedUserOperation {
	return SyncRouter.BaseStructPackedUserOperation{
		Phase:          s.Phase,
		OperationType:  s.OperationType,
		OperationValue: s.OperationValue.ToInt(),
		Sender:         s.Sender,
		Owner:          s.Owner,
		Exec: SyncRouter.BaseStructExecData{
			Nonce:                  uint64(s.Exec.Nonce),
			ChainId:                s.Exec.ChainId.Uint64(),
			CallData:               s.Exec.CallData,
			MainChainGasLimit:      uint64(s.Exec.MainChainGasLimit),
			MainChainGasPrice:      s.Exec.MainChainGasPrice.Uint64(),
			DestChainGasLimit:      uint64(s.Exec.DestChainGasLimit),
			DestChainGasPrice:      s.Exec.DestChainGasPrice.Uint64(),
			ZkVerificationGasLimit: uint64(s.Exec.ZkVerificationGasLimit),
		},
		InnerExec: SyncRouter.BaseStructExecData{
			Nonce:                  uint64(s.InnerExec.Nonce),
			ChainId:                s.InnerExec.ChainId.Uint64(),
			CallData:               s.InnerExec.CallData,
			MainChainGasLimit:      uint64(s.InnerExec.MainChainGasLimit),
			MainChainGasPrice:      s.InnerExec.MainChainGasPrice.Uint64(),
			DestChainGasLimit:      uint64(s.InnerExec.DestChainGasLimit),
			DestChainGasPrice:      s.InnerExec.DestChainGasPrice.Uint64(),
			ZkVerificationGasLimit: uint64(s.InnerExec.ZkVerificationGasLimit),
		},
	}
}
