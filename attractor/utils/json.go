package utils

import (
	"encoding/json"
	"io/ioutil"
)

func WriteConfigJSON(filename string, config interface{}) error {
	data, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func ReadConfigJSON(filename string, config interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, config)
}
