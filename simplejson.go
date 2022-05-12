package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Decode(path string, data interface{}) error {
	d, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	return json.Unmarshal(d, &data)
}

func EncodeWithIndent(path string, data interface{}) error {
	d, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, d, os.ModePerm)
}

func Encode(path string, data interface{}) error {
	d, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, d, os.ModePerm)
}
