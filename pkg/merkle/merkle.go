package merkle

import (
	"crypto/sha256"
	"hash"
)

type MerkleTree struct {
	Root         *MerkleNode
	rootHash     []byte
	Nodes        []*MerkleNode
	MerkleRoot   []byte
	hashStrategy func() hash.Hash
}

type MerkleNode struct {
	Tree  *MerkleTree
	Paren *MerkleNode
	Left  *MerkleNode
	Right *MerkleNode
	leaf  bool
	dup   bool
	Data  []byte // Will see if we store data in merkle tree or in underlying datastore
	// such as the key-value store, later more complicated structures using binary trees or avl trees
	Hash []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	return nil
}

func NewMerkleTree() *MerkleTree {
	tree := &MerkleTree{
		hashStrategy: sha256.New,
	}

	return tree
}

func NewMerkleTreeWithData(data [][]byte) *MerkleTree {

	return nil
}
