package config

import (
	"io/ioutil"
)

// ReadFile reads a file by the provided path and returns its content as bytes.
func ReadFile(filePath string) ([]byte, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}
