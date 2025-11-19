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

var artistsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get Spotify catalog information for a single artist identified by their unique\nSpotify ID.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the artist.\n",
		},
	},
	Action:          handleArtistsRetrieve,
	HideHelpCommand: true,
}

var artistsBulkRetrieve = cli.Command{
	Name:  "bulk-retrieve",
	Usage: "Get Spotify catalog information for several artists based on their Spotify IDs.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for the artists. Maximum: 50 IDs.\n",
		},
	},
	Action:          handleArtistsBulkRetrieve,
	HideHelpCommand: true,
}

var artistsListAlbums = cli.Command{
	Name:  "list-albums",
	Usage: "Get Spotify catalog information about an artist's albums.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the artist.\n",
		},
		&cli.StringFlag{
			Name:  "include-groups",
			Usage: "A comma-separated list of keywords that will be used to filter the response. If not supplied, all album types will be returned. <br/>\nValid values are:<br/>- `album`<br/>- `single`<br/>- `appears_on`<br/>- `compilation`<br/>For example: `include_groups=album,single`.\n",
		},
		&cli.Int64Flag{
			Name:  "limit",
			Usage: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Value: 20,
		},
		&cli.StringFlag{
			Name:  "market",
			Usage: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
		},
		&cli.Int64Flag{
			Name:  "offset",
			Usage: "The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.\n",
		},
	},
	Action:          handleArtistsListAlbums,
	HideHelpCommand: true,
}

var artistsListRelatedArtists = cli.Command{
	Name:  "list-related-artists",
	Usage: "Get Spotify catalog information about artists similar to a given artist.\nSimilarity is based on analysis of the Spotify community's listening history.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the artist.\n",
		},
	},
	Action:          handleArtistsListRelatedArtists,
	HideHelpCommand: true,
}

var artistsTopTracks = cli.Command{
	Name:  "top-tracks",
	Usage: "Get Spotify catalog information about an artist's top tracks by country.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the artist.\n",
		},
		&cli.StringFlag{
			Name:  "market",
			Usage: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
		},
	},
	Action:          handleArtistsTopTracks,
	HideHelpCommand: true,
}

func handleArtistsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err := client.Artists.Get(
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
	return ShowJSON("artists retrieve", json, format, transform)
}

func handleArtistsBulkRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.ArtistBulkGetParams{
		IDs: cmd.Value("ids").(string),
	}
	var res []byte
	_, err := client.Artists.BulkGet(
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
	return ShowJSON("artists bulk-retrieve", json, format, transform)
}

func handleArtistsListAlbums(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.ArtistListAlbumsParams{
		IncludeGroups: spotted.String(cmd.Value("include-groups").(string)),
		Market:        spotted.String(cmd.Value("market").(string)),
	}
	if cmd.IsSet("limit") {
		params.Limit = spotted.Opt(cmd.Value("limit").(int64))
	}
	if cmd.IsSet("offset") {
		params.Offset = spotted.Opt(cmd.Value("offset").(int64))
	}
	var res []byte
	_, err := client.Artists.ListAlbums(
		ctx,
		cmd.Value("id").(string),
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
	return ShowJSON("artists list-albums", json, format, transform)
}

func handleArtistsListRelatedArtists(ctx context.Context, cmd *cli.Command) error {
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
	_, err := client.Artists.ListRelatedArtists(
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
	return ShowJSON("artists list-related-artists", json, format, transform)
}

func handleArtistsTopTracks(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.ArtistTopTracksParams{
		Market: spotted.String(cmd.Value("market").(string)),
	}
	var res []byte
	_, err := client.Artists.TopTracks(
		ctx,
		cmd.Value("id").(string),
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
	return ShowJSON("artists top-tracks", json, format, transform)
}
