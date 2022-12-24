package friendly

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

// Get current working directory without returning an error.
// If os.Getwd() throws an error, it tries os.UserHomeDir().
// If it also throws an error, returns "".
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

// Get environmental variable.
// If it is not defined (=""), return the provided default value.
func Getenv(key string, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}

func prepareCommand(commandLine []string, workingDir string) *exec.Cmd {
	if workingDir == "" {
		workingDir = Getpwd()
	}

	cmd := exec.Command(commandLine[0], commandLine[1:]...)
	cmd.Dir = workingDir

	return cmd
}

// Runs a command and prints its output to stdin.
func Run(commandLine []string, workingDir string) error {
	cmd := prepareCommand(commandLine, workingDir)

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	return cmd.Run()
}

// Runs a command and returns its output as a string.
func GetOutput(commandLine []string, workingDir string) (string, error) {
	cmd := prepareCommand(commandLine, workingDir)
	out, err := cmd.Output()
	return string(out), err
}

// Runs a function in new goroutine if gets an os signal.
func RunOnSignal(s syscall.Signal, callback func()) {
	for {
		channel := make(chan os.Signal, 1)
		signal.Notify(channel, s)
		<-channel
		go callback()
	}
}
