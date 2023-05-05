package badgerdb

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"sync"

	bdb "github.com/dgraph-io/badger/v3"

	"notifier/pkg/entity"
)

const defaultSequenceBandwidth = 1
const sequencePrefix = "seq"

var db *bdb.DB
var mutex sync.Mutex

type Repository[E entity.Storable] struct {
	db *bdb.DB
}

func NewRepository[E entity.Storable]() Repository[E] {
	getDB()

	return Repository[E]{
		db: db,
	}
}

func getDB() {
	if db == nil {
		mutex.Lock()
		defer mutex.Unlock()
		var err error
		db, err = bdb.Open(bdb.DefaultOptions("/tmp/badger"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (r Repository[E]) Save(_ context.Context, data E) (E, error) {
	if data.Key() == "" {
		return data, errors.New("entity does not have a storage key")
	}

	err := r.db.Update(
		func(txn *bdb.Txn) error {
			rawData, err := json.Marshal(data)
			if err != nil {
				return err
			}

			return txn.Set(
				[]byte(r.buildID(
					data.Name(),
					data.Key(),
				)),
				rawData,
			)
		},
	)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r Repository[E]) Get(_ context.Context, key string) (E, error) {
	var t E
	err := r.db.View(
		func(txn *bdb.Txn) error {
			item, err := txn.Get([]byte(r.buildID(t.Name(), key)))
			if err != nil {
				return err
			}

			return item.Value(
				func(val []byte) error {
					return json.Unmarshal(val, &t)
				},
			)
		},
	)
	if err != nil {
		return t, err
	}

	return t, nil
}

func (r Repository[E]) GetFiltered(_ context.Context, filters ...entity.Filter[E]) ([]E, error) {
	var res = make([]E, 0)
	err := r.db.View(
		func(txn *bdb.Txn) error {
			var t E

			options := bdb.DefaultIteratorOptions
			options.Prefix = []byte(t.Name())

			it := txn.NewIterator(options)
			defer it.Close()

		outer:
			for it.Rewind(); it.Valid(); it.Next() {
				err := it.Item().Value(
					func(val []byte) error {
						return json.Unmarshal(val, &t)
					},
				)
				if err != nil {
					//return err
					continue
				}

				for _, filter := range filters {
					if ok := filter(t); !ok {
						continue outer
					}
				}

				res = append(res, t)
			}

			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r Repository[E]) CountFiltered(_ context.Context, filters ...entity.Filter[E]) (uint64, error) {
	var count uint64
	err := r.db.View(
		func(txn *bdb.Txn) error {
			var t E

			options := bdb.DefaultIteratorOptions
			options.Prefix = []byte(t.Name())

			it := txn.NewIterator(options)
			defer it.Close()

		outer:
			for it.Rewind(); it.Valid(); it.Next() {
				err := it.Item().Value(
					func(val []byte) error {
						return json.Unmarshal(val, &t)
					},
				)
				if err != nil {
					return err
				}

				for _, filter := range filters {
					if ok := filter(t); !ok {
						continue outer
					}
				}

				count++
			}

			return nil
		},
	)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r Repository[E]) Delete(_ context.Context, key string) error {
	return r.db.Update(
		func(txn *bdb.Txn) error {
			var t E

			return txn.Delete([]byte(r.buildID(t.Name(), key)))
		},
	)
}

func (r Repository[E]) Sequence(_ context.Context, key string) (uint64, error) {
	var t E
	seq, err := r.db.GetSequence(
		[]byte(r.buildID(
			sequencePrefix,
			t.Name(),
			key,
		)),
		defaultSequenceBandwidth,
	)
	if err != nil {
		return 0, err
	}

	next, err := seq.Next()
	if err != nil {
		return 0, err
	}

	return next, nil
}

func (Repository[E]) buildID(parts ...string) string {
	return strings.Join(parts, "-")
}
