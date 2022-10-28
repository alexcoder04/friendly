//go:build windows
// +build windows

package ffiles

import (
	"os"
	"path/filepath"
	"strconv"
)

func GetRuntimeDir() (string, error) {
	dir := os.TempDir()
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return "", err
	}
	return dir, nil
}
