package ffiles

import (
	"bufio"
	"os"
	"path/filepath"
)

// Creates a new file and writes a string into it.
// Throws an error if the file already exists.
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

// Writes an array of strings into a file as lines.
// Overwrites any existing content and creates the file if it does not exist yet.
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

// Reads a file and returns its lines as an array of strings.
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

// Finds all files in a directory and its subdirectories and returns their names as an array of strings.
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
