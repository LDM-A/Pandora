package merkle

import (
	"crypto/sha256"
	"errors"
)

type MerkleTree struct {
	Root *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

func NewMerkleNode(left *MerkleNode, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
	}

	node.Left = left
	node.Right = right

	return &node
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}

		nodes = newLevel
	}
	t := MerkleTree{&nodes[0]}
	return &t
}

func (mt *MerkleTree) addToTree(root, newNode *MerkleNode) (*MerkleNode, error) {
	if root == nil {
		return nil, errors.New("")
	}

	if root.Left == nil {
		return NewMerkleNode(newNode, root, nil), nil
	}

	if root.Right == nil {
		return NewMerkleNode(root, newNode, nil), nil
	}

	//leftHash := root.Left.Data
	//rightHash := root.Right.Data
	//prevHashes := append(leftHash, rightHash...)

	//hash := sha256.Sum256(prevHashes)
	//root.Data = hash[]

	//root.Left = mt.addToTree(root.Left, newNode)

	return NewMerkleNode(root.Left, root.Right, root.Data), nil
}
