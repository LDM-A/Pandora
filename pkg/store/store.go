package store

import (
	"sync"

	"github.com/LDM-A/pandoratree/pkg/merkle"
)

type Storage interface {
	Get(string) ([]byte, error)
	Put(string, []byte) error
	Delete(string) ([]byte, error)
	Update(string, []byte) error
}

func (kv *KVStorage) Get(key string) ([]byte, error) {
	return nil, nil
}

func (kv *KVStorage) Put(key string, value []byte) error {
	return nil
}

/*
	When Auditor implementation comes will use Actors to send messages
	between auditors to agree on Delete/Update actions and rebalance merkle trees
*/

func (kv *KVStorage) Delete(key string) ([]byte, error) {
	return nil, nil
}

func (kv *KVStorage) Update(key string, value []byte) error {
	return nil
}

/*
	For now before moving to a more complex structure like an AVL tree
	will be using a KV store backed by the merkle tree for verification.
*/

type KVStorage struct {
	mu         sync.RWMutex
	data       map[string][]byte
	merkleTree *merkle.MerkleTree
}

func NewKVStorage() *KVStorage {
	return &KVStorage{
		data:       make(map[string][]byte),
		merkleTree: merkle.NewMerkleTree(),
	}
}
