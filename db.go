package main

import (
	"github.com/dgraph-io/badger/v3"
	"log"
)

// Store is a badgerhold wrapper around a badger DB
type Store struct {
	db *badger.DB
}

// NewBadgerDBFromPath returns a new badgerdb
func NewBadgerDBFromPath(path string) (Store, error) {
	b := Store{}
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
		return b, err
	}
	b.db = db
	return b, nil
}

// Badger returns the underlying Badger DB
func (s *Store) Badger() *badger.DB {
	return s.db
}
