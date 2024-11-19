package poseidon

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	hex2 "github.com/OAB/utils/hex"
	"github.com/ethereum/go-ethereum/common"
	poseidon "github.com/iden3/go-iden3-crypto/goldenposeidon"
	"math"
	"math/big"
	"strings"
)

// maxBigIntLen is 256 bits (32 bytes)
const maxBigIntLen = 32

// wordLength is the number of bits of each ff limb
const wordLength = 64

// h4ToScalar converts array of 4 uint64 into a unique 256 bits scalar.
func h4ToScalar(h4 []uint64) *big.Int {
	if len(h4) == 0 {
		return new(big.Int)
	}
	result := new(big.Int).SetUint64(h4[0])

	for i := 1; i < 4; i++ {
		b2 := new(big.Int).SetUint64(h4[i])
		b2.Lsh(b2, uint(wordLength*i))
		result = result.Add(result, b2)
	}

	return result
}

func H4ToScalar(h4 []uint64) *big.Int {
	return h4ToScalar(h4)
}

// H4ToString converts array of 4 Scalars of 64 bits into a hex string.
func H4ToString(h4 []uint64) string {
	sc := h4ToScalar(h4)

	return fmt.Sprintf("0x%064s", hex.EncodeToString(sc.Bytes()))
}

// H4ToBytes converts array of 4 Scalars of 64 bits into a hex bytes.
func H4ToBytes(h4 []uint64) []byte {
	dst := H4ToString(h4)

	return common.FromHex(dst)
}

// scalar2fea splits a *big.Int into array of 32bit uint64 values.
func scalar2fea(value *big.Int) []uint64 {
	val := make([]uint64, 8)                          //nolint:gomnd
	mask, _ := new(big.Int).SetString("FFFFFFFF", 16) //nolint:gomnd
	val[0] = new(big.Int).And(value, mask).Uint64()
	val[1] = new(big.Int).And(new(big.Int).Rsh(value, 32), mask).Uint64()  //nolint:gomnd
	val[2] = new(big.Int).And(new(big.Int).Rsh(value, 64), mask).Uint64()  //nolint:gomnd
	val[3] = new(big.Int).And(new(big.Int).Rsh(value, 96), mask).Uint64()  //nolint:gomnd
	val[4] = new(big.Int).And(new(big.Int).Rsh(value, 128), mask).Uint64() //nolint:gomnd
	val[5] = new(big.Int).And(new(big.Int).Rsh(value, 160), mask).Uint64() //nolint:gomnd
	val[6] = new(big.Int).And(new(big.Int).Rsh(value, 192), mask).Uint64() //nolint:gomnd
	val[7] = new(big.Int).And(new(big.Int).Rsh(value, 224), mask).Uint64() //nolint:gomnd
	return val
}

func Scalar2fea(value *big.Int) []uint64 {
	return scalar2fea(value)
}

// StringToh4 converts a hex string into array of 4 Scalars of 64 bits.
func StringToh4(str string) ([]uint64, error) {
	if strings.HasPrefix(str, "0x") { // nolint
		str = str[2:]
	}

	bi, ok := new(big.Int).SetString(str, hex2.Base)
	if !ok {
		return nil, fmt.Errorf("could not convert %q into big int", str)
	}

	return scalarToh4(bi), nil
}

// scalarToh4 converts a *big.Int into an array of 4 uint64
func scalarToh4(s *big.Int) []uint64 {
	b := ScalarToFilledByteSlice(s)

	r := make([]uint64, 4) //nolint:gomnd

	f, _ := hex2.DecodeHex("0xFFFFFFFFFFFFFFFF")
	fbe := binary.BigEndian.Uint64(f)

	r[3] = binary.BigEndian.Uint64(b[0:8]) & fbe
	r[2] = binary.BigEndian.Uint64(b[8:16]) & fbe
	r[1] = binary.BigEndian.Uint64(b[16:24]) & fbe
	r[0] = binary.BigEndian.Uint64(b[24:]) & fbe

	return r
}

func ScalarToh4(s *big.Int) []uint64 {
	return scalarToh4(s)
}

// ScalarToFilledByteSlice converts a *big.Int into an array of maxBigIntLen
// bytes.
func ScalarToFilledByteSlice(s *big.Int) []byte {
	buf := make([]byte, maxBigIntLen)
	return s.FillBytes(buf)
}

// HashMessage computes the message hash.
func HashMessage(code []byte) ([]uint64, error) {
	const (
		bytecodeElementsHash = 8
		bytecodeBytesElement = 7

		maxBytesToAdd = bytecodeElementsHash * bytecodeBytesElement
	)

	// add 0x01
	code = append(code, 0x01) // nolint:gomnd

	// add padding
	for len(code)%(56) != 0 { // nolint:gomnd
		code = append(code, 0x00) // nolint:gomnd
	}

	code[len(code)-1] = code[len(code)-1] | 0x80 // nolint:gomnd

	numHashes := int(math.Ceil(float64(len(code)) / float64(maxBytesToAdd)))

	tmpHash := [4]uint64{}
	var err error

	bytesPointer := 0
	for i := 0; i < numHashes; i++ {
		elementsToHash := [12]uint64{}

		for j := 0; j < 4; j++ {
			elementsToHash[j] = tmpHash[j]
		}

		subsetBytecode := code[bytesPointer : int(math.Min(float64(len(code)-1), float64(bytesPointer+maxBytesToAdd)))+1]
		bytesPointer += maxBytesToAdd
		tmpElem := [7]byte{}
		counter := 0
		index := 4
		for j := 0; j < maxBytesToAdd; j++ {
			byteToAdd := []byte{0}

			if j < len(subsetBytecode) {
				byteToAdd = subsetBytecode[j : j+1]
			}

			tmpElem[bytecodeBytesElement-1-counter] = byteToAdd[0]
			counter++

			if counter == bytecodeBytesElement {
				elementsToHash[index] = new(big.Int).SetBytes(tmpElem[:]).Uint64()
				index++
				tmpElem = [7]byte{}
				counter = 0
			}
		}
		tmpHash, err = poseidon.Hash([8]uint64{
			elementsToHash[4],
			elementsToHash[5],
			elementsToHash[6],
			elementsToHash[7],
			elementsToHash[8],
			elementsToHash[9],
			elementsToHash[10],
			elementsToHash[11],
		}, [4]uint64{
			elementsToHash[0],
			elementsToHash[1],
			elementsToHash[2],
			elementsToHash[3],
		})
		if err != nil {
			return nil, err
		}
	}
	return tmpHash[:], nil
}
