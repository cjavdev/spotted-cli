// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/spotted-cli/pkg/jsonflag"
	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var recommendationsGet = cli.Command{
	Name:  "get",
	Usage: "Recommendations are generated based on the available information for a given\nseed entity and matched against similar artists and tracks. If there is\nsufficient information about the provided seeds, a list of tracks will be\nreturned together with pool size details.",
	Flags: []cli.Flag{
		&jsonflag.JSONIntFlag{
			Name:  "limit",
			Usage: "The target size of the list of recommended tracks. For seeds with unusually small pools or when highly restrictive filtering is applied, it may be impossible to generate the requested number of recommended tracks. Debugging information for such cases is available in the response. Default: 20\\. Minimum: 1\\. Maximum: 100.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "limit",
			},
			Value: 20,
		},
		&jsonflag.JSONStringFlag{
			Name:  "market",
			Usage: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "market",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "max-acousticness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_acousticness",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "max-danceability",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_danceability",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "max-duration-ms",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_duration_ms",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "max-energy",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_energy",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "max-instrumentalness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_instrumentalness",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "max-key",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_key",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "max-liveness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_liveness",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "max-loudness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_loudness",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "max-mode",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_mode",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "max-popularity",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_popularity",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "max-speechiness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_speechiness",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "max-tempo",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_tempo",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "max-time-signature",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_time_signature",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "max-valence",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "max_valence",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "min-acousticness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_acousticness",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "min-danceability",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_danceability",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "min-duration-ms",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_duration_ms",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "min-energy",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_energy",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "min-instrumentalness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_instrumentalness",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "min-key",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_key",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "min-liveness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_liveness",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "min-loudness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_loudness",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "min-mode",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_mode",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "min-popularity",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_popularity",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "min-speechiness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_speechiness",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "min-tempo",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_tempo",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "min-time-signature",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_time_signature",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "min-valence",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "min_valence",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "seed-artists",
			Usage: "A comma separated list of [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for seed artists.  Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_genres` and `seed_tracks` are not set_.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "seed_artists",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "seed-genres",
			Usage: "A comma separated list of any genres in the set of [available genre seeds](/documentation/web-api/reference/get-recommendation-genres). Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_artists` and `seed_tracks` are not set_.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "seed_genres",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "seed-tracks",
			Usage: "A comma separated list of [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for a seed track.  Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_artists` and `seed_genres` are not set_.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "seed_tracks",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "target-acousticness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_acousticness",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "target-danceability",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_danceability",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "target-duration-ms",
			Usage: "Target duration of the track (ms)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_duration_ms",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "target-energy",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_energy",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "target-instrumentalness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_instrumentalness",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "target-key",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_key",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "target-liveness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_liveness",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "target-loudness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_loudness",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "target-mode",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_mode",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "target-popularity",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_popularity",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "target-speechiness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_speechiness",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "target-tempo",
			Usage: "Target tempo (BPM)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_tempo",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "target-time-signature",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_time_signature",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name:  "target-valence",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "target_valence",
			},
		},
	},
	Action:          handleRecommendationsGet,
	HideHelpCommand: true,
}

var recommendationsListAvailableGenreSeeds = cli.Command{
	Name:            "list-available-genre-seeds",
	Usage:           "Retrieve a list of available genres seed parameter values for\n[recommendations](/documentation/web-api/reference/get-recommendations).",
	Flags:           []cli.Flag{},
	Action:          handleRecommendationsListAvailableGenreSeeds,
	HideHelpCommand: true,
}

func handleRecommendationsGet(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.RecommendationGetParams{}
	var res []byte
	_, err := cc.client.Recommendations.Get(
		ctx,
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("recommendations get", json, format, transform)
}

func handleRecommendationsListAvailableGenreSeeds(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.Recommendations.ListAvailableGenreSeeds(
		ctx,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("recommendations list-available-genre-seeds", json, format, transform)
}
