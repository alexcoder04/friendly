package ffiles

import (
	"fmt"
	"os"
	"path"
)

// Returns a sensible config directory for `program`.
// Uses `os.UserConfigDir()` under the hood and creates the directory if it does not exist yet.
func GetConfigDirFor(program string) (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	configDir := path.Join(dir, program)
	err = os.MkdirAll(configDir, 0700)
	if err != nil {
		return "", err
	}

	return configDir, nil
}

// Returns a sensible config directory for `program`.
// `os.TempDir()` is used, so the logs may not be persistent depensing on the operating system.
// Creates the directory if it does not exist yet.
func GetLogDirFor(program string) (string, error) {
	logDir := path.Join(os.TempDir(), fmt.Sprintf("%s-%d", program, os.Getuid()))
	err := os.MkdirAll(logDir, 0700)
	if err != nil {
		return "", err
	}
	return logDir, nil
}

// Returns a sensible cache directory for `program`.
// Uses `os.UserCacheDir()` under the hood and creates the directory if it does not exist yet.
func GetCacheDirFor(program string) (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	cacheDir := path.Join(dir, program)
	err = os.MkdirAll(cacheDir, 0700)
	if err != nil {
		return "", err
	}

	return cacheDir, nil
}
