package ffiles

import (
	"io/ioutil"

	"github.com/otiai10/copy"
)

// Copies a file from `source` to `destin`.
func CopyFile(source string, destin string) error {
	bytesRead, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(destin, bytesRead, 0600)
}

// Copies a folder from `source` to `destin`.
func CopyFolder(source string, destination string) error {
	return copy.Copy(source, destination)
}
