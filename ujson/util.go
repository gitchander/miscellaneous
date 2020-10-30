package ujson

import (
	"bytes"
	"encoding/json"
	"errors"
)

type keyVal struct {
	key []byte
	val []byte
}

func merge_keyVal(kv *keyVal) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	buf.Write(kv.key)
	buf.WriteByte(':')
	buf.Write(kv.val)
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// { key : val }
func split_keyVal(data []byte) (*keyVal, error) {

	data, err := trimCurlyBrackets(data)
	if err != nil {
		return nil, err
	}

	i := bytes.IndexByte(data, ':')
	if i == -1 {
		return nil, errors.New("invalid object data")
	}

	kv := &keyVal{
		key: data[:i],
		val: data[i+1:],
	}

	return kv, nil
}

// check and trim {}
func trimCurlyBrackets(data []byte) ([]byte, error) {
	data = bytes.TrimSpace(data)
	if len(data) < 2 {
		return nil, errors.New("it is not object")
	}
	if data[0] != '{' {
		return nil, errors.New("has not open curly bracket")
	}
	n := len(data)
	if data[n-1] != '}' {
		return nil, errors.New("has not close curly bracket")
	}
	return data[1 : n-1], nil
}

func Marshal(key string, val interface{}) ([]byte, error) {

	var kv keyVal
	var err error

	kv.key, err = json.Marshal(key)
	if err != nil {
		return nil, err
	}

	kv.val, err = json.Marshal(val)
	if err != nil {
		return nil, err
	}

	return merge_keyVal(&kv)
}

type MakeValueFunc func(key string) (interface{}, error)

func Unmarshal(data []byte, makeValue MakeValueFunc) (interface{}, error) {

	kv, err := split_keyVal(data)
	if err != nil {
		return nil, err
	}

	var key string
	err = json.Unmarshal(kv.key, &key)
	if err != nil {
		return nil, err
	}

	val, err := makeValue(key)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(kv.val, val)
	if err != nil {
		return nil, err
	}

	return val, nil
}
