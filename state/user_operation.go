package state

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
	"time"
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

	NormalStatus  = 0
	PendingStatus = 1
	BatchStatus   = 2
	SuccessStatus = 3
	FailedStatus  = 4
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
	Uid            uint64         `json:"uid"`
	OpId           uint64         `json:"op_id"`
	Did            string         `json:"did"`
	OperationType  uint8          `json:"operationType"`
	OperationValue *hexutil.Big   `json:"operationValue"`
	Owner          common.Address `json:"owner"`
	Sender         common.Address `json:"sender"`
	Exec           ExecData       `json:"exec"`
	InnerExec      ExecData       `json:"innerExec"`
	Phase          uint8          `json:"phase"`
	Status         uint8          `json:"status"`
	BatchNum       uint64         `json:"batch_num"`
	TimeAt         time.Time      `json:"time_at"`
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
	if u.InnerExec.MainChainGasPrice != nil {
		var sSrcGas big.Int
		sSrcGas.Add(big.NewInt(int64(u.InnerExec.MainChainGasLimit)), big.NewInt(int64(u.InnerExec.ZkVerificationGasLimit)))
		sSrcGas.Mul(&sSrcGas, u.InnerExec.MainChainGasPrice.ToInt())
		totalGas.Add(&totalGas, &sSrcGas)
	}
	if u.InnerExec.DestChainGasPrice != nil {
		var sDestGas big.Int
		sDestGas.Mul(big.NewInt(int64(u.InnerExec.DestChainGasLimit)), u.InnerExec.DestChainGasPrice.ToInt())
		totalGas.Add(&totalGas, &sDestGas)
	}
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
	if d.MainChainGasPrice == nil {
		return common.Hash{}.Bytes()
	}
	value, _ := packutils.PackUints(d.MainChainGasPrice.ToInt(), d.DestChainGasPrice.ToInt())
	return value
}

func (u *UserOperation) PackOperation() []byte {
	var encodeBytes []byte
	encodeBytes = append(encodeBytes, u.OperationType)
	if u.OperationValue == nil {
		encodeBytes = append(encodeBytes, common.LeftPadBytes([]byte{}, 31)...)
	} else {
		encodeBytes = append(encodeBytes, common.LeftPadBytes(u.OperationValue.ToInt().Bytes(), 31)...)
	}
	return encodeBytes
}

func (u *UserOperation) HasInnerExec() byte {
	if u.InnerExec.ChainId.Uint64() < 1 {
		return 0
	} else {
		return 1
	}
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
				{Name: "opInfo0", Type: "bytes32"},
				{Name: "callData0", Type: "bytes32"},
				{Name: "chainGasLimit0", Type: "bytes32"},
				{Name: "zkVerificationGasLimit0", Type: "uint256"},
				{Name: "chainGasPrice0", Type: "bytes32"},
				{Name: "opInfo1", Type: "bytes32"},
				{Name: "callData1", Type: "bytes32"},
				{Name: "chainGasLimit1", Type: "bytes32"},
				{Name: "zkVerificationGasLimit1", Type: "uint256"},
				{Name: "chainGasPrice1", Type: "bytes32"},
			},
		},
		PrimaryType: "UserOperation",
		Domain:      domain,
		Message: apitypes.TypedDataMessage{
			"operation":               hexutil.Encode(u.PackOperation()),
			"sender":                  u.Sender.String(),
			"opInfo0":                 hexutil.Encode(u.PackOpInfo(u.Exec)),
			"callData0":               crypto.Keccak256Hash(u.Exec.CallData).Hex(),
			"chainGasLimit0":          hexutil.Encode(u.PackChainGasLimit(u.Exec)),
			"zkVerificationGasLimit0": u.Exec.ZkVerificationGasLimit.String(),
			"chainGasPrice0":          hexutil.Encode(u.PackChainGasPrice(u.Exec)),
			"opInfo1":                 hexutil.Encode(u.PackOpInfo(u.InnerExec)),
			"callData1":               crypto.Keccak256Hash(u.InnerExec.CallData).Hex(),
			"chainGasLimit1":          hexutil.Encode(u.PackChainGasLimit(u.InnerExec)),
			"zkVerificationGasLimit1": u.InnerExec.ZkVerificationGasLimit.String(),
			"chainGasPrice1":          hexutil.Encode(u.PackChainGasPrice(u.InnerExec)),
		},
	}

	if false {
		// Domain hash
		domainSeparator, _ := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
		log.Info("domainSeparator:", domainSeparator)
		fmt.Println("domainSeparator:", domainSeparator)

		// Primary type hash
		primaryHash := typedData.TypeHash(typedData.PrimaryType)
		fmt.Println("primaryHash:", primaryHash)

		// Message Data
		encodedMessageData, _ := typedData.EncodeData(typedData.PrimaryType, typedData.Message, 1)
		fmt.Println("encodedMessage:", encodedMessageData)
		fmt.Println("encodedMessageHash:", crypto.Keccak256Hash(encodedMessageData))

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
		log.Errorf("signature must be %d bytes long", crypto.SignatureLength)
	}
	if signature[crypto.RecoveryIDOffset] != 27 && signature[crypto.RecoveryIDOffset] != 28 {
		log.Errorf("invalid Ethereum signature (V is not 27 or 28)")
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
func (s *SignedUserOperation) Encode(isPackOwner, isPackHasInner bool) []byte {
	var encodeBytes []byte
	encodeBytes = append(encodeBytes, s.PackOperation()...)
	encodeBytes = append(encodeBytes, common.LeftPadBytes(s.Sender.Bytes(), 32)...)
	if s.Exec.ChainId > 0 {
		encodeBytes = append(encodeBytes, s.PackOpInfo(s.Exec)...)
		encodeBytes = append(encodeBytes, crypto.Keccak256(s.Exec.CallData)...)
		encodeBytes = append(encodeBytes, s.PackChainGasLimit(s.Exec)...)
		encodeBytes = append(encodeBytes, common.LeftPadBytes(packutils.Uint64ToBytes(uint64(s.Exec.ZkVerificationGasLimit)), 32)...)
		encodeBytes = append(encodeBytes, s.PackChainGasPrice(s.Exec)...)
	}
	if isPackHasInner {
		encodeBytes = append(encodeBytes, []byte{s.HasInnerExec()}...)
	}
	if s.InnerExec.ChainId > 0 {
		encodeBytes = append(encodeBytes, s.PackOpInfo(s.InnerExec)...)
		encodeBytes = append(encodeBytes, crypto.Keccak256(s.InnerExec.CallData)...)
		encodeBytes = append(encodeBytes, s.PackChainGasLimit(s.InnerExec)...)
		encodeBytes = append(encodeBytes, common.LeftPadBytes(packutils.Uint64ToBytes(uint64(s.InnerExec.ZkVerificationGasLimit)), 32)...)
		encodeBytes = append(encodeBytes, s.PackChainGasPrice(s.InnerExec)...)
	}
	if isPackOwner {
		encodeBytes = append(encodeBytes, s.Owner.Bytes()...)
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
		encodeBytes = append(encodeBytes, su.Encode(true, true)...)
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
