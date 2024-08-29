package smt

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/holiman/uint256"
	"testing"
)

func TestKey(t *testing.T) {
	userAddress := []byte{0xf3, 0x9f, 0xd6, 0xe5, 0x1a, 0xad, 0x88, 0xf6, 0xf4, 0xce, 0x6a, 0xb8, 0x82, 0x72, 0x79, 0xcf, 0xff, 0xb9, 0x22, 0x66}

	balance := Balance{Amount: uint256.MustFromHex(hexutil.EncodeUint64(1000))}
	nonce := Nonce{Value: 1001, ChainId: 42}

	balanceKey := computeLeafKey(userAddress, balance)
	fmt.Printf("Balance Key: %v\n", balanceKey)

	nonceKey := computeLeafKey(userAddress, nonce)
	fmt.Printf("Nonce Key: %v\n", nonceKey)

	balanceKey2 := computeBalanceKey(userAddress)
	fmt.Printf("Balance Key2: %v\n", balanceKey2)

	nonceKey2 := computeNonceKey(userAddress, 42)
	fmt.Printf("Nonce Key2: %v\n", nonceKey2)

	if balanceKey == balanceKey2 && nonceKey == nonceKey2 {
		fmt.Println("Keys match")
	} else {
		fmt.Println("Keys do not match")
	}
}
