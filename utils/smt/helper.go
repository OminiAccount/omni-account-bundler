package smt

import (
	"math/big"
)

func ComputeNonceIndex(userAddress []byte, chainId uint64) *big.Int {
	nonceKey := ComputeNonceKey(userAddress, chainId)
	return KeyToIndex(nonceKey)
}

func ComputeBalanceIndex(userAddress []byte) *big.Int {
	balanceKey := ComputeBalanceKey(userAddress)
	return KeyToIndex(balanceKey)
}
