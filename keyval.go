package keyval

import "errors"

var (
	// ErrEmptyKey returns, when the key is empty.
	ErrEmptyKey = errors.New("key is empty")
)

// Keyval is a structure contains Store
// Use the New constructor to construct it. 
type Keyval struct {
	store Store
}

// New returns an initialized Keyval with the given Store.
func New(store Store) *Keyval {
	return &Keyval{
		store: store,
	}
}

// Put puts the value with the given key.
func (kv *Keyval) Put(key string, val []byte) error {
	if key == "" {
		return ErrEmptyKey
	}
	return kv.store.Put(key, val)
}

// Put puts the string value with the given key.
func (kv *Keyval) PutString(key string, val string) error {
	return kv.store.Put(key, []byte(val))
}

// Get gets the key and returns the value.
func (kv *Keyval) Get(key string) ([]byte, error) {
	return kv.store.Get(key)
}

// Get gets the key and returns the string value.
func (kv *Keyval) GetString(key string) (string, error) {
	val, err := kv.store.Get(key)
	return string(val), err
}

// Has returns true if the given key exists.
func (kv *Keyval) Has(key string) bool {
	_, err := kv.store.Get(key)
	return !errors.Is(err, ErrNotExist) && err == nil
}

// Keys returns all the stored keys.
func (kv *Keyval) Keys() []string {
	return kv.store.Keys()
}

// Drop drops the given key.
func (kv *Keyval) Drop(key string) error {
	return kv.store.Drop(key)
}

// Drop drops all keys.
func (kv *Keyval) DropAll() {
	for _, k := range kv.store.Keys() {
		kv.store.Drop(k)
	}
}
