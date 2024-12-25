package pool

import (
	"errors"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/state"
	"github.com/OAB/utils/poseidon"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"time"
)

type Batch struct {
	OldStateRoot    common.Hash                  `json:"oldStateRoot"`
	NewStateRoot    common.Hash                  `json:"newStateRoot"`
	OldAccInputHash common.Hash                  `json:"oldAccInputHash"`
	NewAccInputHash common.Hash                  `json:"newAccInputHash"`
	OldNumBatch     uint64                       `json:"oldNumBatch"`
	NewNumBatch     uint64                       `json:"newNumBatch"`
	BatchL2Data     hexutil.Bytes                `json:"batchL2Data"`
	BatchHashData   common.Hash                  `json:"batchHashData"`
	BatchL2SrcData  []*state.SignedUserOperation `json:"batchL2SrcData"`
	Status          uint64                       `json:"status"`
	Timestamp       time.Time                    `json:"timestamp"`
}

func NewBatch(newNumBatch uint64) (*Batch, error) {
	oldNumBatch := uint64(0)
	if newNumBatch >= 1 {
		oldNumBatch = newNumBatch - 1
	} else {
		return nil, errors.New("newNumBatch is greater than 0")
	}
	batch := &Batch{
		OldNumBatch: oldNumBatch,
		NewNumBatch: newNumBatch,
		Timestamp:   time.Now(),
		Status:      state.PendingStatus,
	}
	return batch, nil
}

// encodeCircuitInput The first byte represents the number of UserOp (ff)
// [UserOpsLen,UserOps..]
// UserOp = [operationType,sender,chainId,callDataLen,callData,(mainChainGasLimit,destChainGasLimit),zkVerificationGasLimit,
//
//			 (mainChainGasPrice,destChainGasPrice),signature
//	]
//	callData and paymasterAndData are both dynamic bytes,so the field indicating their length is two bytes(ffff) ,
//	and the other fields are fixed 32 bytes.
func encodeCircuitInput(sus []*state.SignedUserOperation) []byte {
	var encodeBytes []byte
	encodeBytes = append(encodeBytes, byte(len(sus)))
	for _, su := range sus {
		suByte := su.Encode(false, true)
		suByte = append(suByte, su.Signature...)
		encodeBytes = append(encodeBytes, suByte...)
	}

	return encodeBytes
}

func (b *Batch) SetOldStateRoot(value common.Hash) {
	b.OldStateRoot = value
}

func (b *Batch) SetNewStateRoot(value common.Hash) {
	b.NewStateRoot = value
}

func (b *Batch) SetOldAccInputHash(value common.Hash) {
	b.OldAccInputHash = value
}

func (b *Batch) SetNewAccInputHash(sus []*state.SignedUserOperation) {
	b.NewAccInputHash = accInputData(sus)
}

func accInputData(sus []*state.SignedUserOperation) common.Hash {
	var encodeBytes []byte
	for _, su := range sus {
		encodeBytes = append(encodeBytes, su.Encode(true, false)...)
	}
	hashBytes, _ := poseidon.HashMessage(encodeBytes)
	return common.BytesToHash(poseidon.H4ToBytes(hashBytes))
}

func (b *Batch) SetBatchL2Data(sus []*state.SignedUserOperation) {
	b.BatchL2SrcData = sus
	b.BatchL2Data = encodeCircuitInput(sus)
	b.BatchHashData = crypto.Keccak256Hash(b.BatchL2Data)
}
