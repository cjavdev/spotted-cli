package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var whoamiCommand = cli.Command{
	Name:      "whoami",
	Category:  "AUTH",
	Usage:     "Show the currently authenticated Spotify user",
	UsageText: "spotted whoami",
	Action:    whoamiAction,
}

func whoamiAction(ctx context.Context, cmd *cli.Command) error {
	token := GetAccessToken()
	if token == "" {
		fmt.Println("Not logged in.")
		fmt.Println("Run 'spotted login' to authenticate.")
		return nil
	}

	opts := getDefaultRequestOptions(cmd)
	opts = append(opts, option.WithAccessToken(token))

	if cmd.Bool("debug") {
		opts = append(opts, debugMiddlewareOption)
	}

	client := spotted.NewClient(opts...)
	me, err := client.Me.Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to get user info: %w", err)
	}

	jsonData, err := json.Marshal(me)
	if err != nil {
		return err
	}

	result := gjson.ParseBytes(jsonData)
	return ShowJSON(os.Stdout, "User", result, cmd.String("format"), cmd.String("transform"))
}
