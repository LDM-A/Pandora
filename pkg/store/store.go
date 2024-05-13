package store

import (
	"errors"
	"sync"
	//"github.com/LDM-A/pandoratree/pkg/merkle"
)

type Storage interface {
	Get(string) ([]byte, error)
	Has(string) bool
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
func (kv *KVStorage) Has(key string) bool {
	_, ok := kv.data[key]
	return ok
}
func (kv *KVStorage) Get(key string) ([]byte, error) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	value, ok := kv.data[key]
	if !ok {
		return nil, errors.New("value not found")
	}
	return value, nil
}

func (kv *KVStorage) Put(key string, value []byte) error {
	if ok := kv.Has(key); !ok {
		kv.mu.Lock()
		//if err := kv.merkleTree.Add(key, value); err != nil {
		//	return errors.New("Failed to hash data into Merkle tree. Do not store the data")
		//}

		defer kv.mu.Unlock()
		kv.data[key] = value
		return nil
	}
	return errors.New("value already exists")
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
	mu   sync.RWMutex
	data map[string][]byte
	//merkleTree *merkle.MerkleTree
}

func NewKVStorage() *KVStorage {
	return &KVStorage{
		data: make(map[string][]byte),
		//merkleTree: merkle.NewMerkleTree(),
		mu: sync.RWMutex{},
	}
}
