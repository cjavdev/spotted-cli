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

var meFollowingBulkRetrieve = cli.Command{
	Name:  "bulk-retrieve",
	Usage: "Get the current user's followed artists.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "type",
			Usage: "The ID type: currently only `artist` is supported.\n",
		},
		&cli.StringFlag{
			Name:  "after",
			Usage: "The last artist ID retrieved from the previous request.\n",
		},
		&cli.Int64Flag{
			Name:  "limit",
			Usage: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
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
		&cli.StringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the artist or the user [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) to check. For example: `ids=74ASZWbe4lXaubB36ztrGX,08td7MxkoHQkXnWAYD8d6Q`. A maximum of 50 IDs can be sent in one request.\n",
		},
		&cli.StringFlag{
			Name:  "type",
			Usage: "The ID type: either `artist` or `user`.\n",
		},
	},
	Action:          handleMeFollowingCheck,
	HideHelpCommand: true,
}

var meFollowingFollow = cli.Command{
	Name:  "follow",
	Usage: "Add the current user as a follower of one or more artists or other Spotify\nusers.",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:  "id",
			Usage: "A JSON array of the artist or user [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids).\nFor example: `{ids:[\"74ASZWbe4lXaubB36ztrGX\", \"08td7MxkoHQkXnWAYD8d6Q\"]}`. A maximum of 50 IDs can be sent in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._\n",
		},
	},
	Action:          handleMeFollowingFollow,
	HideHelpCommand: true,
}

var meFollowingUnfollow = cli.Command{
	Name:  "unfollow",
	Usage: "Remove the current user as a follower of one or more artists or other Spotify\nusers.",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:  "id",
			Usage: "A JSON array of the artist or user [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `{ids:[\"74ASZWbe4lXaubB36ztrGX\", \"08td7MxkoHQkXnWAYD8d6Q\"]}`. A maximum of 50 IDs can be sent in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._\n",
		},
	},
	Action:          handleMeFollowingUnfollow,
	HideHelpCommand: true,
}

func handleMeFollowingBulkRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeFollowingBulkGetParams{
		// Type:  cmd.Value("type").(spotted.MeFollowingBulkGetParamsType),
		After: spotted.String(cmd.Value("after").(string)),
	}
	if cmd.IsSet("limit") {
		params.Limit = spotted.Opt(cmd.Value("limit").(int64))
	}
	var res []byte
	_, err := client.Me.Following.BulkGet(
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
	return ShowJSON("me:following bulk-retrieve", json, format, transform)
}

func handleMeFollowingCheck(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeFollowingCheckParams{
		IDs:  cmd.Value("ids").(string),
		Type: cmd.Value("type").(spotted.MeFollowingCheckParamsType),
	}
	var res []byte
	_, err := client.Me.Following.Check(
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
	return ShowJSON("me:following check", json, format, transform)
}

func handleMeFollowingFollow(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeFollowingFollowParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"ids": "ids",
	}, &params); err != nil {
		return err
	}
	return client.Me.Following.Follow(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}

func handleMeFollowingUnfollow(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeFollowingUnfollowParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"ids": "ids",
	}, &params); err != nil {
		return err
	}
	return client.Me.Following.Unfollow(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}
