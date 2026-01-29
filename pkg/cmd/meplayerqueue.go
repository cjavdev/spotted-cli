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

var mePlayerQueueAdd = cli.Command{
	Name:    "add",
	Usage:   "Add an item to be played next in the user's current playback queue. This API\nonly works for users who have Spotify Premium. The order of execution is not\nguaranteed when you use this API with other Player API endpoints.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "uri",
			Usage:     "The uri of the item to add to the queue. Must be a track or an episode uri.\n",
			Required:  true,
			QueryPath: "uri",
		},
		&requestflag.Flag[string]{
			Name:      "device-id",
			Usage:     "The id of the device this command is targeting. If\nnot supplied, the user's currently active device is the target.\n",
			QueryPath: "device_id",
		},
	},
	Action:          handleMePlayerQueueAdd,
	HideHelpCommand: true,
}

var mePlayerQueueGet = cli.Command{
	Name:            "get",
	Usage:           "Get the list of objects that make up the user's queue.",
	Suggest:         true,
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

	params := spotted.MePlayerQueueAddParams{}

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

	return client.Me.Player.Queue.Add(ctx, params, options...)
}

func handleMePlayerQueueGet(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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
	_, err = client.Me.Player.Queue.Get(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "me:player:queue get", obj, format, transform)
}
