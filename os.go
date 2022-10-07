package friendly

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func Getpwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		pwd, err = os.UserHomeDir()
		if err != nil {
			return ""
		}
	}
	return pwd
}

func prepareCommand(command string, arguments []string, workingDir string) *exec.Cmd {
	if workingDir == "" {
		workingDir = Getpwd()
	}

	cmd := exec.Command(command, arguments...)
	cmd.Dir = workingDir

	return cmd
}

func Run(command string, arguments []string, workingDir string) error {
	cmd := prepareCommand(command, arguments, workingDir)

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	return cmd.Run()
}

func GetOutput(command string, arguments []string, workingDir string) (string, error) {
	cmd := prepareCommand(command, arguments, workingDir)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

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
