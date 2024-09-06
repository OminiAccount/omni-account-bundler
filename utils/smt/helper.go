package smt

func ComputeNonceIndex(userAddress []byte, chainId uint64) int {
	nonceKey := ComputeNonceKey(userAddress, chainId)
	return KeyToIndex(nonceKey)
}

func ComputeBalanceIndex(userAddress []byte) int {
	balanceKey := ComputeBalanceKey(userAddress)
	return KeyToIndex(balanceKey)
}
