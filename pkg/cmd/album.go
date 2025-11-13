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

var albumsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get Spotify catalog information for a single album.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the album.\n",
		},
		&jsonflag.JSONStringFlag{
			Name:  "market",
			Usage: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "market",
			},
		},
	},
	Action:          handleAlbumsRetrieve,
	HideHelpCommand: true,
}

var albumsBulkRetrieve = cli.Command{
	Name:  "bulk-retrieve",
	Usage: "Get Spotify catalog information for multiple albums identified by their Spotify\nIDs.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for the albums. Maximum: 20 IDs.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "ids",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "market",
			Usage: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "market",
			},
		},
	},
	Action:          handleAlbumsBulkRetrieve,
	HideHelpCommand: true,
}

var albumsListTracks = cli.Command{
	Name:  "list-tracks",
	Usage: "Get Spotify catalog information about an album’s tracks. Optional parameters can\nbe used to limit the number of tracks returned.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the album.\n",
		},
		&jsonflag.JSONIntFlag{
			Name:  "limit",
			Usage: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "limit",
			},
			Value: 20,
		},
		&jsonflag.JSONStringFlag{
			Name:  "market",
			Usage: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "market",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "offset",
			Usage: "The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "offset",
			},
			Value: 0,
		},
	},
	Action:          handleAlbumsListTracks,
	HideHelpCommand: true,
}

func handleAlbumsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.AlbumGetParams{}
	var res []byte
	_, err := cc.client.Albums.Get(
		ctx,
		cmd.Value("id").(string),
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
	return ShowJSON("albums retrieve", json, format, transform)
}

func handleAlbumsBulkRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.AlbumBulkGetParams{}
	var res []byte
	_, err := cc.client.Albums.BulkGet(
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
	return ShowJSON("albums bulk-retrieve", json, format, transform)
}

func handleAlbumsListTracks(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.AlbumListTracksParams{}
	var res []byte
	_, err := cc.client.Albums.ListTracks(
		ctx,
		cmd.Value("id").(string),
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
	return ShowJSON("albums list-tracks", json, format, transform)
}
