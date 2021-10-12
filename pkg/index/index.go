package index

import (
	"bytes"
	"errors"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/spf13/cast"
	"go.uber.org/atomic"
)

var atom atomic.Uint32

// Index idx
type Index struct {
	db *badger.DB
}

type ObjectType int

// Object Types
const (
	ObjectTypeFile ObjectType = iota + 1
	ObjectTypeDir
	ObjectTypeMeta
)

const (
	existPrefix  = "_cid"
	recentPrefix = "_recent"
)

// NewIndex new idx
func NewIndex(path string) (*Index, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}

	idx := &Index{
		db: db,
	}
	return idx, nil
}

// SetExist cache a cid to filesystem
func (i *Index) SetExist(cid string, ot ObjectType) error {
	return i.db.Update(func(txn *badger.Txn) error {
        // Cache recent created folders cid, we will list them in the gallery page.
		if ot == ObjectTypeDir {
			sortStr := cast.ToString(time.Now().UnixNano())
			sortStr += cast.ToString(atom.Add(1))
			txn.Set([]byte(recentPrefix+sortStr), []byte(cid))
		}
		return txn.Set([]byte(existPrefix+cid), []byte(cast.ToString(ot)))
	})
}

// Exist check if a cid exsit
func (i *Index) Exist(cid string) (ok bool, err error) {
	i.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(existPrefix + cid))
		if err == nil {
			ok = true
			return nil
		}
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil
		}
		return err
	})
	return
}

// FIXME not works...
func (i *Index) FilterFileCid(limit int) (ids []string, err error) {
	err = i.db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		// opts.Prefix = []byte(recentPrefix)
		opts.Reverse = true
		it := txn.NewIterator(opts)
		defer it.Close()

		var i int
		for it.Rewind(); it.Valid(); it.Next() {
            if !bytes.HasPrefix(it.Item().Key(), []byte(recentPrefix)) {
                continue
            }
			i++
			// delete key
			if i > limit {
				if err := txn.Delete(it.Item().Key()); err != nil {
					return err
				}
				continue
			}
			v, _ := it.Item().ValueCopy(nil)
			ids = append(ids, string(v))
		}
		return nil
	})
	return
}

func (i *Index) Close() {
	i.db.Close()
}
