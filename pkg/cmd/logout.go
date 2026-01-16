package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

var logoutCommand = cli.Command{
	Name:      "logout",
	Category:  "AUTH",
	Usage:     "Remove stored Spotify credentials",
	UsageText: "spotted logout",
	Action:    logoutAction,
}

func logoutAction(ctx context.Context, cmd *cli.Command) error {
	path, err := CredentialsPath()
	if err != nil {
		return fmt.Errorf("failed to get credentials path: %w", err)
	}

	if err := DeleteCredentials(); err != nil {
		return fmt.Errorf("failed to delete credentials: %w", err)
	}

	fmt.Printf("Logged out. Credentials removed from: %s\n", path)
	return nil
}
