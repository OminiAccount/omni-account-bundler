package smt

import (
	"encoding/json"
	"fmt"
	"github.com/OAB/database/leveldb"
	"github.com/OAB/utils/msgpack"
	"math/big"
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
		deltaProof := tree.SetLeaf(big.NewInt(int64(i)), leaf)
		proofJson, _ := json.MarshalIndent(deltaProof, "", "  ")
		fmt.Printf("Example 2: Setting leaf %d gives delta proof: %s\n", i, string(proofJson))
	}

	// Example 3: Verify a Merkle proof
	leaf := tree.GetLeaf(big.NewInt(3))
	isValid := VerifyMerkleProof(leaf)
	fmt.Println("Example 3: The proof is valid:", isValid)

	// Example 4: Verify a Delta Merkle proof
	isValidDelta := VerifyDeltaMerkleProof(tree.SetLeaf(big.NewInt(2), "0000000000000000000000000000000000000000000000000000000000000012"))
	fmt.Println("Example 4: The delta proof is valid:", isValidDelta)
}

func TestZeroMerkleTree_RollbackMerkleTree(t *testing.T) {
	type Proof map[uint64]DeltaMerkleProof

	type CTest struct {
		Account     uint64
		BatchNumber uint64
		StateProof  Proof
	}
	c := CTest{
		Account:     0,
		BatchNumber: 0,
		StateProof:  make(Proof),
	}

	number := 0
	//diffs := make(map[uint64]DeltaMerkleProof)
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
	tree := NewZeroMerkleTree(256)
	fmt.Printf("%d  %s\n", number, tree.GetRoot())
	// number is 0,tree don't have diff
	for i, leaf := range leavesToSet {
		c.StateProof[uint64(number)] = tree.SetLeaf(big.NewInt(int64(i)), leaf)
		//c.StateProof[uint64(number)].Index.SetUint64(12)
		//c.stateProof[uint64(number)] = deltaProof
		number++
		fmt.Printf("%d  %s\n", number, tree.GetRoot())
	}
	tree.RollbackMerkleTree(c.StateProof[0])
	fmt.Println("roll back to number 0", tree.GetRoot())

	dbs, _ := leveldb.NewLevelDB("./test_db")
	key := []byte("testkey")

	data, err := msgpack.MarshalStruct(c)
	if err != nil {
		fmt.Println("err", err)
	}
	dbs.Put(key, data)

	value, _ := dbs.Get(key)

	decodeData, err := msgpack.UnmarshalStruct[CTest](value)
	if err != nil {
		fmt.Println("err", err)
	}

	tree.RollbackMerkleTree(decodeData.StateProof[1])
	fmt.Println("roll back to number 1", tree.GetRoot())
}
