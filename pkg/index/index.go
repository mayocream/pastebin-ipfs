package index

import (
	"errors"

	"github.com/dgraph-io/badger/v3"
)

// Index idx
type Index struct {
	db *badger.DB
}

const (
	existPrefix = "_cid"
)

// NewIndex new idx
func NewIndex(path string) (*Index, error){
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}

	idx := &Index{
		db: db,
	}
	return idx, nil
}

func (i *Index) SetExist(cid string) error {
	return i.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(existPrefix + cid), nil)
	})
}

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

func (i *Index) Close() {
	i.db.Close()
}
