package main

import (
	"bytes"
	"encoding/gob"
)

// EncodeFunc is a function for encoding a value into bytes
type EncodeFunc func(value interface{}) ([]byte, error)

// DecodeFunc is a function for decoding a value from bytes
type DecodeFunc func(data []byte, value interface{}) error

var encode EncodeFunc
var decode DecodeFunc

// DefaultEncode is the default encoding func for (Gob)
func DefaultEncode(value interface{}) ([]byte, error) {
	var buff bytes.Buffer

	en := gob.NewEncoder(&buff)

	err := en.Encode(value)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

// DefaultDecode is the default decoding func for (Gob)
func DefaultDecode(data []byte, value interface{}) error {
	var buff bytes.Buffer
	de := gob.NewDecoder(&buff)

	_, err := buff.Write(data)
	if err != nil {
		return err
	}

	return de.Decode(value)
}
