// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var playlistsImagesList = cli.Command{
	Name:  "list",
	Usage: "Get the current image associated with a specific playlist.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "playlist-id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
		},
	},
	Action:          handlePlaylistsImagesList,
	HideHelpCommand: true,
}

func handlePlaylistsImagesList(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("playlist-id") && len(unusedArgs) > 0 {
		cmd.Set("playlist-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := client.Playlists.Images.List(
		ctx,
		cmd.Value("playlist-id").(string),
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("playlists:images list", json, format, transform)
}
