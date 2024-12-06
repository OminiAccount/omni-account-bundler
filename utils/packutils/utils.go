package packutils

import (
	"encoding/binary"
	"fmt"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func Uint16ToBytes(num uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, num)
	return buf
}

func BytesToUint16(num []byte) uint16 {
	if len(num) < 2 {
		num = common.LeftPadBytes(num, 2)
	}
	return binary.BigEndian.Uint16(num[len(num)-2:])
}

func Uint64ToBytes(num uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, num)
	return buf
}

func BytesToUint64(num []byte) uint64 {
	if len(num) < 8 {
		num = common.LeftPadBytes(num, 8)
	}
	return binary.BigEndian.Uint64(num[len(num)-8:])
}

func Uint8ToBytes(num uint8) []byte {
	return []byte{num}
}

func BytesToUint8(num []byte) uint8 {
	if len(num) < 1 {
		log.Error("byte slice is too short to convert to uint8")
		return 0
	}
	return num[0]
}

// PackUints packs two uint128 values (represented by *big.Int) into a 32-byte array.
func PackUints(high128, low128 *big.Int) ([]byte, error) {
	if high128.BitLen() > 128 || low128.BitLen() > 128 {
		return nil, fmt.Errorf("high128 or low128 exceeds 128 bits")
	}

	packed := make([]byte, 32)

	// Convert the low128 and high128 to 16-byte representations
	lowBytes := low128.FillBytes(make([]byte, 16))
	highBytes := high128.FillBytes(make([]byte, 16))

	// Copy high128 into the first 16 bytes of the packed array
	copy(packed[:16], highBytes)

	// Copy low128 into the last 16 bytes of the packed array
	copy(packed[16:], lowBytes)

	return packed, nil
}

func UnpackUints(val []byte) (*big.Int, *big.Int, error) {
	if len(val) < 32 {
		return nil, nil, fmt.Errorf("bytes length is too short")
	}

	highBytes := val[:16]
	lowBytes := val[16:]
	highInt := big.NewInt(0).SetBytes(highBytes)
	lowInt := big.NewInt(0).SetBytes(lowBytes)
	return highInt, lowInt, nil
}
