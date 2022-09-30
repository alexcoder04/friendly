package friendly

import (
	"os"
	"path/filepath"
)

func IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func ListFilesRecursively(folder string) ([]string, error) {
	filesList := []string{}
	err := filepath.Walk(folder, func(p string, f os.FileInfo, err error) error {
		if p == folder {
			return nil
		}
		stat, err := os.Stat(p)
		if err != nil {
			return err
		}
		if !stat.IsDir() {
			filesList = append(filesList, p[len(folder)+1:])
		}
		return nil
	})
	if err != nil {
		return filesList, err
	}
	return filesList, nil
}
