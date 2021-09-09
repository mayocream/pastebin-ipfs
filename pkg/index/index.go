package index

import (
	"errors"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/spf13/cast"
)

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
	recentPrefix = "_re"
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

func (i *Index) SetExist(cid string, ot ObjectType) error {
	return i.db.Update(func(txn *badger.Txn) error {
		if ot == ObjectTypeFile {
			// count recent
			txn.Set([]byte(recentPrefix+cast.ToString(time.Now().Unix())), []byte(cid))
		}
		return txn.Set([]byte(existPrefix+cid), []byte(cast.ToString(ot)))
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

func (i *Index) FilterFileCid(limit int) (ids []string, err error) {
	err = i.db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = []byte(recentPrefix)
		opts.Reverse = true
		it := txn.NewIterator(opts)
		for i := 0; it.Valid(); it.Next() {
            i++
            ids = append(ids, string(it.Item().KeyCopy(nil)))
            // delete key
            if i > limit {
                if err := txn.Delete(it.Item().Key()); err != nil {
                    return err
                }
            }
		}
        return nil
	})
    return
}

func (i *Index) Close() {
	i.db.Close()
}
