package smt

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSMT(t *testing.T) {

	// Example 1: Compute zero hashes
	zeroHashes := computeZeroHashes(32)
	zeroHashesJson, _ := json.MarshalIndent(zeroHashes, "", "  ")
	fmt.Println("Example 1: The first 32 zero hashes are:", string(zeroHashesJson))

	// Example 2: Build a Zero Merkle Tree
	leavesToSet := []MerkleNodeValue{
		"0000000000000000000000000000000000000000000000000000000000000001", // 1
		"0000000000000000000000000000000000000000000000000000000000000003", // 3
		"0000000000000000000000000000000000000000000000000000000000000003", // 3
		"0000000000000000000000000000000000000000000000000000000000000007", // 7
		"0000000000000000000000000000000000000000000000000000000000000004", // 4
		"0000000000000000000000000000000000000000000000000000000000000002", // 2
		"0000000000000000000000000000000000000000000000000000000000000000", // 0
		"0000000000000000000000000000000000000000000000000000000000000006", // 6
	}
	tree := NewZeroMerkleTree(3)
	for i, leaf := range leavesToSet {
		deltaProof := tree.SetLeaf(i, leaf)
		proofJson, _ := json.MarshalIndent(deltaProof, "", "  ")
		fmt.Printf("Example 2: Setting leaf %d gives delta proof: %s\n", i, string(proofJson))
	}

	// Example 3: Verify a Merkle proof
	leaf := tree.GetLeaf(3)
	isValid := VerifyMerkleProof(leaf)
	fmt.Println("Example 3: The proof is valid:", isValid)

	// Example 4: Verify a Delta Merkle proof
	isValidDelta := VerifyDeltaMerkleProof(tree.SetLeaf(2, "0000000000000000000000000000000000000000000000000000000000000012"))
	fmt.Println("Example 4: The delta proof is valid:", isValidDelta)
}
