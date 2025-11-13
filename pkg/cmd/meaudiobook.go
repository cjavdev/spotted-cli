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

var meAudiobooksList = cli.Command{
	Name:  "list",
	Usage: "Get a list of the audiobooks saved in the current Spotify user's 'Your Music'\nlibrary.",
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
	Action:          handleMeAudiobooksList,
	HideHelpCommand: true,
}

var meAudiobooksCheck = cli.Command{
	Name:  "check",
	Usage: "Check if one or more audiobooks are already saved in the current Spotify user's\nlibrary.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `ids=18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ`. Maximum: 50 IDs.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "ids",
			},
		},
	},
	Action:          handleMeAudiobooksCheck,
	HideHelpCommand: true,
}

var meAudiobooksRemove = cli.Command{
	Name:  "remove",
	Usage: "Remove one or more audiobooks from the Spotify user's library.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `ids=18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ`. Maximum: 50 IDs.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "ids",
			},
		},
	},
	Action:          handleMeAudiobooksRemove,
	HideHelpCommand: true,
}

var meAudiobooksSave = cli.Command{
	Name:  "save",
	Usage: "Save one or more audiobooks to the current Spotify user's library.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `ids=18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ`. Maximum: 50 IDs.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "ids",
			},
		},
	},
	Action:          handleMeAudiobooksSave,
	HideHelpCommand: true,
}

func handleMeAudiobooksList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeAudiobookListParams{}
	var res []byte
	_, err := cc.client.Me.Audiobooks.List(
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
	return ShowJSON("me:audiobooks list", json, format, transform)
}

func handleMeAudiobooksCheck(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeAudiobookCheckParams{}
	var res []byte
	_, err := cc.client.Me.Audiobooks.Check(
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
	return ShowJSON("me:audiobooks check", json, format, transform)
}

func handleMeAudiobooksRemove(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeAudiobookRemoveParams{}
	return cc.client.Me.Audiobooks.Remove(
		ctx,
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
}

func handleMeAudiobooksSave(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeAudiobookSaveParams{}
	return cc.client.Me.Audiobooks.Save(
		ctx,
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
}
