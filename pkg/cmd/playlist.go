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

var playlistsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get a playlist owned by a Spotify user.",
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
			Usage: "Filters for the query: a comma-separated list of the\nfields to return. If omitted, all fields are returned. For example, to get\njust the playlist''s description and URI: `fields=description,uri`. A dot\nseparator can be used to specify non-reoccurring fields, while parentheses\ncan be used to specify reoccurring fields within objects. For example, to\nget just the added date and user ID of the adder: `fields=tracks.items(added_at,added_by.id)`.\nUse multiple parentheses to drill down into nested objects, for example: `fields=tracks.items(track(name,href,album(name,href)))`.\nFields can be excluded by prefixing them with an exclamation mark, for example:\n`fields=tracks.items(track(name,href,album(!name,href)))`\n",
		},
		&cli.StringFlag{
			Name:  "market",
			Usage: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
		},
	},
	Action:          handlePlaylistsRetrieve,
	HideHelpCommand: true,
}

var playlistsUpdate = cli.Command{
	Name:  "update",
	Usage: "Change a playlist's name and public/private state. (The user must, of course,\nown the playlist.)",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "playlist-id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
		},
		&cli.BoolFlag{
			Name:  "collaborative",
			Usage: "If `true`, the playlist will become collaborative and other users will be able to modify the playlist in their Spotify client. <br/>\n_**Note**: You can only set `collaborative` to `true` on non-public playlists._\n",
		},
		&cli.StringFlag{
			Name:  "description",
			Usage: "Value for playlist description as displayed in Spotify Clients and in the Web API.\n",
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "The new name for the playlist, for example `\"My New Playlist Title\"`\n",
		},
		&cli.BoolFlag{
			Name:  "public",
			Usage: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)\n",
		},
	},
	Action:          handlePlaylistsUpdate,
	HideHelpCommand: true,
}

func handlePlaylistsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("playlist-id") && len(unusedArgs) > 0 {
		cmd.Set("playlist-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.PlaylistGetParams{
		AdditionalTypes: spotted.String(cmd.Value("additional-types").(string)),
		Fields:          spotted.String(cmd.Value("fields").(string)),
		Market:          spotted.String(cmd.Value("market").(string)),
	}
	var res []byte
	_, err := client.Playlists.Get(
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
	return ShowJSON("playlists retrieve", json, format, transform)
}

func handlePlaylistsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("playlist-id") && len(unusedArgs) > 0 {
		cmd.Set("playlist-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.PlaylistUpdateParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"collaborative": "collaborative",
		"description":   "description",
		"name":          "name",
		"public":        "public",
	}, &params); err != nil {
		return err
	}
	return client.Playlists.Update(
		ctx,
		cmd.Value("playlist-id").(string),
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}
