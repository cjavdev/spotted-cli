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

var mePlayerGetCurrentlyPlaying = cli.Command{
	Name:  "get-currently-playing",
	Usage: "Get the object currently being played on the user's Spotify account.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "additional-types",
			Usage: "A comma-separated list of item types that your client supports besides the default `track` type. Valid types are: `track` and `episode`.<br/>\n_**Note**: This parameter was introduced to allow existing clients to maintain their current behaviour and might be deprecated in the future._<br/>\nIn addition to providing this parameter, make sure that your client properly handles cases of new types in the future by checking against the `type` field of each object.\n",
		},
		&cli.StringFlag{
			Name:  "market",
			Usage: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
		},
	},
	Action:          handleMePlayerGetCurrentlyPlaying,
	HideHelpCommand: true,
}

var mePlayerGetDevices = cli.Command{
	Name:            "get-devices",
	Usage:           "Get information about a user’s available Spotify Connect devices. Some device\nmodels are not supported and will not be listed in the API response.",
	Flags:           []cli.Flag{},
	Action:          handleMePlayerGetDevices,
	HideHelpCommand: true,
}

var mePlayerGetState = cli.Command{
	Name:  "get-state",
	Usage: "Get information about the user’s current playback state, including track or\nepisode, progress, and active device.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "additional-types",
			Usage: "A comma-separated list of item types that your client supports besides the default `track` type. Valid types are: `track` and `episode`.<br/>\n_**Note**: This parameter was introduced to allow existing clients to maintain their current behaviour and might be deprecated in the future._<br/>\nIn addition to providing this parameter, make sure that your client properly handles cases of new types in the future by checking against the `type` field of each object.\n",
		},
		&cli.StringFlag{
			Name:  "market",
			Usage: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
		},
	},
	Action:          handleMePlayerGetState,
	HideHelpCommand: true,
}

var mePlayerListRecentlyPlayed = cli.Command{
	Name:  "list-recently-played",
	Usage: "Get tracks from the current user's recently played tracks. _**Note**: Currently\ndoesn't support podcast episodes._",
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:  "after",
			Usage: "A Unix timestamp in milliseconds. Returns all items\nafter (but not including) this cursor position. If `after` is specified, `before`\nmust not be specified.\n",
		},
		&cli.Int64Flag{
			Name:  "before",
			Usage: "A Unix timestamp in milliseconds. Returns all items\nbefore (but not including) this cursor position. If `before` is specified,\n`after` must not be specified.\n",
		},
		&cli.Int64Flag{
			Name:  "limit",
			Usage: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Value: 20,
		},
	},
	Action:          handleMePlayerListRecentlyPlayed,
	HideHelpCommand: true,
}

var mePlayerPausePlayback = cli.Command{
	Name:  "pause-playback",
	Usage: "Pause playback on the user's account. This API only works for users who have\nSpotify Premium. The order of execution is not guaranteed when you use this API\nwith other Player API endpoints.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "device-id",
			Usage: "The id of the device this command is targeting. If not supplied, the user's currently active device is the target.\n",
		},
	},
	Action:          handleMePlayerPausePlayback,
	HideHelpCommand: true,
}

var mePlayerSeekToPosition = cli.Command{
	Name:  "seek-to-position",
	Usage: "Seeks to the given position in the user’s currently playing track. This API only\nworks for users who have Spotify Premium. The order of execution is not\nguaranteed when you use this API with other Player API endpoints.",
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:  "position-ms",
			Usage: "The position in milliseconds to seek to. Must be a\npositive number. Passing in a position that is greater than the length of\nthe track will cause the player to start playing the next song.\n",
		},
		&cli.StringFlag{
			Name:  "device-id",
			Usage: "The id of the device this command is targeting. If\nnot supplied, the user's currently active device is the target.\n",
		},
	},
	Action:          handleMePlayerSeekToPosition,
	HideHelpCommand: true,
}

var mePlayerSetRepeatMode = cli.Command{
	Name:  "set-repeat-mode",
	Usage: "Set the repeat mode for the user's playback. This API only works for users who\nhave Spotify Premium. The order of execution is not guaranteed when you use this\nAPI with other Player API endpoints.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "state",
			Usage: "**track**, **context** or **off**.<br/>\n**track** will repeat the current track.<br/>\n**context** will repeat the current context.<br/>\n**off** will turn repeat off.\n",
		},
		&cli.StringFlag{
			Name:  "device-id",
			Usage: "The id of the device this command is targeting. If\nnot supplied, the user's currently active device is the target.\n",
		},
	},
	Action:          handleMePlayerSetRepeatMode,
	HideHelpCommand: true,
}

var mePlayerSetVolume = cli.Command{
	Name:  "set-volume",
	Usage: "Set the volume for the user’s current playback device. This API only works for\nusers who have Spotify Premium. The order of execution is not guaranteed when\nyou use this API with other Player API endpoints.",
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:  "volume-percent",
			Usage: "The volume to set. Must be a value from 0 to 100 inclusive.\n",
		},
		&cli.StringFlag{
			Name:  "device-id",
			Usage: "The id of the device this command is targeting. If not supplied, the user's currently active device is the target.\n",
		},
	},
	Action:          handleMePlayerSetVolume,
	HideHelpCommand: true,
}

var mePlayerSkipNext = cli.Command{
	Name:  "skip-next",
	Usage: "Skips to next track in the user’s queue. This API only works for users who have\nSpotify Premium. The order of execution is not guaranteed when you use this API\nwith other Player API endpoints.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "device-id",
			Usage: "The id of the device this command is targeting. If not supplied, the user's currently active device is the target.",
		},
	},
	Action:          handleMePlayerSkipNext,
	HideHelpCommand: true,
}

var mePlayerSkipPrevious = cli.Command{
	Name:  "skip-previous",
	Usage: "Skips to previous track in the user’s queue. This API only works for users who\nhave Spotify Premium. The order of execution is not guaranteed when you use this\nAPI with other Player API endpoints.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "device-id",
			Usage: "The id of the device this command is targeting. If\nnot supplied, the user's currently active device is the target.\n",
		},
	},
	Action:          handleMePlayerSkipPrevious,
	HideHelpCommand: true,
}

var mePlayerStartPlayback = cli.Command{
	Name:  "start-playback",
	Usage: "Start a new context or resume current playback on the user's active device. This\nAPI only works for users who have Spotify Premium. The order of execution is not\nguaranteed when you use this API with other Player API endpoints.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "device-id",
			Usage: "The id of the device this command is targeting. If not supplied, the user's currently active device is the target.",
		},
		&cli.StringFlag{
			Name:  "context-uri",
			Usage: "Optional. Spotify URI of the context to play.\nValid contexts are albums, artists & playlists.\n`{context_uri:\"spotify:album:1Je1IMUlBXcx1Fz0WE7oPT\"}`\n",
		},
		&cli.Int64Flag{
			Name:  "position-ms",
			Usage: "Indicates from what position to start playback. Must be a positive number. Passing in a position that is greater than the length of the track will cause the player to start playing the next song.\n",
		},
		&cli.StringSliceFlag{
			Name:  "uris",
			Usage: "Optional. A JSON array of the Spotify track URIs to play.\nFor example: `{\"uris\": [\"spotify:track:4iV5W9uYEdYUVa79Axb7Rh\", \"spotify:track:1301WleyT98MSxVHPZCA6M\"]}`\n",
		},
	},
	Action:          handleMePlayerStartPlayback,
	HideHelpCommand: true,
}

var mePlayerToggleShuffle = cli.Command{
	Name:  "toggle-shuffle",
	Usage: "Toggle shuffle on or off for user’s playback. This API only works for users who\nhave Spotify Premium. The order of execution is not guaranteed when you use this\nAPI with other Player API endpoints.",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "state",
			Usage: "**true** : Shuffle user's playback.<br/>\n**false** : Do not shuffle user's playback.\n",
		},
		&cli.StringFlag{
			Name:  "device-id",
			Usage: "The id of the device this command is targeting. If\nnot supplied, the user's currently active device is the target.\n",
		},
	},
	Action:          handleMePlayerToggleShuffle,
	HideHelpCommand: true,
}

var mePlayerTransfer = cli.Command{
	Name:  "transfer",
	Usage: "Transfer playback to a new device and optionally begin playback. This API only\nworks for users who have Spotify Premium. The order of execution is not\nguaranteed when you use this API with other Player API endpoints.",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:  "device-id",
			Usage: "A JSON array containing the ID of the device on which playback should be started/transferred.<br/>For example:`{device_ids:[\"74ASZWbe4lXaubB36ztrGX\"]}`<br/>_**Note**: Although an array is accepted, only a single device_id is currently supported. Supplying more than one will return `400 Bad Request`_\n",
		},
		&cli.BoolFlag{
			Name:  "play",
			Usage: "**true**: ensure playback happens on new device.<br/>**false** or not provided: keep the current playback state.\n",
		},
	},
	Action:          handleMePlayerTransfer,
	HideHelpCommand: true,
}

func handleMePlayerGetCurrentlyPlaying(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlayerGetCurrentlyPlayingParams{
		AdditionalTypes: spotted.String(cmd.Value("additional-types").(string)),
		Market:          spotted.String(cmd.Value("market").(string)),
	}
	var res []byte
	_, err := client.Me.Player.GetCurrentlyPlaying(
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
	return ShowJSON("me:player get-currently-playing", json, format, transform)
}

func handleMePlayerGetDevices(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := client.Me.Player.GetDevices(
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
	return ShowJSON("me:player get-devices", json, format, transform)
}

func handleMePlayerGetState(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlayerGetStateParams{
		AdditionalTypes: spotted.String(cmd.Value("additional-types").(string)),
		Market:          spotted.String(cmd.Value("market").(string)),
	}
	var res []byte
	_, err := client.Me.Player.GetState(
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
	return ShowJSON("me:player get-state", json, format, transform)
}

func handleMePlayerListRecentlyPlayed(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlayerListRecentlyPlayedParams{
		After:  spotted.Int(cmd.Value("after").(int64)),
		Before: spotted.Int(cmd.Value("before").(int64)),
	}
	if cmd.IsSet("limit") {
		params.Limit = spotted.Opt(cmd.Value("limit").(int64))
	}
	var res []byte
	_, err := client.Me.Player.ListRecentlyPlayed(
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
	return ShowJSON("me:player list-recently-played", json, format, transform)
}

func handleMePlayerPausePlayback(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlayerPausePlaybackParams{
		DeviceID: spotted.String(cmd.Value("device-id").(string)),
	}
	return client.Me.Player.PausePlayback(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}

func handleMePlayerSeekToPosition(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlayerSeekToPositionParams{
		PositionMs: cmd.Value("position-ms").(int64),
		DeviceID:   spotted.String(cmd.Value("device-id").(string)),
	}
	return client.Me.Player.SeekToPosition(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}

func handleMePlayerSetRepeatMode(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlayerSetRepeatModeParams{
		State:    cmd.Value("state").(string),
		DeviceID: spotted.String(cmd.Value("device-id").(string)),
	}
	return client.Me.Player.SetRepeatMode(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}

func handleMePlayerSetVolume(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlayerSetVolumeParams{
		VolumePercent: cmd.Value("volume-percent").(int64),
		DeviceID:      spotted.String(cmd.Value("device-id").(string)),
	}
	return client.Me.Player.SetVolume(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}

func handleMePlayerSkipNext(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlayerSkipNextParams{
		DeviceID: spotted.String(cmd.Value("device-id").(string)),
	}
	return client.Me.Player.SkipNext(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}

func handleMePlayerSkipPrevious(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlayerSkipPreviousParams{
		DeviceID: spotted.String(cmd.Value("device-id").(string)),
	}
	return client.Me.Player.SkipPrevious(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}

func handleMePlayerStartPlayback(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlayerStartPlaybackParams{
		DeviceID: spotted.String(cmd.Value("device-id").(string)),
	}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"context-uri": "context_uri",
		"position-ms": "position_ms",
		"uris":        "uris",
	}, &params); err != nil {
		return err
	}
	return client.Me.Player.StartPlayback(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}

func handleMePlayerToggleShuffle(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlayerToggleShuffleParams{
		State:    cmd.Value("state").(bool),
		DeviceID: spotted.String(cmd.Value("device-id").(string)),
	}
	return client.Me.Player.ToggleShuffle(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}

func handleMePlayerTransfer(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.MePlayerTransferParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"device-ids": "device_ids",
		"play":       "play",
	}, &params); err != nil {
		return err
	}
	return client.Me.Player.Transfer(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}
