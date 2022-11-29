//go:build windows
// +build windows

package ffiles

import (
	"os"
)

func GetRuntimeDir() (string, error) {
	dir := os.TempDir()
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return "", err
	}
	return dir, nil
}
