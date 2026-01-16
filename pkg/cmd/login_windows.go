//go:build windows

package cmd

import (
	"os/exec"
)

func openBrowserOS(url string) error {
	cmd := exec.Command("cmd", "/c", "start", url)
	return cmd.Start()
}
