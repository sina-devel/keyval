package keyval

import (
	"errors"
	"sync"
)

var (
	// ErrNotExist returns when the key is not exist.
	ErrNotExist = errors.New("key is not exist")
)

// Store is an interface that Keyval uses to work on data.
type Store interface {
	// Put puts the value with the given key.
	Put(key string, val []byte) error

	// Get gets the key and returns the value.
	//
	// note: if key is not exist should return the ErrNotExist.
	Get(key string) ([]byte, error)

	// Drop drops the given key.
	Drop(key string) error

	// Keys returns all the stored keys.
	Keys() []string
}

// MemoryStore implements Store interface.
type MemoryStore struct {
	mu   sync.Mutex
	data map[string][]byte
}

// NewMemoryStore returns an initialized MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string][]byte),
	}
}

// Put puts the value with the given key.
//
// note: if the key already exists, it overwrites it.
func (m *MemoryStore) Put(key string, val []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = val
	return nil
}

// Get gets the key and returns the value.
func (m *MemoryStore) Get(key string) ([]byte, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if val, ok := m.data[key]; ok {
		return val, nil
	}

	return nil, ErrNotExist
}

// Drop drops the given key.
func (m *MemoryStore) Drop(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.data[key]; ok {
		delete(m.data, key)
		return nil
	}

	return ErrNotExist
}

// Keys returns all the stored keys.
func (m *MemoryStore) Keys() []string {
	m.mu.Lock()
	defer m.mu.Unlock()

	var keys []string
	for k := range m.data {
		keys = append(keys, k)
	}

	return keys
}
