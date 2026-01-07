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

var usersPlaylistsCreate = cli.Command{
	Name:  "create",
	Usage: "Create a playlist for a Spotify user. (The playlist will be empty until you\n[add tracks](/documentation/web-api/reference/add-tracks-to-playlist).) Each\nuser is generally limited to a maximum of 11000 playlists.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Usage:    "The user's [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids).\n",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name for the new playlist, for example `\"Your Coolest Playlist\"`. This name does not need to be unique; a user may have several playlists with the same name.\n",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "collaborative",
			Usage:    "Defaults to `false`. If `true` the playlist will be collaborative. _**Note**: to create a collaborative playlist you must also set `public` to `false`. To create collaborative playlists you must have granted `playlist-modify-private` and `playlist-modify-public` [scopes](/documentation/web-api/concepts/scopes/#list-of-scopes)._\n",
			BodyPath: "collaborative",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "value for playlist description as displayed in Spotify Clients and in the Web API.\n",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "published",
			Usage:    "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)\n",
			BodyPath: "published",
		},
	},
	Action:          handleUsersPlaylistsCreate,
	HideHelpCommand: true,
}

var usersPlaylistsList = cli.Command{
	Name:  "list",
	Usage: "Get a list of the playlists owned or followed by a Spotify user.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Usage:    "The user's [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids).\n",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "The index of the first playlist to return. Default:\n0 (the first object). Maximum offset: 100.000\\. Use with `limit` to get the\nnext set of playlists.\n",
			QueryPath: "offset",
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.Playlists.New(
		ctx,
		cmd.Value("user-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:playlists create", obj, format, transform)
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
		_, err = client.Users.Playlists.List(
			ctx,
			cmd.Value("user-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "users:playlists list", obj, format, transform)
	} else {
		iter := client.Users.Playlists.ListAutoPaging(
			ctx,
			cmd.Value("user-id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "users:playlists list", iter, format, transform)
	}
}
