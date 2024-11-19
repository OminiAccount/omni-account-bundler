package state

import (
	"errors"
	"github.com/OAB/pool"
	"github.com/OAB/utils/packutils"
	"github.com/OAB/utils/poseidon"

	//"github.com/OAB/utils/smt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Batch struct {
	OldStateRoot    common.Hash   `json:"oldStateRoot"`
	NewStateRoot    common.Hash   `json:"newStateRoot"`
	OldAccInputHash common.Hash   `json:"oldAccInputHash"`
	NewAccInputHash common.Hash   `json:"newAccInputHash"`
	OldNumBatch     uint64        `json:"oldNumBatch"`
	NewNumBatch     uint64        `json:"newNumBatch"`
	ChainID         uint64        `json:"chainID"`
	ForkID          uint64        `json:"forkID"`
	BatchL2Data     hexutil.Bytes `json:"batchL2Data"`
	Timestamp       uint64        `json:"timestamp"`
	BatchHashData   common.Hash   `json:"batchHashData"`
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
		Timestamp:   0,
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
func encodeCircuitInput(sus []*pool.SignedUserOperation) []byte {
	var encodeBytes []byte
	encodeBytes = append(encodeBytes, byte(len(sus)))
	for _, su := range sus {
		encodeBytes = append(encodeBytes, su.PackOperation()...)
		encodeBytes = append(encodeBytes, common.LeftPadBytes(su.Sender.Bytes(), 32)...)
		encodeBytes = append(encodeBytes, su.PackOpInfo()...)
		encodeBytes = append(encodeBytes, crypto.Keccak256(su.CallData)...)
		encodeBytes = append(encodeBytes, su.PackChainGasLimit()...)
		encodeBytes = append(encodeBytes, common.LeftPadBytes(packutils.Uint64ToBytes(uint64(su.ZkVerificationGasLimit)), 32)...)
		encodeBytes = append(encodeBytes, su.PackChainGasPrice()...)
		encodeBytes = append(encodeBytes, su.Signature...)
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

func (b *Batch) SetNewAccInputHash(value common.Hash) {
	b.NewAccInputHash = value
}

func (b *Batch) SetBatchL2Data(sus []*pool.SignedUserOperation) {
	b.BatchL2Data = encodeCircuitInput(sus)
	b.SetBatchHashData(sus)
}

func batchHashData(sus []*pool.SignedUserOperation) common.Hash {
	var encodeBytes []byte
	for _, su := range sus {
		encodeBytes = append(encodeBytes, su.Encode()...)
	}

	hashBytes, _ := poseidon.HashMessage(encodeBytes)

	return common.BytesToHash(poseidon.H4ToBytes(hashBytes))
}

func (b *Batch) SetBatchHashData(sus []*pool.SignedUserOperation) {
	b.BatchHashData = batchHashData(sus)
}

type ProofResult struct {
	Number       uint64        `json:"number"`
	Proof        hexutil.Bytes `json:"proof"`
	PublicValues hexutil.Bytes `json:"public_values"`
}
