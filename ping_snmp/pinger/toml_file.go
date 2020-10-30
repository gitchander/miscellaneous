package main

import (
	"bytes"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

func configSaveToml(v interface{}, filename string) error {
	buf := bytes.NewBuffer(nil)
	encoder := toml.NewEncoder(buf)
	err := encoder.Encode(v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, buf.Bytes(), 0664)
}

func configLoadToml(v interface{}, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return toml.Unmarshal(data, v)
}
