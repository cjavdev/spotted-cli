// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cjavdev/spotted-cli/internal/apiquery"
	"github.com/cjavdev/spotted-cli/internal/requestflag"
	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var meFollowingBulkRetrieve = cli.Command{
	Name:  "bulk-retrieve",
	Usage: "Get the current user's followed artists.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "The ID type: currently only `artist` is supported.\n",
			QueryPath: "type",
		},
		&requestflag.Flag[string]{
			Name:      "after",
			Usage:     "The last artist ID retrieved from the previous request.\n",
			QueryPath: "after",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Default:   20,
			QueryPath: "limit",
		},
	},
	Action:          handleMeFollowingBulkRetrieve,
	HideHelpCommand: true,
}

var meFollowingCheck = cli.Command{
	Name:  "check",
	Usage: "Check to see if the current user is following one or more artists or other\nSpotify users.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ids",
			Usage:     "A comma-separated list of the artist or the user [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) to check. For example: `ids=74ASZWbe4lXaubB36ztrGX,08td7MxkoHQkXnWAYD8d6Q`. A maximum of 50 IDs can be sent in one request.\n",
			QueryPath: "ids",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "The ID type: either `artist` or `user`.\n",
			QueryPath: "type",
		},
	},
	Action:          handleMeFollowingCheck,
	HideHelpCommand: true,
}

var meFollowingFollow = cli.Command{
	Name:  "follow",
	Usage: "Add the current user as a follower of one or more artists or other Spotify\nusers.",
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "id",
			Usage:    "A JSON array of the artist or user [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids).\nFor example: `{ids:[\"74ASZWbe4lXaubB36ztrGX\", \"08td7MxkoHQkXnWAYD8d6Q\"]}`. A maximum of 50 IDs can be sent in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._\n",
			BodyPath: "ids",
		},
	},
	Action:          handleMeFollowingFollow,
	HideHelpCommand: true,
}

var meFollowingUnfollow = cli.Command{
	Name:  "unfollow",
	Usage: "Remove the current user as a follower of one or more artists or other Spotify\nusers.",
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "id",
			Usage:    "A JSON array of the artist or user [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `{ids:[\"74ASZWbe4lXaubB36ztrGX\", \"08td7MxkoHQkXnWAYD8d6Q\"]}`. A maximum of 50 IDs can be sent in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._\n",
			BodyPath: "ids",
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
	params := spotted.MeFollowingBulkGetParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Me.Following.BulkGet(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "me:following bulk-retrieve", obj, format, transform)
}

func handleMeFollowingCheck(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeFollowingCheckParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Me.Following.Check(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "me:following check", obj, format, transform)
}

func handleMeFollowingFollow(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeFollowingFollowParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	return client.Me.Following.Follow(ctx, params, options...)
}

func handleMeFollowingUnfollow(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeFollowingUnfollowParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	return client.Me.Following.Unfollow(ctx, params, options...)
}
