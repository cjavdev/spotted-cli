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

var playlistsFollowersCheck = cli.Command{
	Name:  "check",
	Usage: "Check to see if the current user is following a specified playlist.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "playlist-id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
		},
		&cli.StringFlag{
			Name:  "ids",
			Usage: "**Deprecated** A single item list containing current user's [Spotify Username](/documentation/web-api/concepts/spotify-uris-ids). Maximum: 1 id.\n",
		},
	},
	Action:          handlePlaylistsFollowersCheck,
	HideHelpCommand: true,
}

var playlistsFollowersFollow = cli.Command{
	Name:  "follow",
	Usage: "Add the current user as a follower of a playlist.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "playlist-id",
			Usage: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.\n",
		},
		&cli.BoolFlag{
			Name:  "components-schemas-properties-published",
			Usage: "Defaults to `true`. If `true` the playlist will be included in user's public playlists (added to profile), if `false` it will remain private. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)\n",
		},
	},
	Action:          handlePlaylistsFollowersFollow,
	HideHelpCommand: true,
}

var playlistsFollowersUnfollow = cli.Command{
	Name:  "unfollow",
	Usage: "Remove the current user as a follower of a playlist.",
	Flags: []cli.Flag{
		&cli.StringFlag{
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
	params := spotted.PlaylistFollowerCheckParams{
		IDs: spotted.String(cmd.Value("ids").(string)),
	}
	var res []byte
	_, err := client.Playlists.Followers.Check(
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
	return ShowJSON("playlists:followers check", json, format, transform)
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
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"components-schemas-properties-published": "$\\.components\\.schemas\\.*\\.properties\\.published",
	}, &params); err != nil {
		return err
	}
	return client.Playlists.Followers.Follow(
		ctx,
		cmd.Value("playlist-id").(string),
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
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
	return client.Playlists.Followers.Unfollow(
		ctx,
		cmd.Value("playlist-id").(string),
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
}
