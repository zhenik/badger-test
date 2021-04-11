package main

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
)

// Get retrieves a value from badgerhold and puts it into result.  Result must be a pointer
func (s *Store) Get(key, result interface{}) error {
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get()
		handle(err)

		var valNot, valCopy []byte
		err := item.Value(func(val []byte) error {
			// This func with val would only be called if item.Value encounters no error.

			// Accessing val here is valid.
			fmt.Printf("The answer is: %s\n", val)

			// Copying or parsing val is valid.
			valCopy = append([]byte{}, val...)

			// Assigning val slice to another variable is NOT OK.
			valNot = val // Do not do this.
			return nil
		})
		handle(err)

		// DO NOT access val here. It is the most common cause of bugs.
		fmt.Printf("NEVER do this. %s\n", valNot)

		// You must copy it to use it outside item.Value(...).
		fmt.Printf("The answer is: %s\n", valCopy)

		// Alternatively, you could also use item.ValueCopy().
		valCopy, err = item.ValueCopy(nil)
		handle(err)
		fmt.Printf("The answer is: %s\n", valCopy)

		return nil
	})
	//return s.Badger().View(func(tx *badger.Txn) error {
	//	//return s.TxGet(tx, key, result)
	//	return s.Get(key, result)
	//})
}
