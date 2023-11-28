package merkle

type MerkleTree struct {
	Root       *MerkleNode
	Nodes      []*MerkleNode
	MerkleRoot []byte
}

type MerkleNode struct {
	Tree  *MerkleTree
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
	Hash  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	return nil
}

func NewMerkleTree() *MerkleTree {
	return nil
}

func NewMerkleTreeWithData(data [][]byte) *MerkleTree {

	return nil
}
