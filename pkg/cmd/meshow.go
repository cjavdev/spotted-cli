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

var meShowsList = cli.Command{
	Name:  "list",
	Usage: "Get a list of shows saved in the current Spotify user's library. Optional\nparameters can be used to limit the number of shows returned.",
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
	Action:          handleMeShowsList,
	HideHelpCommand: true,
}

var meShowsCheck = cli.Command{
	Name:  "check",
	Usage: "Check if one or more shows is already saved in the current Spotify user's\nlibrary.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for the shows. Maximum: 50 IDs.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "ids",
			},
		},
	},
	Action:          handleMeShowsCheck,
	HideHelpCommand: true,
}

var meShowsRemove = cli.Command{
	Name:  "remove",
	Usage: "Delete one or more shows from current Spotify user's library.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "ids",
			Usage: "A JSON array of the [Spotify IDs](https://developer.spotify.com/documentation/web-api/#spotify-uris-and-ids).  \nA maximum of 50 items can be specified in one request. *Note: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored.*",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "ids.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+id",
			Usage: "A JSON array of the [Spotify IDs](https://developer.spotify.com/documentation/web-api/#spotify-uris-and-ids).  \nA maximum of 50 items can be specified in one request. *Note: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored.*",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "ids.-1",
			},
		},
	},
	Action:          handleMeShowsRemove,
	HideHelpCommand: true,
}

var meShowsSave = cli.Command{
	Name:  "save",
	Usage: "Save one or more shows to current Spotify user's library.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "ids",
			Usage: "A JSON array of the [Spotify IDs](https://developer.spotify.com/documentation/web-api/#spotify-uris-and-ids).  \nA maximum of 50 items can be specified in one request. *Note: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored.*",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "ids.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+id",
			Usage: "A JSON array of the [Spotify IDs](https://developer.spotify.com/documentation/web-api/#spotify-uris-and-ids).  \nA maximum of 50 items can be specified in one request. *Note: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored.*",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "ids.-1",
			},
		},
	},
	Action:          handleMeShowsSave,
	HideHelpCommand: true,
}

func handleMeShowsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeShowListParams{}
	var res []byte
	_, err := cc.client.Me.Shows.List(
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
	return ShowJSON("me:shows list", json, format, transform)
}

func handleMeShowsCheck(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeShowCheckParams{}
	var res []byte
	_, err := cc.client.Me.Shows.Check(
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
	return ShowJSON("me:shows check", json, format, transform)
}

func handleMeShowsRemove(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeShowRemoveParams{}
	return cc.client.Me.Shows.Remove(
		ctx,
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
}

func handleMeShowsSave(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeShowSaveParams{}
	return cc.client.Me.Shows.Save(
		ctx,
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
}
