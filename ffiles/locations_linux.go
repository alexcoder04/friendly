//go:build linux
// +build linux

package ffiles

import (
	"os"
	"path/filepath"
	"strconv"
)

// Gets the runtime directory and creates it if it does not exist yet.
// `$XDG_RUNTIME_DIR` is used on Linux, `os.TempDir()` on Windows.
func GetRuntimeDir() (string, error) {
	var dir string
	if os.Getenv("XDG_RUNTIME_DIR") != "" {
		dir = os.Getenv("XDG_RUNTIME_DIR")
	} else {
		dir = filepath.Join("/run/user", strconv.Itoa(os.Getuid()))
	}
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return "", err
	}
	return dir, nil
}
