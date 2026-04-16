// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cjavdev/spotted-cli/internal/apiquery"
	"github.com/cjavdev/spotted-cli/internal/binaryparam"
	"github.com/cjavdev/spotted-cli/internal/requestflag"
	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var playlistsImagesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Replace the image used to represent a specific playlist.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "playlist-id",
			Usage:    "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "body",
			Usage:     "Base64 encoded JPEG image data, maximum payload size is 256 KB.",
			Required:  true,
			BodyRoot:  true,
			FileInput: true,
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handlePlaylistsImagesUpdate,
	HideHelpCommand: true,
}

var playlistsImagesList = cli.Command{
	Name:    "list",
	Usage:   "Get the current image associated with a specific playlist.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "playlist-id",
			Usage:    "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
			Required: true,
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

	bodyReader, stdinInUse, err := binaryparam.FileOrStdin(os.Stdin, cmd.Value("body").(string))
	if err != nil {
		return fmt.Errorf("Failed on param '%s': %w", "body", err)
	}
	defer bodyReader.Close()

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationOctetStream,
		stdinInUse,
	)
	if err != nil {
		return err
	}

	response, err := client.Playlists.Images.Update(
		ctx,
		cmd.Value("playlist-id").(string),
		bodyReader,
		options...,
	)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, os.Stdout, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Playlists.Images.List(ctx, cmd.Value("playlist-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		Title:          "playlists:images list",
		Transform:      transform,
	})
}
