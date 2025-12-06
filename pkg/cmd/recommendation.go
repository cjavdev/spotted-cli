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

var recommendationsGet = cli.Command{
	Name:  "get",
	Usage: "Recommendations are generated based on the available information for a given\nseed entity and matched against similar artists and tracks. If there is\nsufficient information about the provided seeds, a list of tracks will be\nreturned together with pool size details.",
	Flags: []cli.Flag{
		&requestflag.IntFlag{
			Name:  "limit",
			Usage: "The target size of the list of recommended tracks. For seeds with unusually small pools or when highly restrictive filtering is applied, it may be impossible to generate the requested number of recommended tracks. Debugging information for such cases is available in the response. Default: 20\\. Minimum: 1\\. Maximum: 100.\n",
			Value: requestflag.Value[int64](20),
			Config: requestflag.RequestConfig{
				QueryPath: "limit",
			},
		},
		&requestflag.StringFlag{
			Name:  "market",
			Usage: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
			Config: requestflag.RequestConfig{
				QueryPath: "market",
			},
		},
		&requestflag.FloatFlag{
			Name:  "max-acousticness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_acousticness",
			},
		},
		&requestflag.FloatFlag{
			Name:  "max-danceability",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_danceability",
			},
		},
		&requestflag.IntFlag{
			Name:  "max-duration-ms",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_duration_ms",
			},
		},
		&requestflag.FloatFlag{
			Name:  "max-energy",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_energy",
			},
		},
		&requestflag.FloatFlag{
			Name:  "max-instrumentalness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_instrumentalness",
			},
		},
		&requestflag.IntFlag{
			Name:  "max-key",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_key",
			},
		},
		&requestflag.FloatFlag{
			Name:  "max-liveness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_liveness",
			},
		},
		&requestflag.FloatFlag{
			Name:  "max-loudness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_loudness",
			},
		},
		&requestflag.IntFlag{
			Name:  "max-mode",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_mode",
			},
		},
		&requestflag.IntFlag{
			Name:  "max-popularity",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_popularity",
			},
		},
		&requestflag.FloatFlag{
			Name:  "max-speechiness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_speechiness",
			},
		},
		&requestflag.FloatFlag{
			Name:  "max-tempo",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_tempo",
			},
		},
		&requestflag.IntFlag{
			Name:  "max-time-signature",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_time_signature",
			},
		},
		&requestflag.FloatFlag{
			Name:  "max-valence",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "max_valence",
			},
		},
		&requestflag.FloatFlag{
			Name:  "min-acousticness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_acousticness",
			},
		},
		&requestflag.FloatFlag{
			Name:  "min-danceability",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_danceability",
			},
		},
		&requestflag.IntFlag{
			Name:  "min-duration-ms",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_duration_ms",
			},
		},
		&requestflag.FloatFlag{
			Name:  "min-energy",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_energy",
			},
		},
		&requestflag.FloatFlag{
			Name:  "min-instrumentalness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_instrumentalness",
			},
		},
		&requestflag.IntFlag{
			Name:  "min-key",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_key",
			},
		},
		&requestflag.FloatFlag{
			Name:  "min-liveness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_liveness",
			},
		},
		&requestflag.FloatFlag{
			Name:  "min-loudness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_loudness",
			},
		},
		&requestflag.IntFlag{
			Name:  "min-mode",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_mode",
			},
		},
		&requestflag.IntFlag{
			Name:  "min-popularity",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_popularity",
			},
		},
		&requestflag.FloatFlag{
			Name:  "min-speechiness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_speechiness",
			},
		},
		&requestflag.FloatFlag{
			Name:  "min-tempo",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_tempo",
			},
		},
		&requestflag.IntFlag{
			Name:  "min-time-signature",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_time_signature",
			},
		},
		&requestflag.FloatFlag{
			Name:  "min-valence",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "min_valence",
			},
		},
		&requestflag.StringFlag{
			Name:  "seed-artists",
			Usage: "A comma separated list of [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for seed artists.  Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_genres` and `seed_tracks` are not set_.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "seed_artists",
			},
		},
		&requestflag.StringFlag{
			Name:  "seed-genres",
			Usage: "A comma separated list of any genres in the set of [available genre seeds](/documentation/web-api/reference/get-recommendation-genres). Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_artists` and `seed_tracks` are not set_.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "seed_genres",
			},
		},
		&requestflag.StringFlag{
			Name:  "seed-tracks",
			Usage: "A comma separated list of [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for a seed track.  Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_artists` and `seed_genres` are not set_.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "seed_tracks",
			},
		},
		&requestflag.FloatFlag{
			Name:  "target-acousticness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "target_acousticness",
			},
		},
		&requestflag.FloatFlag{
			Name:  "target-danceability",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "target_danceability",
			},
		},
		&requestflag.IntFlag{
			Name:  "target-duration-ms",
			Usage: "Target duration of the track (ms)",
			Config: requestflag.RequestConfig{
				QueryPath: "target_duration_ms",
			},
		},
		&requestflag.FloatFlag{
			Name:  "target-energy",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "target_energy",
			},
		},
		&requestflag.FloatFlag{
			Name:  "target-instrumentalness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "target_instrumentalness",
			},
		},
		&requestflag.IntFlag{
			Name:  "target-key",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "target_key",
			},
		},
		&requestflag.FloatFlag{
			Name:  "target-liveness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "target_liveness",
			},
		},
		&requestflag.FloatFlag{
			Name:  "target-loudness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "target_loudness",
			},
		},
		&requestflag.IntFlag{
			Name:  "target-mode",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "target_mode",
			},
		},
		&requestflag.IntFlag{
			Name:  "target-popularity",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "target_popularity",
			},
		},
		&requestflag.FloatFlag{
			Name:  "target-speechiness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "target_speechiness",
			},
		},
		&requestflag.FloatFlag{
			Name:  "target-tempo",
			Usage: "Target tempo (BPM)",
			Config: requestflag.RequestConfig{
				QueryPath: "target_tempo",
			},
		},
		&requestflag.IntFlag{
			Name:  "target-time-signature",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "target_time_signature",
			},
		},
		&requestflag.FloatFlag{
			Name:  "target-valence",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			Config: requestflag.RequestConfig{
				QueryPath: "target_valence",
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
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.RecommendationGetParams{}

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
	_, err = client.Recommendations.Get(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "recommendations get", obj, format, transform)
}

func handleRecommendationsListAvailableGenreSeeds(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Recommendations.ListAvailableGenreSeeds(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "recommendations list-available-genre-seeds", obj, format, transform)
}
