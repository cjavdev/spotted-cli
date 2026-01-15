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

var meShowsList = cli.Command{
	Name:    "list",
	Usage:   "Get a list of shows saved in the current Spotify user's library. Optional\nparameters can be used to limit the number of shows returned.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.\n",
			Default:   0,
			QueryPath: "offset",
		},
	},
	Action:          handleMeShowsList,
	HideHelpCommand: true,
}

var meShowsCheck = cli.Command{
	Name:    "check",
	Usage:   "Check if one or more shows is already saved in the current Spotify user's\nlibrary.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ids",
			Usage:     "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for the shows. Maximum: 50 IDs.\n",
			Required:  true,
			QueryPath: "ids",
		},
	},
	Action:          handleMeShowsCheck,
	HideHelpCommand: true,
}

func handleMeShowsList(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MeShowListParams{}

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
		_, err = client.Me.Shows.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "me:shows list", obj, format, transform)
	} else {
		iter := client.Me.Shows.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "me:shows list", iter, format, transform)
	}
}

func handleMeShowsCheck(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MeShowCheckParams{}

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
	_, err = client.Me.Shows.Check(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "me:shows check", obj, format, transform)
}
