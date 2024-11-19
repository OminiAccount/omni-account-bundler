package merkletree

// leafType specifies type of the leaf
type leafType uint8

const (
	// LeafTypeBalance specifies that leaf stores Balance
	LeafTypeBalance leafType = 0
	// LeafTypeNonce specifies that leaf stores Nonce
	LeafTypeNonce leafType = 1
)
