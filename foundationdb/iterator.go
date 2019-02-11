package foundationdb

import "github.com/apple/foundationdb/bindings/go/src/fdb"

// Iterator holds the fields required to provide fdb iteration
// it implements KVIterator
type Iterator struct {
	err      error
	done     bool
	current  *fdb.KeyValue
	tx       fdb.Transaction
	iterator *fdb.RangeIterator
}

// Seek will advance the iterator to the specified key
func (i *Iterator) Seek(key []byte) {
	panic("seek sounds pretty expensive - not implemented")
}

// Next will advance the iterator to the next key
func (i *Iterator) Next() {
	if !i.iterator.Advance() {
		i.done = true
		i.current = nil
		return
	}
	kv, err := i.iterator.Get()
	if err != nil {
		i.err = err
		i.current = nil
		return
	}
	i.current = &kv
}

// Key returns the key pointed to by the iterator
// The bytes returned are **ONLY** valid until the next call to Seek/Next/Close
// Continued use after that requires that they be copied.
func (i *Iterator) Key() []byte {
	if i.current == nil {
		return nil
	}
	return i.current.Key
}

// Value returns the value pointed to by the iterator
// The bytes returned are **ONLY** valid until the next call to Seek/Next/Close
// Continued use after that requires that they be copied.
func (i *Iterator) Value() []byte {
	if i.current == nil {
		return nil
	}
	return i.current.Value
}

// Valid returns whether or not the iterator is in a valid state
func (i *Iterator) Valid() bool {
	if i.iterator == nil {
		return false
	}
	return i.err == nil && !i.done
}

// Current returns Key(),Value(),Valid() in a single operation
func (i *Iterator) Current() ([]byte, []byte, bool) {
	if !i.Valid() || i.current == nil {
		return nil, nil, false
	}
	return i.current.Key, i.current.Value, true
}

// Close closes the iterator
func (i *Iterator) Close() error {
	i.tx.Cancel()
	i.current = nil
	i.done = true
	return nil
}
