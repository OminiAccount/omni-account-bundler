package smt

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/holiman/uint256"
	"math/big"
)

type LeafValue interface {
	Identifier() byte
}

type Balance struct {
	Amount *uint256.Int
}

func (b Balance) Identifier() byte {
	return 0
}

type Nonce struct {
	Value   uint64
	ChainId uint64
}

func (n Nonce) Identifier() byte {
	return 1
}

func computeLeafKey(userAddress []byte, value LeafValue) string {
	hasher := sha256.New()
	paddedInput := make([]byte, 32)

	// Padding the first byte based on the type identifier
	paddedInput[0] = value.Identifier()
	copy(paddedInput[1:], userAddress)

	if n, ok := value.(Nonce); ok {
		// If the type is Nonce, continue to fill the ChainId
		copy(paddedInput[21:], toBytes(n.ChainId))
	}

	hasher.Write(paddedInput)
	leafKey := hasher.Sum(nil)
	return hex.EncodeToString(leafKey)
}

func ComputeBalanceKey(userAddress []byte) string {
	balance := Balance{Amount: uint256.MustFromHex(hexutil.EncodeUint64(0))} // The Amount value here is not important in this case
	return computeLeafKey(userAddress, balance)
}

func ComputeNonceKey(userAddress []byte, chainId uint64) string {
	nonce := Nonce{Value: 0, ChainId: chainId}
	return computeLeafKey(userAddress, nonce)
}

// toBytes converts an uint64 to a big-endian byte array
func toBytes(val uint64) []byte {
	bytes := make([]byte, 8)
	for i := uint(0); i < 8; i++ {
		bytes[7-i] = byte(val >> (i * 8))
	}
	return bytes
}

func KeyToIndex(key string) *big.Int {
	truncatedStr := key[len(key)-64:]
	index, err := hex.DecodeString(truncatedStr)
	if err != nil {
		panic("Invalid hex string")
	}
	bigIndex := new(big.Int).SetBytes(index)

	return bigIndex
}
