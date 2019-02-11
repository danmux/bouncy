package foundationdb

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/blevesearch/bleve/index/store"
)

// Reader is an **ISOLATED** reader
// In this context isolated is defined to mean that
// writes/deletes made after the KVReader is opened
// are not observed.
// Because there is usually a cost associated with
// keeping isolated readers active, users should
// close them as soon as they are no longer needed.
type Reader struct {
	db fdb.Database
}

// Get returns the value associated with the key
// If the key does not exist, nil is returned.
// The caller owns the bytes returned.
func (r *Reader) Get(key []byte) ([]byte, error) {
	val, err := r.db.Transact(func(tr fdb.Transaction) (interface{}, error) {
		return tr.Get(fdb.Key(key)).Get()
	})
	if err != nil {
		return nil, err
	}
	return val.([]byte), nil
}

// MultiGet retrieves multiple values in one call.
func (r *Reader) MultiGet(keys [][]byte) ([][]byte, error) {
	panic("reader multi get unimplemented")
}

// PrefixIterator returns a KVIterator that will
// visit all K/V pairs with the provided prefix
func (r *Reader) PrefixIterator(prefix []byte) store.KVIterator {
	tx, err := r.db.CreateTransaction()
	if err != nil {
		return &Iterator{
			err: err,
		}
	}

	prefixRange, err := fdb.PrefixRange(prefix)
	if err != nil {
		return &Iterator{
			err: err,
		}
	}

	it := &Iterator{
		tx:       tx,
		iterator: tx.GetRange(prefixRange, fdb.RangeOptions{}).Iterator(),
	}

	// the iterator is expected to be at the first record
	it.Next()

	return it
}

// RangeIterator returns a KVIterator that will
// visit all K/V pairs >= start AND < end
func (r *Reader) RangeIterator(start, end []byte) store.KVIterator {
	panic("range iterator unimplemented " + string(start) + " " + string(end))

}

// Close closes the Reader
func (r *Reader) Close() error {
	return nil
}
