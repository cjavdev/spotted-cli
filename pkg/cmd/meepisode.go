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

var meEpisodesList = cli.Command{
	Name:  "list",
	Usage: "Get a list of the episodes saved in the current Spotify user's library.<br/>\nThis API endpoint is in **beta** and could change without warning. Please share\nany feedback that you have, or issues that you discover, in our\n[developer community forum](https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer).",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "market",
			Usage:     "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
			QueryPath: "market",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.\n",
			QueryPath: "offset",
		},
	},
	Action:          handleMeEpisodesList,
	HideHelpCommand: true,
}

var meEpisodesCheck = cli.Command{
	Name:  "check",
	Usage: "Check if one or more episodes is already saved in the current Spotify user's\n'Your Episodes' library.<br/> This API endpoint is in **beta** and could change\nwithout warning. Please share any feedback that you have, or issues that you\ndiscover, in our\n[developer community forum](https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer)..",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ids",
			Usage:     "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for the episodes. Maximum: 50 IDs.\n",
			Required:  true,
			QueryPath: "ids",
		},
	},
	Action:          handleMeEpisodesCheck,
	HideHelpCommand: true,
}

func handleMeEpisodesList(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MeEpisodeListParams{}

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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Me.Episodes.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "me:episodes list", obj, format, transform)
	} else {
		iter := client.Me.Episodes.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "me:episodes list", iter, format, transform)
	}
}

func handleMeEpisodesCheck(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MeEpisodeCheckParams{}

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
	_, err = client.Me.Episodes.Check(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "me:episodes check", obj, format, transform)
}
