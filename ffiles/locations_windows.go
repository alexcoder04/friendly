//go:build windows
// +build windows

package ffiles

import (
	"os"
)

// Gets the runtime directory and creates it if it does not exist yet.
// `$XDG_RUNTIME_DIR` is used on Linux, `os.TempDir()` on Windows.
func GetRuntimeDir() (string, error) {
	dir := os.TempDir()
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return "", err
	}
	return dir, nil
}
