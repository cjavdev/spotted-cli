// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/cjavdev/spotted-cli/internal/autocomplete"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command *cli.Command
)

func init() {
	Command = &cli.Command{
		Name:    "spotted",
		Usage:   "CLI for the spotted API",
		Suggest: true,
		Version: Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "albums",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&albumsRetrieve,
					&albumsBulkRetrieve,
					&albumsListTracks,
				},
			},
			{
				Name:     "artists",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&artistsRetrieve,
					&artistsBulkRetrieve,
					&artistsListAlbums,
					&artistsListRelatedArtists,
					&artistsTopTracks,
				},
			},
			{
				Name:     "shows",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&showsRetrieve,
					&showsBulkRetrieve,
					&showsListEpisodes,
				},
			},
			{
				Name:     "episodes",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&episodesRetrieve,
					&episodesBulkRetrieve,
				},
			},
			{
				Name:     "audiobooks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&audiobooksRetrieve,
					&audiobooksBulkRetrieve,
					&audiobooksListChapters,
				},
			},
			{
				Name:     "me",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&meRetrieve,
				},
			},
			{
				Name:     "me:audiobooks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&meAudiobooksList,
					&meAudiobooksCheck,
					&meAudiobooksRemove,
					&meAudiobooksSave,
				},
			},
			{
				Name:     "me:playlists",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mePlaylistsList,
				},
			},
			{
				Name:     "me:top",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&meTopListTopArtists,
					&meTopListTopTracks,
				},
			},
			{
				Name:     "me:albums",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&meAlbumsList,
					&meAlbumsCheck,
					&meAlbumsRemove,
					&meAlbumsSave,
				},
			},
			{
				Name:     "me:tracks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&meTracksList,
					&meTracksCheck,
					&meTracksRemove,
					&meTracksSave,
				},
			},
			{
				Name:     "me:episodes",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&meEpisodesList,
					&meEpisodesCheck,
					&meEpisodesRemove,
					&meEpisodesSave,
				},
			},
			{
				Name:     "me:shows",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&meShowsList,
					&meShowsCheck,
					&meShowsRemove,
					&meShowsSave,
				},
			},
			{
				Name:     "me:following",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&meFollowingBulkRetrieve,
					&meFollowingCheck,
					&meFollowingFollow,
					&meFollowingUnfollow,
				},
			},
			{
				Name:     "me:player",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mePlayerGetCurrentlyPlaying,
					&mePlayerGetDevices,
					&mePlayerGetState,
					&mePlayerListRecentlyPlayed,
					&mePlayerPausePlayback,
					&mePlayerSeekToPosition,
					&mePlayerSetRepeatMode,
					&mePlayerSetVolume,
					&mePlayerSkipNext,
					&mePlayerSkipPrevious,
					&mePlayerStartPlayback,
					&mePlayerToggleShuffle,
					&mePlayerTransfer,
				},
			},
			{
				Name:     "me:player:queue",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mePlayerQueueAdd,
					&mePlayerQueueGet,
				},
			},
			{
				Name:     "chapters",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&chaptersRetrieve,
					&chaptersBulkRetrieve,
				},
			},
			{
				Name:     "tracks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&tracksRetrieve,
					&tracksBulkRetrieve,
				},
			},
			{
				Name:     "search",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&searchQuery,
				},
			},
			{
				Name:     "playlists",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&playlistsRetrieve,
					&playlistsUpdate,
				},
			},
			{
				Name:     "playlists:tracks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&playlistsTracksUpdate,
					&playlistsTracksList,
					&playlistsTracksAdd,
					&playlistsTracksRemove,
				},
			},
			{
				Name:     "playlists:followers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&playlistsFollowersCheck,
					&playlistsFollowersFollow,
					&playlistsFollowersUnfollow,
				},
			},
			{
				Name:     "playlists:images",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&playlistsImagesList,
				},
			},
			{
				Name:     "users",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersRetrieveProfile,
				},
			},
			{
				Name:     "users:playlists",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersPlaylistsCreate,
					&usersPlaylistsList,
				},
			},
			{
				Name:     "browse",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&browseGetFeaturedPlaylists,
					&browseGetNewReleases,
				},
			},
			{
				Name:     "browse:categories",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&browseCategoriesRetrieve,
					&browseCategoriesList,
					&browseCategoriesGetPlaylists,
				},
			},
			{
				Name:     "audio-features",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&audioFeaturesRetrieve,
					&audioFeaturesBulkRetrieve,
				},
			},
			{
				Name:     "audio-analysis",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&audioAnalysisRetrieve,
				},
			},
			{
				Name:     "recommendations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&recommendationsGet,
					&recommendationsListAvailableGenreSeeds,
				},
			},
			{
				Name:     "markets",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&marketsList,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "spotted @manpages [-o spotted.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "spotted.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "spotted.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
