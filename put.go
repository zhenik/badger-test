package main

import (
	"github.com/dgraph-io/badger/v3"
)

// Upsert inserts the record into the badger if it doesn't exist.  If it does already exist, then it updates
// the existing record
func (s *Store) Upsert(key string, data interface{}) error {
	encodedKey := encodeKey(key)
	encodedValue, err := encodeValue(data)
	if err != nil {
		return err
	}
	err = s.Badger().Update(func(txn *badger.Txn) error {
		txn.Set(encodedKey, encodedValue)
		return nil
	})
	return err
}
