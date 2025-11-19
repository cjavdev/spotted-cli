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

var meAudiobooksList = cli.Command{
	Name:  "list",
	Usage: "Get a list of the audiobooks saved in the current Spotify user's 'Your Music'\nlibrary.",
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:  "limit",
			Usage: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Value: 20,
		},
		&cli.Int64Flag{
			Name:  "offset",
			Usage: "The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.\n",
		},
	},
	Action:          handleMeAudiobooksList,
	HideHelpCommand: true,
}

var meAudiobooksCheck = cli.Command{
	Name:  "check",
	Usage: "Check if one or more audiobooks are already saved in the current Spotify user's\nlibrary.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `ids=18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ`. Maximum: 50 IDs.\n",
		},
	},
	Action:          handleMeAudiobooksCheck,
	HideHelpCommand: true,
}

var meAudiobooksRemove = cli.Command{
	Name:  "remove",
	Usage: "Remove one or more audiobooks from the Spotify user's library.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `ids=18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ`. Maximum: 50 IDs.\n",
		},
	},
	Action:          handleMeAudiobooksRemove,
	HideHelpCommand: true,
}

var meAudiobooksSave = cli.Command{
	Name:  "save",
	Usage: "Save one or more audiobooks to the current Spotify user's library.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "ids",
			Usage: "A comma-separated list of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `ids=18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ`. Maximum: 50 IDs.\n",
		},
	},
	Action:          handleMeAudiobooksSave,
	HideHelpCommand: true,
}

func handleMeAudiobooksList(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeAudiobookListParams{}
	if cmd.IsSet("limit") {
		params.Limit = spotted.Opt(cmd.Value("limit").(int64))
	}
	if cmd.IsSet("offset") {
		params.Offset = spotted.Opt(cmd.Value("offset").(int64))
	}
	var res []byte
	_, err := client.Me.Audiobooks.List(
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
	return ShowJSON("me:audiobooks list", json, format, transform)
}

func handleMeAudiobooksCheck(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeAudiobookCheckParams{
		IDs: cmd.Value("ids").(string),
	}
	var res []byte
	_, err := client.Me.Audiobooks.Check(
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
	return ShowJSON("me:audiobooks check", json, format, transform)
}

func handleMeAudiobooksRemove(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeAudiobookRemoveParams{
		IDs: cmd.Value("ids").(string),
	}
	return client.Me.Audiobooks.Remove(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}

func handleMeAudiobooksSave(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MeAudiobookSaveParams{
		IDs: cmd.Value("ids").(string),
	}
	return client.Me.Audiobooks.Save(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}
