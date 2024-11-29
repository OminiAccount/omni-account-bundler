package packutils

import (
	"encoding/binary"
	"fmt"
	"math/big"
)

func Uint16ToBytes(num uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, num)
	return buf
}

func Uint64ToBytes(num uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, num)
	return buf
}

func BytesToUint64(num []byte) uint64 {
	if len(num) < 8 {
		panic("byte slice is too short to convert to uint64")
	}
	return binary.BigEndian.Uint64(num[:8])
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
