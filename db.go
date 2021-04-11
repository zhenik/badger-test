package main

import (
	"github.com/dgraph-io/badger/v3"
	"log"
)

// Badger is a simple helper for db access
type Badger struct {
	db *badger.DB
}

// NewBadgerDBFromPath returns a new badgerdb
func NewBadgerDBFromPath(path string) (Badger, error) {
	b := Badger{}
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
		return b, err
	}
	b.db = db
	return b, nil
}
