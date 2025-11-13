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

var meFollowingBulkRetrieve = cli.Command{
	Name:  "bulk-retrieve",
	Usage: "Get the current user's followed artists.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "type",
			Usage: "The ID type: currently only `artist` is supported.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "type",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "after",
			Usage: "The last artist ID retrieved from the previous request.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "after",
			},
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
	},
	Action:          handleMeFollowingBulkRetrieve,
	HideHelpCommand: true,
}

var meFollowingCheck = cli.Command{
	Name:  "check",
	Usage: "Check to see if the current user is following one or more artists or other\nSpotify users.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the artist or the user [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) to check. For example: `ids=74ASZWbe4lXaubB36ztrGX,08td7MxkoHQkXnWAYD8d6Q`. A maximum of 50 IDs can be sent in one request.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "ids",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "type",
			Usage: "The ID type: either `artist` or `user`.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "type",
			},
		},
	},
	Action:          handleMeFollowingCheck,
	HideHelpCommand: true,
}

var meFollowingFollow = cli.Command{
	Name:  "follow",
	Usage: "Add the current user as a follower of one or more artists or other Spotify\nusers.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "ids",
			Usage: "A JSON array of the artist or user [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids).\nFor example: `{ids:[\"74ASZWbe4lXaubB36ztrGX\", \"08td7MxkoHQkXnWAYD8d6Q\"]}`. A maximum of 50 IDs can be sent in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "ids.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+id",
			Usage: "A JSON array of the artist or user [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids).\nFor example: `{ids:[\"74ASZWbe4lXaubB36ztrGX\", \"08td7MxkoHQkXnWAYD8d6Q\"]}`. A maximum of 50 IDs can be sent in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "ids.-1",
			},
		},
	},
	Action:          handleMeFollowingFollow,
	HideHelpCommand: true,
}

var meFollowingUnfollow = cli.Command{
	Name:  "unfollow",
	Usage: "Remove the current user as a follower of one or more artists or other Spotify\nusers.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "ids",
			Usage: "A JSON array of the artist or user [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `{ids:[\"74ASZWbe4lXaubB36ztrGX\", \"08td7MxkoHQkXnWAYD8d6Q\"]}`. A maximum of 50 IDs can be sent in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "ids.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+id",
			Usage: "A JSON array of the artist or user [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `{ids:[\"74ASZWbe4lXaubB36ztrGX\", \"08td7MxkoHQkXnWAYD8d6Q\"]}`. A maximum of 50 IDs can be sent in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "ids.-1",
			},
		},
	},
	Action:          handleMeFollowingUnfollow,
	HideHelpCommand: true,
}

func handleMeFollowingBulkRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeFollowingBulkGetParams{}
	var res []byte
	_, err := cc.client.Me.Following.BulkGet(
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
	return ShowJSON("me:following bulk-retrieve", json, format, transform)
}

func handleMeFollowingCheck(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeFollowingCheckParams{}
	var res []byte
	_, err := cc.client.Me.Following.Check(
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
	return ShowJSON("me:following check", json, format, transform)
}

func handleMeFollowingFollow(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeFollowingFollowParams{}
	return cc.client.Me.Following.Follow(
		ctx,
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
}

func handleMeFollowingUnfollow(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeFollowingUnfollowParams{}
	return cc.client.Me.Following.Unfollow(
		ctx,
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
}
