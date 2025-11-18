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

var playlistsTracksUpdate = cli.Command{
	Name:  "update",
	Usage: "Either reorder or replace items in a playlist depending on the request's\nparameters. To reorder items, include `range_start`, `insert_before`,\n`range_length` and `snapshot_id` in the request's body. To replace items,\ninclude `uris` as either a query parameter or in the request's body. Replacing\nitems in a playlist will overwrite its existing items. This operation can be\nused for replacing or clearing items in a playlist. <br/> **Note**: Replace and\nreorder are mutually exclusive operations which share the same endpoint, but\nhave different parameters. These operations can't be applied together in a\nsingle request.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "playlist-id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
		},
		&cli.Int64Flag{
			Name:  "insert-before",
			Usage: "The position where the items should be inserted.<br/>To reorder the items to the end of the playlist, simply set _insert_before_ to the position after the last item.<br/>Examples:<br/>To reorder the first item to the last position in a playlist with 10 items, set _range_start_ to 0, and _insert_before_ to 10.<br/>To reorder the last item in a playlist with 10 items to the start of the playlist, set _range_start_ to 9, and _insert_before_ to 0.\n",
		},
		&cli.Int64Flag{
			Name:  "range-length",
			Usage: "The amount of items to be reordered. Defaults to 1 if not set.<br/>The range of items to be reordered begins from the _range_start_ position, and includes the _range_length_ subsequent items.<br/>Example:<br/>To move the items at index 9-10 to the start of the playlist, _range_start_ is set to 9, and _range_length_ is set to 2.\n",
		},
		&cli.Int64Flag{
			Name:  "range-start",
			Usage: "The position of the first item to be reordered.\n",
		},
		&cli.StringFlag{
			Name:  "snapshot-id",
			Usage: "The playlist's snapshot ID against which you want to make the changes.\n",
		},
		&cli.StringSliceFlag{
			Name: "uris",
		},
	},
	Action:          handlePlaylistsTracksUpdate,
	HideHelpCommand: true,
}

var playlistsTracksList = cli.Command{
	Name:  "list",
	Usage: "Get full details of the items of a playlist owned by a Spotify user.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "playlist-id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
		},
		&cli.StringFlag{
			Name:  "additional-types",
			Usage: "A comma-separated list of item types that your client supports besides the default `track` type. Valid types are: `track` and `episode`.<br/>\n_**Note**: This parameter was introduced to allow existing clients to maintain their current behaviour and might be deprecated in the future._<br/>\nIn addition to providing this parameter, make sure that your client properly handles cases of new types in the future by checking against the `type` field of each object.\n",
		},
		&cli.StringFlag{
			Name:  "fields",
			Usage: "Filters for the query: a comma-separated list of the\nfields to return. If omitted, all fields are returned. For example, to get\njust the total number of items and the request limit:<br/>`fields=total,limit`<br/>A\ndot separator can be used to specify non-reoccurring fields, while parentheses\ncan be used to specify reoccurring fields within objects. For example, to\nget just the added date and user ID of the adder:<br/>`fields=items(added_at,added_by.id)`<br/>Use\nmultiple parentheses to drill down into nested objects, for example:<br/>`fields=items(track(name,href,album(name,href)))`<br/>Fields\ncan be excluded by prefixing them with an exclamation mark, for example:<br/>`fields=items.track.album(!external_urls,images)`\n",
		},
		&cli.Int64Flag{
			Name:  "limit",
			Usage: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 100.\n",
			Value: 20,
		},
		&cli.StringFlag{
			Name:  "market",
			Usage: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
		},
		&cli.Int64Flag{
			Name:  "offset",
			Usage: "The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.\n",
		},
	},
	Action:          handlePlaylistsTracksList,
	HideHelpCommand: true,
}

var playlistsTracksAdd = cli.Command{
	Name:  "add",
	Usage: "Add one or more items to a user's playlist.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "playlist-id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
		},
		&cli.Int64Flag{
			Name:  "position",
			Usage: "The position to insert the items, a zero-based index. For example, to insert the items in the first position: `position=0` ; to insert the items in the third position: `position=2`. If omitted, the items will be appended to the playlist. Items are added in the order they appear in the uris array. For example: `{\"uris\": [\"spotify:track:4iV5W9uYEdYUVa79Axb7Rh\",\"spotify:track:1301WleyT98MSxVHPZCA6M\"], \"position\": 3}`\n",
		},
		&cli.StringSliceFlag{
			Name:  "uris",
			Usage: "A JSON array of the [Spotify URIs](/documentation/web-api/concepts/spotify-uris-ids) to add. For example: `{\"uris\": [\"spotify:track:4iV5W9uYEdYUVa79Axb7Rh\",\"spotify:track:1301WleyT98MSxVHPZCA6M\", \"spotify:episode:512ojhOuo1ktJprKbVcKyQ\"]}`<br/>A maximum of 100 items can be added in one request. _**Note**: if the `uris` parameter is present in the query string, any URIs listed here in the body will be ignored._\n",
		},
	},
	Action:          handlePlaylistsTracksAdd,
	HideHelpCommand: true,
}

var playlistsTracksRemove = cli.Command{
	Name:  "remove",
	Usage: "Remove one or more items from a user's playlist.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "playlist-id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
		},
		&cli.StringFlag{
			Name:  "snapshot-id",
			Usage: "The playlist's snapshot ID against which you want to make the changes.\nThe API will validate that the specified items exist and in the specified positions and make the changes,\neven if more recent changes have been made to the playlist.\n",
		},
	},
	Action:          handlePlaylistsTracksRemove,
	HideHelpCommand: true,
}

func handlePlaylistsTracksUpdate(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("playlist-id") && len(unusedArgs) > 0 {
		cmd.Set("playlist-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.PlaylistTrackUpdateParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"insert-before": "insert_before",
		"range-length":  "range_length",
		"range-start":   "range_start",
		"snapshot-id":   "snapshot_id",
		"uris":          "uris",
	}, &params); err != nil {
		return err
	}
	var res []byte
	_, err := client.Playlists.Tracks.Update(
		ctx,
		cmd.Value("playlist-id").(string),
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
	return ShowJSON("playlists:tracks update", json, format, transform)
}

func handlePlaylistsTracksList(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("playlist-id") && len(unusedArgs) > 0 {
		cmd.Set("playlist-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.PlaylistTrackListParams{
		AdditionalTypes: spotted.String(cmd.Value("additional-types").(string)),
		Fields:          spotted.String(cmd.Value("fields").(string)),
		Market:          spotted.String(cmd.Value("market").(string)),
	}
	if cmd.IsSet("limit") {
		params.Limit = spotted.Opt(cmd.Value("limit").(int64))
	}
	if cmd.IsSet("offset") {
		params.Offset = spotted.Opt(cmd.Value("offset").(int64))
	}
	var res []byte
	_, err := client.Playlists.Tracks.List(
		ctx,
		cmd.Value("playlist-id").(string),
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
	return ShowJSON("playlists:tracks list", json, format, transform)
}

func handlePlaylistsTracksAdd(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("playlist-id") && len(unusedArgs) > 0 {
		cmd.Set("playlist-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.PlaylistTrackAddParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"position": "position",
		"uris":     "uris",
	}, &params); err != nil {
		return err
	}
	var res []byte
	_, err := client.Playlists.Tracks.Add(
		ctx,
		cmd.Value("playlist-id").(string),
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
	return ShowJSON("playlists:tracks add", json, format, transform)
}

func handlePlaylistsTracksRemove(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("playlist-id") && len(unusedArgs) > 0 {
		cmd.Set("playlist-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.PlaylistTrackRemoveParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"snapshot-id": "snapshot_id",
	}, &params); err != nil {
		return err
	}
	var res []byte
	_, err := client.Playlists.Tracks.Remove(
		ctx,
		cmd.Value("playlist-id").(string),
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
	return ShowJSON("playlists:tracks remove", json, format, transform)
}
