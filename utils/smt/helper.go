package smt

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
)

func ComputeNonceIndex(userAddress []byte, chainId uint64) int {
	nonceKey := ComputeNonceKey(userAddress, chainId)
	return KeyToIndex(nonceKey)
}

func ComputeBalanceIndex(userAddress []byte) int {
	balanceKey := ComputeBalanceKey(userAddress)
	return KeyToIndex(balanceKey)
}

func IntToHexBigInt(value int) *hexutil.Big {
	bigInt, _ := new(big.Int).SetString(strconv.Itoa(value), 10)
	return (*hexutil.Big)(bigInt)
}

func HexBigIntToInt(value *hexutil.Big) {
}
