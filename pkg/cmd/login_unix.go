//go:build !windows

package cmd

import (
	"os/exec"
	"runtime"
)

func openBrowserOS(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	return cmd.Start()
}
