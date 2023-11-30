package store

import (
	"errors"
	"sync"

	"github.com/LDM-A/pandoratree/pkg/merkle"
)

type Storage interface {
	Get(string) ([]byte, error)
	Put(string, []byte) error
	Delete(string) ([]byte, error)
	Update(string, []byte) error
}

type VerifiedStorage interface {
	VerifiedGet(string) ([]byte, error)
	VerifiedPut(string, []byte) error
	VerifiedDelete(string) ([]byte, error)
	VerifiedUpdate(string, []byte) error
}

// VerifiedGet returns the value associated with the proof, and the data needed for an auditor to confirm its validity
// Returns an error if proofing the Get failed server side. (If the server sends an OK this does not mean it has not been tampered with)
func (kv *KVStorage) VerifiedGet(key string) ([]byte, error) {
	return nil, nil
}

func (kv *KVStorage) VerifiedPut(key string) ([]byte, error) {
	return nil, nil
}

func (kv *KVStorage) VerifiedDelete(key string) ([]byte, error) {
	return nil, nil
}

func (kv *KVStorage) VerifiedUpdate(key string) ([]byte, error) {
	return nil, nil
}

/*
	When Auditor implementation comes will use Actors to send messages
	between auditors to agree on Delete/Update actions and rebalance merkle trees
*/

func (kv *KVStorage) Get(key string) ([]byte, error) {

	value := kv.data[key]
	if value != nil {
		return nil, errors.New("Value not found")
	}
	return value, nil
}

func (kv *KVStorage) Put(key string, value []byte) error {
	if _, found := kv.Get(key); found != nil {
		kv.data[key] = value
		return nil
	}
	return errors.New("Value already exists.")
}

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
		mu:         sync.RWMutex{},
	}
}
