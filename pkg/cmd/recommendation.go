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
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The target size of the list of recommended tracks. For seeds with unusually small pools or when highly restrictive filtering is applied, it may be impossible to generate the requested number of recommended tracks. Debugging information for such cases is available in the response. Default: 20\\. Minimum: 1\\. Maximum: 100.\n",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "market",
			Usage:     "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
			QueryPath: "market",
		},
		&requestflag.Flag[float64]{
			Name:      "max-acousticness",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_acousticness",
		},
		&requestflag.Flag[float64]{
			Name:      "max-danceability",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_danceability",
		},
		&requestflag.Flag[int64]{
			Name:      "max-duration-ms",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_duration_ms",
		},
		&requestflag.Flag[float64]{
			Name:      "max-energy",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_energy",
		},
		&requestflag.Flag[float64]{
			Name:      "max-instrumentalness",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_instrumentalness",
		},
		&requestflag.Flag[int64]{
			Name:      "max-key",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_key",
		},
		&requestflag.Flag[float64]{
			Name:      "max-liveness",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_liveness",
		},
		&requestflag.Flag[float64]{
			Name:      "max-loudness",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_loudness",
		},
		&requestflag.Flag[int64]{
			Name:      "max-mode",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_mode",
		},
		&requestflag.Flag[int64]{
			Name:      "max-popularity",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_popularity",
		},
		&requestflag.Flag[float64]{
			Name:      "max-speechiness",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_speechiness",
		},
		&requestflag.Flag[float64]{
			Name:      "max-tempo",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_tempo",
		},
		&requestflag.Flag[int64]{
			Name:      "max-time-signature",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_time_signature",
		},
		&requestflag.Flag[float64]{
			Name:      "max-valence",
			Usage:     "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
			QueryPath: "max_valence",
		},
		&requestflag.Flag[float64]{
			Name:      "min-acousticness",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_acousticness",
		},
		&requestflag.Flag[float64]{
			Name:      "min-danceability",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_danceability",
		},
		&requestflag.Flag[int64]{
			Name:      "min-duration-ms",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_duration_ms",
		},
		&requestflag.Flag[float64]{
			Name:      "min-energy",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_energy",
		},
		&requestflag.Flag[float64]{
			Name:      "min-instrumentalness",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_instrumentalness",
		},
		&requestflag.Flag[int64]{
			Name:      "min-key",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_key",
		},
		&requestflag.Flag[float64]{
			Name:      "min-liveness",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_liveness",
		},
		&requestflag.Flag[float64]{
			Name:      "min-loudness",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_loudness",
		},
		&requestflag.Flag[int64]{
			Name:      "min-mode",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_mode",
		},
		&requestflag.Flag[int64]{
			Name:      "min-popularity",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_popularity",
		},
		&requestflag.Flag[float64]{
			Name:      "min-speechiness",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_speechiness",
		},
		&requestflag.Flag[float64]{
			Name:      "min-tempo",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_tempo",
		},
		&requestflag.Flag[int64]{
			Name:      "min-time-signature",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_time_signature",
		},
		&requestflag.Flag[float64]{
			Name:      "min-valence",
			Usage:     "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
			QueryPath: "min_valence",
		},
		&requestflag.Flag[string]{
			Name:      "seed-artists",
			Usage:     "A comma separated list of [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for seed artists.  Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_genres` and `seed_tracks` are not set_.\n",
			QueryPath: "seed_artists",
		},
		&requestflag.Flag[string]{
			Name:      "seed-genres",
			Usage:     "A comma separated list of any genres in the set of [available genre seeds](/documentation/web-api/reference/get-recommendation-genres). Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_artists` and `seed_tracks` are not set_.\n",
			QueryPath: "seed_genres",
		},
		&requestflag.Flag[string]{
			Name:      "seed-tracks",
			Usage:     "A comma separated list of [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for a seed track.  Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_artists` and `seed_genres` are not set_.\n",
			QueryPath: "seed_tracks",
		},
		&requestflag.Flag[float64]{
			Name:      "target-acousticness",
			Usage:     "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			QueryPath: "target_acousticness",
		},
		&requestflag.Flag[float64]{
			Name:      "target-danceability",
			Usage:     "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			QueryPath: "target_danceability",
		},
		&requestflag.Flag[int64]{
			Name:      "target-duration-ms",
			Usage:     "Target duration of the track (ms)",
			QueryPath: "target_duration_ms",
		},
		&requestflag.Flag[float64]{
			Name:      "target-energy",
			Usage:     "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			QueryPath: "target_energy",
		},
		&requestflag.Flag[float64]{
			Name:      "target-instrumentalness",
			Usage:     "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			QueryPath: "target_instrumentalness",
		},
		&requestflag.Flag[int64]{
			Name:      "target-key",
			Usage:     "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			QueryPath: "target_key",
		},
		&requestflag.Flag[float64]{
			Name:      "target-liveness",
			Usage:     "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			QueryPath: "target_liveness",
		},
		&requestflag.Flag[float64]{
			Name:      "target-loudness",
			Usage:     "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			QueryPath: "target_loudness",
		},
		&requestflag.Flag[int64]{
			Name:      "target-mode",
			Usage:     "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			QueryPath: "target_mode",
		},
		&requestflag.Flag[int64]{
			Name:      "target-popularity",
			Usage:     "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			QueryPath: "target_popularity",
		},
		&requestflag.Flag[float64]{
			Name:      "target-speechiness",
			Usage:     "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			QueryPath: "target_speechiness",
		},
		&requestflag.Flag[float64]{
			Name:      "target-tempo",
			Usage:     "Target tempo (BPM)",
			QueryPath: "target_tempo",
		},
		&requestflag.Flag[int64]{
			Name:      "target-time-signature",
			Usage:     "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			QueryPath: "target_time_signature",
		},
		&requestflag.Flag[float64]{
			Name:      "target-valence",
			Usage:     "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
			QueryPath: "target_valence",
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
		EmptyBody,
		false,
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
		EmptyBody,
		false,
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
