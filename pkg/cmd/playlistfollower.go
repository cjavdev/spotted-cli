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

var playlistsFollowersCheck = cli.Command{
	Name:  "check",
	Usage: "Check to see if the current user is following a specified playlist.",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "playlist-id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
		},
		&requestflag.StringFlag{
			Name:  "ids",
			Usage: "**Deprecated** A single item list containing current user's [Spotify Username](/documentation/web-api/concepts/spotify-uris-ids). Maximum: 1 id.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "ids",
			},
		},
	},
	Action:          handlePlaylistsFollowersCheck,
	HideHelpCommand: true,
}

var playlistsFollowersFollow = cli.Command{
	Name:  "follow",
	Usage: "Add the current user as a follower of a playlist.",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "playlist-id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
		},
		&requestflag.BoolFlag{
			Name:  "public",
			Usage: "Defaults to `true`. If `true` the playlist will be included in user's public playlists (added to profile), if `false` it will remain private. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)\n",
			Config: requestflag.RequestConfig{
				BodyPath: "public",
			},
		},
	},
	Action:          handlePlaylistsFollowersFollow,
	HideHelpCommand: true,
}

var playlistsFollowersUnfollow = cli.Command{
	Name:  "unfollow",
	Usage: "Remove the current user as a follower of a playlist.",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "playlist-id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
		},
	},
	Action:          handlePlaylistsFollowersUnfollow,
	HideHelpCommand: true,
}

func handlePlaylistsFollowersCheck(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("playlist-id") && len(unusedArgs) > 0 {
		cmd.Set("playlist-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.PlaylistFollowerCheckParams{}

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
	_, err = client.Playlists.Followers.Check(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "playlist-id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "playlists:followers check", obj, format, transform)
}

func handlePlaylistsFollowersFollow(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("playlist-id") && len(unusedArgs) > 0 {
		cmd.Set("playlist-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.PlaylistFollowerFollowParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	return client.Playlists.Followers.Follow(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "playlist-id"),
		params,
		options...,
	)
}

func handlePlaylistsFollowersUnfollow(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("playlist-id") && len(unusedArgs) > 0 {
		cmd.Set("playlist-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	return client.Playlists.Followers.Unfollow(ctx, requestflag.CommandRequestValue[string](cmd, "playlist-id"), options...)
}
