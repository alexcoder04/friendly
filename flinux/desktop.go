//go:build linux
// +build linux

package flinux

import (
	"os"
	"path"
	"strings"

	"github.com/alexcoder04/friendly/ffiles"
)

func GetDisplayServer() string {
	sessionType := os.Getenv("XDG_SESSION_TYPE")
	if sessionType != "" {
		return sessionType
	}
	if os.Getenv("WAYLAND_DISPLAY") != "" {
		return "wayland"
	}
	if os.Getenv("DISPLAY") != "" {
		return "x11"
	}
	return ""
}

func GuiRunning() bool {
	dispServer := GetDisplayServer()
	if dispServer == "" {
		return false
	}

	if dispServer == "wayland" {
		display := os.Getenv("WAYLAND_DISPLAY")
		if display == "" {
			return false
		}
		runTimeDir, err := ffiles.GetRuntimeDir()
		if err != nil {
			return false
		}
		waySock := path.Join(runTimeDir, display)
		return ffiles.IsFile(waySock)
	}

	display := strings.Replace(os.Getenv("DISPLAY"), ":", "", 1)
	if display == "" {
		return false
	}
	x11Sock := path.Join(os.TempDir(), ".X11-unix", "X"+display)
	return ffiles.IsFile(x11Sock)
}
