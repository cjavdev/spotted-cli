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

var usersPlaylistsCreate = cli.Command{
	Name:  "create",
	Usage: "Create a playlist for a Spotify user. (The playlist will be empty until you\n[add tracks](/documentation/web-api/reference/add-tracks-to-playlist).) Each\nuser is generally limited to a maximum of 11000 playlists.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "user-id",
			Usage: "The user's [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids).\n",
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "The name for the new playlist, for example `\"Your Coolest Playlist\"`. This name does not need to be unique; a user may have several playlists with the same name.\n",
		},
		&cli.BoolFlag{
			Name:  "components-schemas-properties-published",
			Usage: "Defaults to `true`. The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private. To be able to create private playlists, the user must have granted the `playlist-modify-private` [scope](/documentation/web-api/concepts/scopes/#list-of-scopes). For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)\n",
		},
		&cli.BoolFlag{
			Name:  "collaborative",
			Usage: "Defaults to `false`. If `true` the playlist will be collaborative. _**Note**: to create a collaborative playlist you must also set `public` to `false`. To create collaborative playlists you must have granted `playlist-modify-private` and `playlist-modify-public` [scopes](/documentation/web-api/concepts/scopes/#list-of-scopes)._\n",
		},
		&cli.StringFlag{
			Name:  "description",
			Usage: "value for playlist description as displayed in Spotify Clients and in the Web API.\n",
		},
	},
	Action:          handleUsersPlaylistsCreate,
	HideHelpCommand: true,
}

var usersPlaylistsList = cli.Command{
	Name:  "list",
	Usage: "Get a list of the playlists owned or followed by a Spotify user.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "user-id",
			Usage: "The user's [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids).\n",
		},
		&cli.Int64Flag{
			Name:  "limit",
			Usage: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Value: 20,
		},
		&cli.Int64Flag{
			Name:  "offset",
			Usage: "The index of the first playlist to return. Default:\n0 (the first object). Maximum offset: 100.000\\. Use with `limit` to get the\nnext set of playlists.\n",
		},
	},
	Action:          handleUsersPlaylistsList,
	HideHelpCommand: true,
}

func handleUsersPlaylistsCreate(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.UserPlaylistNewParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"name": "name",
		"components-schemas-properties-published": "$\\.components\\.schemas\\.*\\.properties\\.published",
		"collaborative": "collaborative",
		"description":   "description",
	}, &params); err != nil {
		return err
	}
	var res []byte
	_, err := client.Users.Playlists.New(
		ctx,
		cmd.Value("user-id").(string),
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
	return ShowJSON("users:playlists create", json, format, transform)
}

func handleUsersPlaylistsList(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.UserPlaylistListParams{}
	if cmd.IsSet("limit") {
		params.Limit = spotted.Opt(cmd.Value("limit").(int64))
	}
	if cmd.IsSet("offset") {
		params.Offset = spotted.Opt(cmd.Value("offset").(int64))
	}
	var res []byte
	_, err := client.Users.Playlists.List(
		ctx,
		cmd.Value("user-id").(string),
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
	return ShowJSON("users:playlists list", json, format, transform)
}
