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

var audioFeaturesRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get audio feature information for a single track identified by its unique\nSpotify ID.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.\n",
		},
	},
	Action:          handleAudioFeaturesRetrieve,
	HideHelpCommand: true,
}

var audioFeaturesBulkRetrieve = cli.Command{
	Name:  "bulk-retrieve",
	Usage: "Get audio features for multiple tracks based on their Spotify IDs.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids)\nfor the tracks. Maximum: 100 IDs.\n",
		},
	},
	Action:          handleAudioFeaturesBulkRetrieve,
	HideHelpCommand: true,
}

func handleAudioFeaturesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := client.AudioFeatures.Get(
		ctx,
		cmd.Value("id").(string),
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("audio-features retrieve", json, format, transform)
}

func handleAudioFeaturesBulkRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.AudioFeatureBulkGetParams{
		IDs: cmd.Value("ids").(string),
	}
	var res []byte
	_, err := client.AudioFeatures.BulkGet(
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
	return ShowJSON("audio-features bulk-retrieve", json, format, transform)
}
