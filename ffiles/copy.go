package ffiles

import (
	"io/ioutil"

	"github.com/otiai10/copy"
)

func CopyFile(source string, destin string) error {
	bytesRead, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(destin, bytesRead, 0600)
}

func CopyFolder(source string, destination string) error {
	return copy.Copy(source, destination)
}
