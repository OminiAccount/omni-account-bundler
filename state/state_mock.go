package state

import (
	"fmt"
	"github.com/OAAC/utils/smt"
	"github.com/ethereum/go-ethereum/common"
)

func (s *State) mockTree() {
	fmt.Println("Mock tree")
	sender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	balanceKey := smt.ComputeBalanceKey(common.HexToAddress(sender).Bytes())

	// init 1 ETH balance for mock user
	s.tree.SetLeaf(
		int(smt.KeyToIndex(balanceKey)),
		"0000000000000000000000000000000000000000000000000de0b6b3a7640000",
	)
	chain_id := 42161
	nonceKey := smt.ComputeNonceKey(common.HexToAddress(sender).Bytes(), uint64(chain_id))

	s.tree.SetLeaf(
		int(smt.KeyToIndex(nonceKey)),
		"0000000000000000000000000000000000000000000000000000000000000007",
	)
}
