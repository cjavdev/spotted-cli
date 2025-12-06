// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/cjavdev/spotted-cli/internal/apiquery"
	"github.com/cjavdev/spotted-cli/internal/requestflag"
	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var playlistsImagesUpdate = cli.Command{
	Name:  "update",
	Usage: "Replace the image used to represent a specific playlist.",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "playlist-id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
		},
		&requestflag.StringFlag{
			Name:  "body",
			Usage: "Base64 encoded JPEG image data, maximum payload size is 256 KB.",
			Config: requestflag.RequestConfig{
				BodyPath: "body",
			},
		},
	},
	Action:          handlePlaylistsImagesUpdate,
	HideHelpCommand: true,
}

var playlistsImagesList = cli.Command{
	Name:  "list",
	Usage: "Get the current image associated with a specific playlist.",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "playlist-id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
		},
	},
	Action:          handlePlaylistsImagesList,
	HideHelpCommand: true,
}

func handlePlaylistsImagesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("playlist-id") && len(unusedArgs) > 0 {
		cmd.Set("playlist-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("body") && len(unusedArgs) > 0 {
		cmd.Set("body", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Playlists.Images.Update(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "playlist-id"),
		requestflag.CommandRequestValue[io.Reader](cmd, "body"),
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "playlists:images update", obj, format, transform)
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
	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Playlists.Images.List(ctx, requestflag.CommandRequestValue[string](cmd, "playlist-id"), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "playlists:images list", obj, format, transform)
}
