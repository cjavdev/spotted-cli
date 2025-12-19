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

var mePlayerGetCurrentlyPlaying = cli.Command{
	Name:  "get-currently-playing",
	Usage: "Get the object currently being played on the user's Spotify account.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "additional-types",
			Usage:     "A comma-separated list of item types that your client supports besides the default `track` type. Valid types are: `track` and `episode`.<br/>\n_**Note**: This parameter was introduced to allow existing clients to maintain their current behaviour and might be deprecated in the future._<br/>\nIn addition to providing this parameter, make sure that your client properly handles cases of new types in the future by checking against the `type` field of each object.\n",
			QueryPath: "additional_types",
		},
		&requestflag.Flag[string]{
			Name:      "market",
			Usage:     "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
			QueryPath: "market",
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
		&requestflag.Flag[string]{
			Name:      "additional-types",
			Usage:     "A comma-separated list of item types that your client supports besides the default `track` type. Valid types are: `track` and `episode`.<br/>\n_**Note**: This parameter was introduced to allow existing clients to maintain their current behaviour and might be deprecated in the future._<br/>\nIn addition to providing this parameter, make sure that your client properly handles cases of new types in the future by checking against the `type` field of each object.\n",
			QueryPath: "additional_types",
		},
		&requestflag.Flag[string]{
			Name:      "market",
			Usage:     "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
			QueryPath: "market",
		},
	},
	Action:          handleMePlayerGetState,
	HideHelpCommand: true,
}

var mePlayerListRecentlyPlayed = cli.Command{
	Name:  "list-recently-played",
	Usage: "Get tracks from the current user's recently played tracks. _**Note**: Currently\ndoesn't support podcast episodes._",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "after",
			Usage:     "A Unix timestamp in milliseconds. Returns all items\nafter (but not including) this cursor position. If `after` is specified, `before`\nmust not be specified.\n",
			QueryPath: "after",
		},
		&requestflag.Flag[int64]{
			Name:      "before",
			Usage:     "A Unix timestamp in milliseconds. Returns all items\nbefore (but not including) this cursor position. If `before` is specified,\n`after` must not be specified.\n",
			QueryPath: "before",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Default:   20,
			QueryPath: "limit",
		},
	},
	Action:          handleMePlayerListRecentlyPlayed,
	HideHelpCommand: true,
}

var mePlayerPausePlayback = cli.Command{
	Name:  "pause-playback",
	Usage: "Pause playback on the user's account. This API only works for users who have\nSpotify Premium. The order of execution is not guaranteed when you use this API\nwith other Player API endpoints.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "device-id",
			Usage:     "The id of the device this command is targeting. If not supplied, the user's currently active device is the target.\n",
			QueryPath: "device_id",
		},
	},
	Action:          handleMePlayerPausePlayback,
	HideHelpCommand: true,
}

var mePlayerSeekToPosition = cli.Command{
	Name:  "seek-to-position",
	Usage: "Seeks to the given position in the user’s currently playing track. This API only\nworks for users who have Spotify Premium. The order of execution is not\nguaranteed when you use this API with other Player API endpoints.",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "position-ms",
			Usage:     "The position in milliseconds to seek to. Must be a\npositive number. Passing in a position that is greater than the length of\nthe track will cause the player to start playing the next song.\n",
			QueryPath: "position_ms",
		},
		&requestflag.Flag[string]{
			Name:      "device-id",
			Usage:     "The id of the device this command is targeting. If\nnot supplied, the user's currently active device is the target.\n",
			QueryPath: "device_id",
		},
	},
	Action:          handleMePlayerSeekToPosition,
	HideHelpCommand: true,
}

var mePlayerSetRepeatMode = cli.Command{
	Name:  "set-repeat-mode",
	Usage: "Set the repeat mode for the user's playback. This API only works for users who\nhave Spotify Premium. The order of execution is not guaranteed when you use this\nAPI with other Player API endpoints.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     "**track**, **context** or **off**.<br/>\n**track** will repeat the current track.<br/>\n**context** will repeat the current context.<br/>\n**off** will turn repeat off.\n",
			QueryPath: "state",
		},
		&requestflag.Flag[string]{
			Name:      "device-id",
			Usage:     "The id of the device this command is targeting. If\nnot supplied, the user's currently active device is the target.\n",
			QueryPath: "device_id",
		},
	},
	Action:          handleMePlayerSetRepeatMode,
	HideHelpCommand: true,
}

var mePlayerSetVolume = cli.Command{
	Name:  "set-volume",
	Usage: "Set the volume for the user’s current playback device. This API only works for\nusers who have Spotify Premium. The order of execution is not guaranteed when\nyou use this API with other Player API endpoints.",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "volume-percent",
			Usage:     "The volume to set. Must be a value from 0 to 100 inclusive.\n",
			QueryPath: "volume_percent",
		},
		&requestflag.Flag[string]{
			Name:      "device-id",
			Usage:     "The id of the device this command is targeting. If not supplied, the user's currently active device is the target.\n",
			QueryPath: "device_id",
		},
	},
	Action:          handleMePlayerSetVolume,
	HideHelpCommand: true,
}

var mePlayerSkipNext = cli.Command{
	Name:  "skip-next",
	Usage: "Skips to next track in the user’s queue. This API only works for users who have\nSpotify Premium. The order of execution is not guaranteed when you use this API\nwith other Player API endpoints.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "device-id",
			Usage:     "The id of the device this command is targeting. If not supplied, the user's currently active device is the target.",
			QueryPath: "device_id",
		},
	},
	Action:          handleMePlayerSkipNext,
	HideHelpCommand: true,
}

var mePlayerSkipPrevious = cli.Command{
	Name:  "skip-previous",
	Usage: "Skips to previous track in the user’s queue. This API only works for users who\nhave Spotify Premium. The order of execution is not guaranteed when you use this\nAPI with other Player API endpoints.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "device-id",
			Usage:     "The id of the device this command is targeting. If\nnot supplied, the user's currently active device is the target.\n",
			QueryPath: "device_id",
		},
	},
	Action:          handleMePlayerSkipPrevious,
	HideHelpCommand: true,
}

var mePlayerStartPlayback = cli.Command{
	Name:  "start-playback",
	Usage: "Start a new context or resume current playback on the user's active device. This\nAPI only works for users who have Spotify Premium. The order of execution is not\nguaranteed when you use this API with other Player API endpoints.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "device-id",
			Usage:     "The id of the device this command is targeting. If not supplied, the user's currently active device is the target.",
			QueryPath: "device_id",
		},
		&requestflag.Flag[string]{
			Name:     "context-uri",
			Usage:    "Optional. Spotify URI of the context to play.\nValid contexts are albums, artists & playlists.\n`{context_uri:\"spotify:album:1Je1IMUlBXcx1Fz0WE7oPT\"}`\n",
			BodyPath: "context_uri",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "offset",
			Usage:    "Optional. Indicates from where in the context playback should start. Only available when context_uri corresponds to an album or playlist object\n\"position\" is zero based and can’t be negative. Example: `\"offset\": {\"position\": 5}`\n\"uri\" is a string representing the uri of the item to start at. Example: `\"offset\": {\"uri\": \"spotify:track:1301WleyT98MSxVHPZCA6M\"}`\n",
			BodyPath: "offset",
		},
		&requestflag.Flag[int64]{
			Name:     "position-ms",
			Usage:    "Indicates from what position to start playback. Must be a positive number. Passing in a position that is greater than the length of the track will cause the player to start playing the next song.\n",
			BodyPath: "position_ms",
		},
		&requestflag.Flag[bool]{
			Name:     "published",
			Usage:    "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)\n",
			BodyPath: "published",
		},
		&requestflag.Flag[[]string]{
			Name:     "uris",
			Usage:    "Optional. A JSON array of the Spotify track URIs to play.\nFor example: `{\"uris\": [\"spotify:track:4iV5W9uYEdYUVa79Axb7Rh\", \"spotify:track:1301WleyT98MSxVHPZCA6M\"]}`\n",
			BodyPath: "uris",
		},
	},
	Action:          handleMePlayerStartPlayback,
	HideHelpCommand: true,
}

var mePlayerToggleShuffle = cli.Command{
	Name:  "toggle-shuffle",
	Usage: "Toggle shuffle on or off for user’s playback. This API only works for users who\nhave Spotify Premium. The order of execution is not guaranteed when you use this\nAPI with other Player API endpoints.",
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:      "state",
			Usage:     "**true** : Shuffle user's playback.<br/>\n**false** : Do not shuffle user's playback.\n",
			QueryPath: "state",
		},
		&requestflag.Flag[string]{
			Name:      "device-id",
			Usage:     "The id of the device this command is targeting. If\nnot supplied, the user's currently active device is the target.\n",
			QueryPath: "device_id",
		},
	},
	Action:          handleMePlayerToggleShuffle,
	HideHelpCommand: true,
}

var mePlayerTransfer = cli.Command{
	Name:  "transfer",
	Usage: "Transfer playback to a new device and optionally begin playback. This API only\nworks for users who have Spotify Premium. The order of execution is not\nguaranteed when you use this API with other Player API endpoints.",
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "device-id",
			Usage:    "A JSON array containing the ID of the device on which playback should be started/transferred.<br/>For example:`{device_ids:[\"74ASZWbe4lXaubB36ztrGX\"]}`<br/>_**Note**: Although an array is accepted, only a single device_id is currently supported. Supplying more than one will return `400 Bad Request`_\n",
			BodyPath: "device_ids",
		},
		&requestflag.Flag[bool]{
			Name:     "play",
			Usage:    "**true**: ensure playback happens on new device.<br/>**false** or not provided: keep the current playback state.\n",
			BodyPath: "play",
		},
		&requestflag.Flag[bool]{
			Name:     "published",
			Usage:    "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)\n",
			BodyPath: "published",
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

	params := spotted.MePlayerGetCurrentlyPlayingParams{}

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
	_, err = client.Me.Player.GetCurrentlyPlaying(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "me:player get-currently-playing", obj, format, transform)
}

func handleMePlayerGetDevices(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Me.Player.GetDevices(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "me:player get-devices", obj, format, transform)
}

func handleMePlayerGetState(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MePlayerGetStateParams{}

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
	_, err = client.Me.Player.GetState(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "me:player get-state", obj, format, transform)
}

func handleMePlayerListRecentlyPlayed(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MePlayerListRecentlyPlayedParams{}

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
		_, err = client.Me.Player.ListRecentlyPlayed(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "me:player list-recently-played", obj, format, transform)
	} else {
		iter := client.Me.Player.ListRecentlyPlayedAutoPaging(ctx, params, options...)
		return streamOutput("me:player list-recently-played", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "me:player list-recently-played", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleMePlayerPausePlayback(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MePlayerPausePlaybackParams{}

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

	return client.Me.Player.PausePlayback(ctx, params, options...)
}

func handleMePlayerSeekToPosition(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MePlayerSeekToPositionParams{}

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

	return client.Me.Player.SeekToPosition(ctx, params, options...)
}

func handleMePlayerSetRepeatMode(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MePlayerSetRepeatModeParams{}

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

	return client.Me.Player.SetRepeatMode(ctx, params, options...)
}

func handleMePlayerSetVolume(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MePlayerSetVolumeParams{}

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

	return client.Me.Player.SetVolume(ctx, params, options...)
}

func handleMePlayerSkipNext(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MePlayerSkipNextParams{}

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

	return client.Me.Player.SkipNext(ctx, params, options...)
}

func handleMePlayerSkipPrevious(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MePlayerSkipPreviousParams{}

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

	return client.Me.Player.SkipPrevious(ctx, params, options...)
}

func handleMePlayerStartPlayback(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MePlayerStartPlaybackParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.Me.Player.StartPlayback(ctx, params, options...)
}

func handleMePlayerToggleShuffle(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MePlayerToggleShuffleParams{}

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

	return client.Me.Player.ToggleShuffle(ctx, params, options...)
}

func handleMePlayerTransfer(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := spotted.MePlayerTransferParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.Me.Player.Transfer(ctx, params, options...)
}
