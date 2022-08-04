package linux

import (
	"os"
	"path"
	"strings"

	"github.com/adrg/xdg"
	"github.com/alexcoder04/friendly"
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
		waySock := path.Join(xdg.RuntimeDir, display)
		return friendly.IsFile(waySock)
	}

	display := strings.Replace(os.Getenv("DISPLAY"), ":", "", 1)
	if display == "" {
		return false
	}
	x11Sock := path.Join(os.TempDir(), ".X11-unix", "X"+display)
	return friendly.IsFile(x11Sock)
}
