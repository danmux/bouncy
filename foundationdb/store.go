package foundationdb

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/blevesearch/bleve/index/store"
	"github.com/blevesearch/bleve/registry"
)

const (
	// Name is the name of this KVStore
	Name = "foundationdb"
)

// Store is a FoundationDB implementation of the KVStore interface
type Store struct {
	mo store.MergeOperator
	db fdb.Database
}

// New returns a new Store
func New(mo store.MergeOperator, config map[string]interface{}) (store.KVStore, error) {
	err := fdb.APIVersion(600)
	if err != nil {
		return nil, err
	}
	db, err := fdb.OpenDefault()
	if err != nil {
		return nil, err
	}
	return &Store{
		mo: mo,
		db: db,
	}, nil
}

// Writer returns a KVWriter which can be used to
// make changes to the FDB.  If a writer cannot
// be obtained a non-nil error is returned.
func (s *Store) Writer() (store.KVWriter, error) {
	return &Writer{
		store: s,
	}, nil
}

// Reader returns a KVReader which can be used to
// read data from the KVStore.  If a reader cannot
// be obtained a non-nil error is returned.
func (s *Store) Reader() (store.KVReader, error) {
	return &Reader{
		db: s.db,
	}, nil
}

// Close closes the KVStore
func (s *Store) Close() error {
	return nil
}

// make this store available to bleve
func init() {
	registry.RegisterKVStore(Name, New)
}
