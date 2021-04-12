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

// GetValue retrieves a value []byte from data store and return a copy of it
func (s *Store) GetValue(key string) ([]byte, error) {
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

// GetItem retrieves an Item from data store. Throw
func (s *Store) GetItem(key string) (Item, error) {
	valueBytes, _ := s.GetValue(key)

	var item Item
	d := gob.NewDecoder(bytes.NewReader(valueBytes))
	if err := d.Decode(&item); err != nil {
		log.Println("Decoding error")
		return Item{}, err
	}
	log.Println("Item decoded", item)
	return item, nil
}

// GetValues retrieves a value from badgerhold and puts it into result.  Result must be a pointer
func (s *Store) GetValues() ([]KVP, error) {
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

func (s *Store) GetItems() ([]Item, error) {
	valuesBytes, err := s.GetValues()
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
