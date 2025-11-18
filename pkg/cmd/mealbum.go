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

var meAlbumsList = cli.Command{
	Name:  "list",
	Usage: "Get a list of the albums saved in the current Spotify user's 'Your Music'\nlibrary.",
	Flags: []cli.Flag{
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
	Action:          handleMeAlbumsList,
	HideHelpCommand: true,
}

var meAlbumsCheck = cli.Command{
	Name:  "check",
	Usage: "Check if one or more albums is already saved in the current Spotify user's 'Your\nMusic' library.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for the albums. Maximum: 20 IDs.\n",
		},
	},
	Action:          handleMeAlbumsCheck,
	HideHelpCommand: true,
}

var meAlbumsRemove = cli.Command{
	Name:  "remove",
	Usage: "Remove one or more albums from the current user's 'Your Music' library.",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:  "id",
			Usage: "A JSON array of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `[\"4iV5W9uYEdYUVa79Axb7Rh\", \"1301WleyT98MSxVHPZCA6M\"]`<br/>A maximum of 50 items can be specified in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._\n",
		},
	},
	Action:          handleMeAlbumsRemove,
	HideHelpCommand: true,
}

var meAlbumsSave = cli.Command{
	Name:  "save",
	Usage: "Save one or more albums to the current user's 'Your Music' library.",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:  "id",
			Usage: "A JSON array of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `[\"4iV5W9uYEdYUVa79Axb7Rh\", \"1301WleyT98MSxVHPZCA6M\"]`<br/>A maximum of 50 items can be specified in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._\n",
		},
	},
	Action:          handleMeAlbumsSave,
	HideHelpCommand: true,
}

func handleMeAlbumsList(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeAlbumListParams{
		Market: spotted.String(cmd.Value("market").(string)),
	}
	if cmd.IsSet("limit") {
		params.Limit = spotted.Opt(cmd.Value("limit").(int64))
	}
	if cmd.IsSet("offset") {
		params.Offset = spotted.Opt(cmd.Value("offset").(int64))
	}
	var res []byte
	_, err := client.Me.Albums.List(
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
	return ShowJSON("me:albums list", json, format, transform)
}

func handleMeAlbumsCheck(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeAlbumCheckParams{
		IDs: cmd.Value("ids").(string),
	}
	var res []byte
	_, err := client.Me.Albums.Check(
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
	return ShowJSON("me:albums check", json, format, transform)
}

func handleMeAlbumsRemove(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeAlbumRemoveParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"ids": "ids",
	}, &params); err != nil {
		return err
	}
	return client.Me.Albums.Remove(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}

func handleMeAlbumsSave(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeAlbumSaveParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"ids": "ids",
	}, &params); err != nil {
		return err
	}
	return client.Me.Albums.Save(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}
