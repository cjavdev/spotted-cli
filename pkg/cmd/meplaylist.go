// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/spotted-cli/pkg/jsonflag"
	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var mePlaylistsList = cli.Command{
	Name:  "list",
	Usage: "Get a list of the playlists owned or followed by the current Spotify user.",
	Flags: []cli.Flag{
		&jsonflag.JSONIntFlag{
			Name:  "limit",
			Usage: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "limit",
			},
			Value: 20,
		},
		&jsonflag.JSONIntFlag{
			Name:  "offset",
			Usage: "'The index of the first playlist to return. Default:\n0 (the first object). Maximum offset: 100.000\\. Use with `limit` to get the\nnext set of playlists.'\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "offset",
			},
			Value: 0,
		},
	},
	Action:          handleMePlaylistsList,
	HideHelpCommand: true,
}

func handleMePlaylistsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlaylistListParams{}
	var res []byte
	_, err := cc.client.Me.Playlists.List(
		ctx,
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("me:playlists list", json, format, transform)
}
