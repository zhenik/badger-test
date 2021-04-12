package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"log"
)

// KVP simple named key value pair storage
type KVP struct {
	Key   []byte
	Value []byte
}

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
	log.Println("Item decoded", item)
	return item, nil
}

// GetValuesBytes retrieves a value from badgerhold and puts it into result.  Result must be a pointer
func (s *Store) GetValuesBytes() ([]KVP, error) {
	var results []KVP

	err := s.Badger().View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				//fmt.Printf("key=%s, value=%s\n", k, v)
				res := KVP{k, v}
				results = append(results, res)
				return nil
			})

			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return []KVP{}, err
	}
	return results, nil
}

func (s *Store) GetAll() ([]Item, error) {
	valuesBytes, err := s.GetValuesBytes()
	if err != nil {
		return nil, err
	}
	var results []Item

	for _, row := range valuesBytes {
		item, err := decodeValue(row.Value)
		if err != nil {
			return nil, err
		}
		results = append(results, item)
	}
	return results, nil
}

func (s *Store) GetAllKeys() ([]string, error) {
	var results []string

	err := s.Badger().View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			row := it.Item()
			k := row.Key()
			results = append(results, string(k))
			fmt.Printf("key=%s\n", k)
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return results, nil
}
