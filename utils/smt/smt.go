package smt

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// MerkleNodeValue represents the value of each node in the Merkle tree
type MerkleNodeValue string

// MerkleProof represents a Merkle proof
type MerkleProof struct {
	Root     MerkleNodeValue   `json:"root"`
	Siblings []MerkleNodeValue `json:"siblings"`
	Index    int               `json:"index"`
	Value    MerkleNodeValue   `json:"value"`
}

// DeltaMerkleProof represents a proof of the state transition in the Merkle tree
type DeltaMerkleProof struct {
	Index    int               `json:"index"`
	Siblings []MerkleNodeValue `json:"siblings"`
	OldRoot  MerkleNodeValue   `json:"old_root"`
	OldValue MerkleNodeValue   `json:"old_value"`
	NewRoot  MerkleNodeValue   `json:"new_root"`
	NewValue MerkleNodeValue   `json:"new_value"`
}

// hash computes the SHA-256 hash of two Merkle node values
func hash(leftNode, rightNode MerkleNodeValue) MerkleNodeValue {
	hasher := sha256.New()
	leftBytes, _ := hex.DecodeString(string(leftNode))
	rightBytes, _ := hex.DecodeString(string(rightNode))
	hasher.Write(leftBytes)
	hasher.Write(rightBytes)
	return MerkleNodeValue(hex.EncodeToString(hasher.Sum(nil)))
}

// computeZeroHashes computes the zero hashes for the SMT
func computeZeroHashes(height int) []MerkleNodeValue {
	currentZeroHash := MerkleNodeValue("0000000000000000000000000000000000000000000000000000000000000000")
	zeroHashes := []MerkleNodeValue{currentZeroHash}

	for i := 1; i <= height; i++ {
		currentZeroHash = hash(currentZeroHash, currentZeroHash)
		zeroHashes = append(zeroHashes, currentZeroHash)
	}

	return zeroHashes
}

// NodeStore stores the nodes of the SMT
type NodeStore struct {
	Nodes      map[string]MerkleNodeValue
	Height     int
	ZeroHashes []MerkleNodeValue
}

// NewNodeStore creates a new NodeStore
func NewNodeStore(height int) *NodeStore {
	return &NodeStore{
		Nodes:      make(map[string]MerkleNodeValue),
		Height:     height,
		ZeroHashes: computeZeroHashes(height),
	}
}

// Contains checks if a node exists in the store
func (ns *NodeStore) Contains(level, index int) bool {
	key := fmt.Sprintf("%d_%d", level, index)
	_, exists := ns.Nodes[key]
	return exists
}

// Set sets a node value in the store
func (ns *NodeStore) Set(level, index int, value MerkleNodeValue) {
	key := fmt.Sprintf("%d_%d", level, index)
	ns.Nodes[key] = value
}

// Get retrieves a node value from the store
func (ns *NodeStore) Get(level, index int) MerkleNodeValue {
	if ns.Contains(level, index) {
		key := fmt.Sprintf("%d_%d", level, index)
		return ns.Nodes[key]
	}
	return ns.ZeroHashes[ns.Height-level]
}

// ZeroMerkleTree represents the Sparse Merkle Tree
type ZeroMerkleTree struct {
	Height    int
	NodeStore *NodeStore
}

// NewZeroMerkleTree creates a new ZeroMerkleTree
func NewZeroMerkleTree(height int) *ZeroMerkleTree {
	return &ZeroMerkleTree{
		Height:    height,
		NodeStore: NewNodeStore(height),
	}
}

// SetLeaf sets a leaf value in the Merkle tree
func (zmt *ZeroMerkleTree) SetLeaf(index int, value MerkleNodeValue) DeltaMerkleProof {
	oldRoot := zmt.NodeStore.Get(0, 0)
	oldValue := zmt.NodeStore.Get(zmt.Height, index)

	var siblings []MerkleNodeValue
	currentIndex := index
	currentValue := value

	for level := zmt.Height; level >= 1; level-- {
		zmt.NodeStore.Set(level, currentIndex, currentValue)

		if currentIndex%2 == 0 {
			rightSibling := zmt.NodeStore.Get(level, currentIndex+1)
			currentValue = hash(currentValue, rightSibling)
			siblings = append(siblings, rightSibling)
		} else {
			leftSibling := zmt.NodeStore.Get(level, currentIndex-1)
			currentValue = hash(leftSibling, currentValue)
			siblings = append(siblings, leftSibling)
		}

		currentIndex /= 2
	}

	zmt.NodeStore.Set(0, 0, currentValue)

	return DeltaMerkleProof{
		Index:    index,
		Siblings: siblings,
		OldRoot:  oldRoot,
		OldValue: oldValue,
		NewRoot:  currentValue,
		NewValue: value,
	}
}

// GetLeaf retrieves a leaf value from the Merkle tree
func (zmt *ZeroMerkleTree) GetLeaf(index int) MerkleProof {
	var siblings []MerkleNodeValue

	value := zmt.NodeStore.Get(zmt.Height, index)
	currentIndex := index
	currentValue := value

	for level := zmt.Height; level >= 1; level-- {
		if currentIndex%2 == 0 {
			rightSibling := zmt.NodeStore.Get(level, currentIndex+1)
			currentValue = hash(currentValue, rightSibling)
			siblings = append(siblings, rightSibling)
		} else {
			leftSibling := zmt.NodeStore.Get(level, currentIndex-1)
			currentValue = hash(leftSibling, currentValue)
			siblings = append(siblings, leftSibling)
		}

		currentIndex /= 2
	}

	root := currentValue

	return MerkleProof{
		Root:     root,
		Siblings: siblings,
		Index:    index,
		Value:    value,
	}
}

// GetRoot returns the root of the Merkle tree
func (zmt *ZeroMerkleTree) GetRoot() MerkleNodeValue {
	return zmt.NodeStore.Get(0, 0)
}

// ComputeMerkleRootFromProof computes the Merkle root from a proof
func ComputeMerkleRootFromProof(siblings []MerkleNodeValue, index int, value MerkleNodeValue) MerkleNodeValue {
	merklePathNodeValue := value
	merklePathNodeIndex := index

	for _, sibling := range siblings {
		if merklePathNodeIndex%2 == 0 {
			merklePathNodeValue = hash(merklePathNodeValue, sibling)
		} else {
			merklePathNodeValue = hash(sibling, merklePathNodeValue)
		}
		merklePathNodeIndex /= 2
	}

	return merklePathNodeValue
}

// VerifyMerkleProof verifies a Merkle proof
func VerifyMerkleProof(proof MerkleProof) bool {
	return proof.Root == ComputeMerkleRootFromProof(proof.Siblings, proof.Index, proof.Value)
}

// VerifyDeltaMerkleProof verifies a delta Merkle proof
func VerifyDeltaMerkleProof(deltaProof DeltaMerkleProof) bool {
	oldProof := MerkleProof{
		Siblings: deltaProof.Siblings,
		Index:    deltaProof.Index,
		Root:     deltaProof.OldRoot,
		Value:    deltaProof.OldValue,
	}

	newProof := MerkleProof{
		Siblings: deltaProof.Siblings,
		Index:    deltaProof.Index,
		Root:     deltaProof.NewRoot,
		Value:    deltaProof.NewValue,
	}

	return VerifyMerkleProof(oldProof) && VerifyMerkleProof(newProof)
}
