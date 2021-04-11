package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

func encodeKey(key string) []byte {
	return []byte(key)
}

// DefaultEncode is the default encoding func for badgerhold (Gob)
func encodeValue(value interface{}) ([]byte, error) {
	var buff bytes.Buffer

	en := gob.NewEncoder(&buff)

	err := en.Encode(value)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func decodeValue(value []byte) (Item, error) {
	var item Item
	d := gob.NewDecoder(bytes.NewReader(value))
	err := d.Decode(&item)
	if err != nil {
		log.Println("Decoding error")
		return Item{}, err
	}
	log.Println("Item decoded", item)
	return item, nil
}
