package friendly

import (
	"bufio"
	"io/ioutil"
	"os"

	"github.com/otiai10/copy"
)

func WriteLines(file string, lines []string) error {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	for _, l := range lines {
		_, err := w.WriteString(l + "\n")
		if err != nil {
			return err
		}
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return nil
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

func CopyFile(source string, destin string) error {
	bytesRead, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destin, bytesRead, 0600)
	if err != nil {
		return err
	}

	return nil
}

func CopyFolder(source string, destination string) error {
	return copy.Copy(source, destination)
}
