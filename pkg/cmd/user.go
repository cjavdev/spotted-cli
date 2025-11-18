// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var usersRetrieveProfile = cli.Command{
	Name:  "retrieve-profile",
	Usage: "Get public profile information about a Spotify user.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "user-id",
			Usage: "The user's [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids).\n",
		},
	},
	Action:          handleUsersRetrieveProfile,
	HideHelpCommand: true,
}

func handleUsersRetrieveProfile(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := client.Users.GetProfile(
		ctx,
		cmd.Value("user-id").(string),
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("users retrieve-profile", json, format, transform)
}
