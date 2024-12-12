package state

import (
	"errors"
	"github.com/OAB/pool"
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/poseidon"
	"github.com/OAB/lib/common/hexutil"
	"github.com/ethereum/go-ethereum/common"
)

const (
	opLength    = 290
	OpTypeE     = 1
	OpValueE    = 32
	SenderE     = 64
	OpInfoE     = 96
	CallDataE   = 128
	GasLimitE   = 160
	ZkGasLimitE = 192
	GasPriceE   = 224
	SignatureE  = 289
)

type Batch struct {
	OldStateRoot    common.Hash                 `json:"oldStateRoot"`
	NewStateRoot    common.Hash                 `json:"newStateRoot"`
	OldAccInputHash common.Hash                 `json:"oldAccInputHash"`
	NewAccInputHash common.Hash                 `json:"newAccInputHash"`
	OldNumBatch     uint64                      `json:"oldNumBatch"`
	NewNumBatch     uint64                      `json:"newNumBatch"`
	ChainID         uint64                      `json:"chainID"`
	ForkID          uint64                      `json:"forkID"`
	BatchL2Data     hexutil.Bytes               `json:"batchL2Data"`
	Timestamp       uint64                      `json:"timestamp"`
	BatchHashData   common.Hash                 `json:"batchHashData"`
	BatchL2SrcData  []*pool.SignedUserOperation `json:"batchL2SrcData"`
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
		suByte := su.Encode()
		suByte = append(suByte, su.Signature...)
		encodeBytes = append(encodeBytes, suByte...)
	}

	return encodeBytes
}

func decodeCircuitInput(val []byte) []*pool.SignedUserOperation {
	if len(val) < opLength {
		log.Error("decodeCircuitInput param length short")
		return nil
	}
	//l := val[0]
	var sus []*pool.SignedUserOperation
	/*for i := 0; i < len(val); {
		if len(val) < (i + opLength) {
			break
		}
		uo := make([]byte, opLength)
		copy(uo, val[(i+1):(i+opLength)])
		i += opLength - 1
		item := &pool.SignedUserOperation{
			UserOperation: &pool.UserOperation{},
			Signature:     uo[GasPriceE:SignatureE],
		}
		item.UserOperation.OperationType = packutils.BytesToUint8(uo[0:OpTypeE])
		item.UserOperation.OperationValue = hexutil.Uint64(packutils.BytesToUint64(uo[OpTypeE:OpValueE]))
		item.UserOperation.Sender = common.BytesToAddress(uo[OpValueE:SenderE])
		nonce, chainid, err := packutils.UnpackUints(uo[SenderE:OpInfoE])
		if err != nil {
			log.Errorf("decodeCircuitInput unpack uints error: %v", err)
		}
		item.UserOperation.Nonce = hexutil.Uint64(nonce.Uint64())
		padded := make([]byte, 16)
		copy(padded[16-len(chainid.Bytes()):], chainid.Bytes())
		var chex []hexutil.Uint64
		c1 := hexutil.Uint64(big.NewInt(0).SetBytes(padded[8:]).Uint64())
		if c1.Uint64() > 0 {
			chex = append(chex, c1)
		}
		c2 := hexutil.Uint64(big.NewInt(0).SetBytes(padded[:8]).Uint64())
		if c2.Uint64() > 0 {
			chex = append(chex, c2)
		}
		item.UserOperation.ChainId = chex
		item.UserOperation.CallData = uo[OpInfoE:CallDataE]
		mainGasLimit, destGasLimit, err := packutils.UnpackUints(uo[CallDataE:GasLimitE])
		if err != nil {
			log.Errorf("decodeCircuitInput unpack uints error: %v", err)
		}
		item.UserOperation.MainChainGasLimit = hexutil.Uint64(mainGasLimit.Uint64())
		item.UserOperation.DestChainGasLimit = hexutil.Uint64(destGasLimit.Uint64())
		item.UserOperation.ZkVerificationGasLimit = hexutil.Uint64(packutils.BytesToUint64(uo[GasLimitE:ZkGasLimitE]))
		mainGasPrice, destGasPrice, err := packutils.UnpackUints(uo[ZkGasLimitE:GasPriceE])
		if err != nil {
			log.Errorf("decodeCircuitInput unpack uints error: %v", err)
		}
		item.UserOperation.MainChainGasPrice = (*hexutil.Big)(mainGasPrice)
		item.UserOperation.DestChainGasPrice = (*hexutil.Big)(destGasPrice)
		sus = append(sus, item)
	}
	if len(sus) != int(l) {
		log.Warn("source data decode length mismatch")
	}*/
	return sus
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
	b.BatchL2SrcData = sus
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
	FinalNumber  uint64        `json:"finalNumber"`
	Proof        hexutil.Bytes `json:"proof"`
	PublicValues hexutil.Bytes `json:"public_values"`
}
