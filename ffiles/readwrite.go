package ffiles

import (
	"bufio"
	"os"
	"path/filepath"
)

func WriteNewFile(file string, content string) error {
	exists, err := Exists(file)
	if err != nil {
		return err
	}
	if exists {
		return os.ErrExist
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	_, err = w.WriteString(content)
	if err != nil {
		return err
	}

	return w.Flush()
}

func WriteLines(file string, lines []string) error {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, l := range lines {
		_, err := w.WriteString(l + "\n")
		if err != nil {
			return err
		}
	}

	return w.Flush()
}

func ReadLines(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return []string{}, err
	}
	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)

	var lines []string = []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, err
}

func ListFilesRecursively(folder string) ([]string, error) {
	filesList := []string{}
	err := filepath.Walk(folder, func(p string, f os.FileInfo, _ error) error {
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
	return filesList, err
}
