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

var mePlaylistsList = cli.Command{
	Name:  "list",
	Usage: "Get a list of the playlists owned or followed by the current Spotify user.",
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:  "limit",
			Usage: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Value: 20,
		},
		&cli.Int64Flag{
			Name:  "offset",
			Usage: "'The index of the first playlist to return. Default:\n0 (the first object). Maximum offset: 100.000\\. Use with `limit` to get the\nnext set of playlists.'\n",
		},
	},
	Action:          handleMePlaylistsList,
	HideHelpCommand: true,
}

func handleMePlaylistsList(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlaylistListParams{}
	if cmd.IsSet("limit") {
		params.Limit = spotted.Opt(cmd.Value("limit").(int64))
	}
	if cmd.IsSet("offset") {
		params.Offset = spotted.Opt(cmd.Value("offset").(int64))
	}
	var res []byte
	_, err := client.Me.Playlists.List(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
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
