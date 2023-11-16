package merkle

type MerkleTree struct {
	Root *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	return nil
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	root := &MerkleTree{&nodes[0]}
	return root
}
