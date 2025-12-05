// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

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
		&requestflag.StringFlag{
			Name:  "user-id",
			Usage: "The user's [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids).\n",
		},
		&requestflag.StringFlag{
			Name:  "name",
			Usage: "The name for the new playlist, for example `\"Your Coolest Playlist\"`. This name does not need to be unique; a user may have several playlists with the same name.\n",
			Config: requestflag.RequestConfig{
				BodyPath: "name",
			},
		},
		&requestflag.BoolFlag{
			Name:  "collaborative",
			Usage: "Defaults to `false`. If `true` the playlist will be collaborative. _**Note**: to create a collaborative playlist you must also set `public` to `false`. To create collaborative playlists you must have granted `playlist-modify-private` and `playlist-modify-public` [scopes](/documentation/web-api/concepts/scopes/#list-of-scopes)._\n",
			Config: requestflag.RequestConfig{
				BodyPath: "collaborative",
			},
		},
		&requestflag.StringFlag{
			Name:  "description",
			Usage: "value for playlist description as displayed in Spotify Clients and in the Web API.\n",
			Config: requestflag.RequestConfig{
				BodyPath: "description",
			},
		},
		&requestflag.BoolFlag{
			Name:  "public",
			Usage: "Defaults to `true`. The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private. To be able to create private playlists, the user must have granted the `playlist-modify-private` [scope](/documentation/web-api/concepts/scopes/#list-of-scopes). For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)\n",
			Config: requestflag.RequestConfig{
				BodyPath: "public",
			},
		},
	},
	Action:          handleUsersPlaylistsCreate,
	HideHelpCommand: true,
}

var usersPlaylistsList = cli.Command{
	Name:  "list",
	Usage: "Get a list of the playlists owned or followed by a Spotify user.",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "user-id",
			Usage: "The user's [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids).\n",
		},
		&requestflag.IntFlag{
			Name:  "limit",
			Usage: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Value: requestflag.Value[int64](20),
			Config: requestflag.RequestConfig{
				QueryPath: "limit",
			},
		},
		&requestflag.IntFlag{
			Name:  "offset",
			Usage: "The index of the first playlist to return. Default:\n0 (the first object). Maximum offset: 100.000\\. Use with `limit` to get the\nnext set of playlists.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "offset",
			},
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
	)
	if err != nil {
		return err
	}
	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.Playlists.New(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "user-id"),
		params,
		options...,
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

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}
	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.Playlists.List(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "user-id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("users:playlists list", json, format, transform)
}
