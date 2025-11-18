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

var mePlayerQueueAdd = cli.Command{
	Name:  "add",
	Usage: "Add an item to be played next in the user's current playback queue. This API\nonly works for users who have Spotify Premium. The order of execution is not\nguaranteed when you use this API with other Player API endpoints.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "uri",
			Usage: "The uri of the item to add to the queue. Must be a track or an episode uri.\n",
		},
		&cli.StringFlag{
			Name:  "device-id",
			Usage: "The id of the device this command is targeting. If\nnot supplied, the user's currently active device is the target.\n",
		},
	},
	Action:          handleMePlayerQueueAdd,
	HideHelpCommand: true,
}

var mePlayerQueueGet = cli.Command{
	Name:            "get",
	Usage:           "Get the list of objects that make up the user's queue.",
	Flags:           []cli.Flag{},
	Action:          handleMePlayerQueueGet,
	HideHelpCommand: true,
}

func handleMePlayerQueueAdd(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlayerQueueAddParams{
		Uri:      cmd.Value("uri").(string),
		DeviceID: spotted.String(cmd.Value("device-id").(string)),
	}
	return client.Me.Player.Queue.Add(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}

func handleMePlayerQueueGet(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := client.Me.Player.Queue.Get(
		ctx,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("me:player:queue get", json, format, transform)
}
