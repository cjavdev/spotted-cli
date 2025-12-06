// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cjavdev/spotted-cli/internal/apiquery"
	"github.com/cjavdev/spotted-cli/internal/requestflag"
	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var mePlaylistsList = cli.Command{
	Name:  "list",
	Usage: "Get a list of the playlists owned or followed by the current Spotify user.",
	Flags: []cli.Flag{
		&requestflag.IntFlag{
			Name:  "limit",
			Usage: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Value: requestflag.Value[int64](20),
			Config: requestflag.RequestConfig{
				QueryPath: "limit",
			},
		},
		&requestflag.IntFlag{
			Name:  "offset",
			Usage: "'The index of the first playlist to return. Default:\n0 (the first object). Maximum offset: 100.000\\. Use with `limit` to get the\nnext set of playlists.'\n",
			Config: requestflag.RequestConfig{
				QueryPath: "offset",
			},
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

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Me.Playlists.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "me:playlists list", obj, format, transform)
	} else {
		iter := client.Me.Playlists.ListAutoPaging(ctx, params, options...)
		return streamOutput("me:playlists list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "me:playlists list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}
