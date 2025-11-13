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

var meTracksList = cli.Command{
	Name:  "list",
	Usage: "Get a list of the songs saved in the current Spotify user's 'Your Music'\nlibrary.",
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
	Action:          handleMeTracksList,
	HideHelpCommand: true,
}

var meTracksCheck = cli.Command{
	Name:  "check",
	Usage: "Check if one or more tracks is already saved in the current Spotify user's 'Your\nMusic' library.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `ids=4iV5W9uYEdYUVa79Axb7Rh,1301WleyT98MSxVHPZCA6M`. Maximum: 50 IDs.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "ids",
			},
		},
	},
	Action:          handleMeTracksCheck,
	HideHelpCommand: true,
}

var meTracksRemove = cli.Command{
	Name:  "remove",
	Usage: "Remove one or more tracks from the current user's 'Your Music' library.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "ids",
			Usage: "A JSON array of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `[\"4iV5W9uYEdYUVa79Axb7Rh\", \"1301WleyT98MSxVHPZCA6M\"]`<br/>A maximum of 50 items can be specified in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "ids.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+id",
			Usage: "A JSON array of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `[\"4iV5W9uYEdYUVa79Axb7Rh\", \"1301WleyT98MSxVHPZCA6M\"]`<br/>A maximum of 50 items can be specified in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "ids.-1",
			},
		},
	},
	Action:          handleMeTracksRemove,
	HideHelpCommand: true,
}

var meTracksSave = cli.Command{
	Name:  "save",
	Usage: "Save one or more tracks to the current user's 'Your Music' library.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "ids",
			Usage: "A JSON array of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `[\"4iV5W9uYEdYUVa79Axb7Rh\", \"1301WleyT98MSxVHPZCA6M\"]`<br/>A maximum of 50 items can be specified in one request. _**Note**: if the `timestamped_ids` is present in the body, any IDs listed in the query parameters (deprecated) or the `ids` field in the body will be ignored._\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "ids.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+id",
			Usage: "A JSON array of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `[\"4iV5W9uYEdYUVa79Axb7Rh\", \"1301WleyT98MSxVHPZCA6M\"]`<br/>A maximum of 50 items can be specified in one request. _**Note**: if the `timestamped_ids` is present in the body, any IDs listed in the query parameters (deprecated) or the `ids` field in the body will be ignored._\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "ids.-1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "timestamped-ids.id",
			Usage: "A JSON array of objects containing track IDs with their corresponding timestamps. Each object must include a track ID and an `added_at` timestamp. This allows you to specify when tracks were added to maintain a specific chronological order in the user's library.<br/>A maximum of 50 items can be specified in one request. _**Note**: if the `timestamped_ids` is present in the body, any IDs listed in the query parameters (deprecated) or the `ids` field in the body will be ignored._\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "timestamped_ids.#.id",
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name:  "timestamped-ids.added_at",
			Usage: "A JSON array of objects containing track IDs with their corresponding timestamps. Each object must include a track ID and an `added_at` timestamp. This allows you to specify when tracks were added to maintain a specific chronological order in the user's library.<br/>A maximum of 50 items can be specified in one request. _**Note**: if the `timestamped_ids` is present in the body, any IDs listed in the query parameters (deprecated) or the `ids` field in the body will be ignored._\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "timestamped_ids.#.added_at",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name:  "+timestamped-id",
			Usage: "A JSON array of objects containing track IDs with their corresponding timestamps. Each object must include a track ID and an `added_at` timestamp. This allows you to specify when tracks were added to maintain a specific chronological order in the user's library.<br/>A maximum of 50 items can be specified in one request. _**Note**: if the `timestamped_ids` is present in the body, any IDs listed in the query parameters (deprecated) or the `ids` field in the body will be ignored._\n",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "timestamped_ids.-1",
				SetValue: map[string]interface{}{},
			},
		},
	},
	Action:          handleMeTracksSave,
	HideHelpCommand: true,
}

func handleMeTracksList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeTrackListParams{}
	var res []byte
	_, err := cc.client.Me.Tracks.List(
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
	return ShowJSON("me:tracks list", json, format, transform)
}

func handleMeTracksCheck(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeTrackCheckParams{}
	var res []byte
	_, err := cc.client.Me.Tracks.Check(
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
	return ShowJSON("me:tracks check", json, format, transform)
}

func handleMeTracksRemove(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeTrackRemoveParams{}
	return cc.client.Me.Tracks.Remove(
		ctx,
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
}

func handleMeTracksSave(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeTrackSaveParams{}
	return cc.client.Me.Tracks.Save(
		ctx,
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
}
