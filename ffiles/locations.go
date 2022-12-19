package ffiles

import (
	"fmt"
	"os"
	"path"
)

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

func GetLogDirFor(program string) (string, error) {
	logDir := path.Join(os.TempDir(), fmt.Sprintf("%s-%d", program, os.Getuid()))
	err := os.MkdirAll(logDir, 0700)
	if err != nil {
		return "", err
	}
	return logDir, nil
}

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
