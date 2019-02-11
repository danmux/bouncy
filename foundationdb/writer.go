package foundationdb

import (
	"fmt"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/blevesearch/bleve/index/store"
)

// Writer is an abstraction for mutating the KVStore
// Writer does **NOT** enforce restrictions of a single writer
// if the underlying KVStore allows concurrent writes, the
// KVWriter interface should also do so, it is up to the caller
// to do this in a way that is safe and makes sense
type Writer struct {
	store *Store
}

// NewBatch returns a KVBatch for performing batch operations on this kvstore
func (w *Writer) NewBatch() store.KVBatch {
	return store.NewEmulatedBatch(w.store.mo)
}

// NewBatchEx returns a KVBatch and an associated byte array
// that's pre-sized based on the KVBatchOptions.  The caller can
// use the returned byte array for keys and values associated with
// the batch.  Once the batch is either executed or closed, the
// associated byte array should no longer be accessed by the
// caller.
func (w *Writer) NewBatchEx(options store.KVBatchOptions) ([]byte, store.KVBatch, error) {
	return make([]byte, options.TotalBytes), w.NewBatch(), nil
}

// ExecuteBatch will execute the KVBatch, the provided KVBatch **MUST** have
// been created by the same KVStore (though not necessarily the same KVWriter)
// Batch execution is atomic, either all the operations or none will be performed
func (w *Writer) ExecuteBatch(batch store.KVBatch) error {

	emulatedBatch, ok := batch.(*store.EmulatedBatch)
	if !ok {
		return fmt.Errorf("wrong type of batch")
	}

	// transaction is only mutations and we don't care about the result
	_, err := w.store.db.Transact(func(tr fdb.Transaction) (interface{}, error) {
		for k, mergeOps := range emulatedBatch.Merger.Merges {
			key := fdb.Key(k)
			// TODO concurrent?
			existingVal, e := tr.Get(key).Get()
			if e != nil {
				return nil, e
			}
			mergedVal, fullMergeOk := w.store.mo.FullMerge([]byte(k), existingVal, mergeOps)
			if !fullMergeOk {
				return nil, fmt.Errorf("merge operator returned failure")
			}
			tr.Set(key, mergedVal)
		}
		return nil, nil
	})
	if err != nil {
		return err
	}

	_, err = w.store.db.Transact(func(tr fdb.Transaction) (interface{}, error) {
		for _, op := range emulatedBatch.Ops {
			if op.V != nil {
				tr.Set(fdb.Key(op.K), op.V)
			} else {
				tr.Clear(fdb.Key(op.K))
			}
		}
		return nil, nil
	})

	return err
}

// Close closes the writer
func (w *Writer) Close() error {
	return nil
}
