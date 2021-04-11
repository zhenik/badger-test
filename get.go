package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/dgraph-io/badger/v3"
	"log"
)
// ErrNotFound is returned when no data is found for the given key
var ErrNotFound = errors.New("no data found for this key")

// GetValueBytes retrieves a value from badgerhold and puts it into result.  Result must be a pointer
func (s *Store) GetValueBytes(key string) ([]byte, error) {
	encodedKey := encodeKey(key)

	var valueCopy []byte
	err := s.Badger().View(func(txn *badger.Txn) error {
		item, err := txn.Get(encodedKey)
		if err != nil {
			return ErrNotFound
		}
		err = item.Value(func(val []byte) error {
			valueCopy = append([]byte{}, val...)
			return nil
		})
		return err
	})
	return valueCopy, err
}

func (s *Store) Get(key string) (Item, error) {
	valueBytes, _ := s.GetValueBytes(key)


	var item Item
	d := gob.NewDecoder(bytes.NewReader(valueBytes))
	if err := d.Decode(&item); err != nil {
		log.Println("Decoding error")
	}
	log.Println("Item decoded",item)
	return item, nil
}

