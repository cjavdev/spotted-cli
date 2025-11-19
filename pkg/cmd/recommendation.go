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

var recommendationsGet = cli.Command{
	Name:  "get",
	Usage: "Recommendations are generated based on the available information for a given\nseed entity and matched against similar artists and tracks. If there is\nsufficient information about the provided seeds, a list of tracks will be\nreturned together with pool size details.",
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:  "limit",
			Usage: "The target size of the list of recommended tracks. For seeds with unusually small pools or when highly restrictive filtering is applied, it may be impossible to generate the requested number of recommended tracks. Debugging information for such cases is available in the response. Default: 20\\. Minimum: 1\\. Maximum: 100.\n",
			Value: 20,
		},
		&cli.StringFlag{
			Name:  "market",
			Usage: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
		},
		&cli.Float64Flag{
			Name:  "max-acousticness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Float64Flag{
			Name:  "max-danceability",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Int64Flag{
			Name:  "max-duration-ms",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Float64Flag{
			Name:  "max-energy",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Float64Flag{
			Name:  "max-instrumentalness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Int64Flag{
			Name:  "max-key",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Float64Flag{
			Name:  "max-liveness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Float64Flag{
			Name:  "max-loudness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Int64Flag{
			Name:  "max-mode",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Int64Flag{
			Name:  "max-popularity",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Float64Flag{
			Name:  "max-speechiness",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Float64Flag{
			Name:  "max-tempo",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Int64Flag{
			Name:  "max-time-signature",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Float64Flag{
			Name:  "max-valence",
			Usage: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.\n",
		},
		&cli.Float64Flag{
			Name:  "min-acousticness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.Float64Flag{
			Name:  "min-danceability",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.Int64Flag{
			Name:  "min-duration-ms",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.Float64Flag{
			Name:  "min-energy",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.Float64Flag{
			Name:  "min-instrumentalness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.Int64Flag{
			Name:  "min-key",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.Float64Flag{
			Name:  "min-liveness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.Float64Flag{
			Name:  "min-loudness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.Int64Flag{
			Name:  "min-mode",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.Int64Flag{
			Name:  "min-popularity",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.Float64Flag{
			Name:  "min-speechiness",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.Float64Flag{
			Name:  "min-tempo",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.Int64Flag{
			Name:  "min-time-signature",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.Float64Flag{
			Name:  "min-valence",
			Usage: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.\n",
		},
		&cli.StringFlag{
			Name:  "seed-artists",
			Usage: "A comma separated list of [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for seed artists.  Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_genres` and `seed_tracks` are not set_.\n",
		},
		&cli.StringFlag{
			Name:  "seed-genres",
			Usage: "A comma separated list of any genres in the set of [available genre seeds](/documentation/web-api/reference/get-recommendation-genres). Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_artists` and `seed_tracks` are not set_.\n",
		},
		&cli.StringFlag{
			Name:  "seed-tracks",
			Usage: "A comma separated list of [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for a seed track.  Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_artists` and `seed_genres` are not set_.\n",
		},
		&cli.Float64Flag{
			Name:  "target-acousticness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
		},
		&cli.Float64Flag{
			Name:  "target-danceability",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
		},
		&cli.Int64Flag{
			Name:  "target-duration-ms",
			Usage: "Target duration of the track (ms)",
		},
		&cli.Float64Flag{
			Name:  "target-energy",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
		},
		&cli.Float64Flag{
			Name:  "target-instrumentalness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
		},
		&cli.Int64Flag{
			Name:  "target-key",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
		},
		&cli.Float64Flag{
			Name:  "target-liveness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
		},
		&cli.Float64Flag{
			Name:  "target-loudness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
		},
		&cli.Int64Flag{
			Name:  "target-mode",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
		},
		&cli.Int64Flag{
			Name:  "target-popularity",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
		},
		&cli.Float64Flag{
			Name:  "target-speechiness",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
		},
		&cli.Float64Flag{
			Name:  "target-tempo",
			Usage: "Target tempo (BPM)",
		},
		&cli.Int64Flag{
			Name:  "target-time-signature",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
		},
		&cli.Float64Flag{
			Name:  "target-valence",
			Usage: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.\n",
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
	params := spotted.RecommendationGetParams{
		Market:                 spotted.String(cmd.Value("market").(string)),
		MaxAcousticness:        spotted.Float(cmd.Value("max-acousticness").(float64)),
		MaxDanceability:        spotted.Float(cmd.Value("max-danceability").(float64)),
		MaxDurationMs:          spotted.Int(cmd.Value("max-duration-ms").(int64)),
		MaxEnergy:              spotted.Float(cmd.Value("max-energy").(float64)),
		MaxInstrumentalness:    spotted.Float(cmd.Value("max-instrumentalness").(float64)),
		MaxKey:                 spotted.Int(cmd.Value("max-key").(int64)),
		MaxLiveness:            spotted.Float(cmd.Value("max-liveness").(float64)),
		MaxLoudness:            spotted.Float(cmd.Value("max-loudness").(float64)),
		MaxMode:                spotted.Int(cmd.Value("max-mode").(int64)),
		MaxPopularity:          spotted.Int(cmd.Value("max-popularity").(int64)),
		MaxSpeechiness:         spotted.Float(cmd.Value("max-speechiness").(float64)),
		MaxTempo:               spotted.Float(cmd.Value("max-tempo").(float64)),
		MaxTimeSignature:       spotted.Int(cmd.Value("max-time-signature").(int64)),
		MaxValence:             spotted.Float(cmd.Value("max-valence").(float64)),
		MinAcousticness:        spotted.Float(cmd.Value("min-acousticness").(float64)),
		MinDanceability:        spotted.Float(cmd.Value("min-danceability").(float64)),
		MinDurationMs:          spotted.Int(cmd.Value("min-duration-ms").(int64)),
		MinEnergy:              spotted.Float(cmd.Value("min-energy").(float64)),
		MinInstrumentalness:    spotted.Float(cmd.Value("min-instrumentalness").(float64)),
		MinKey:                 spotted.Int(cmd.Value("min-key").(int64)),
		MinLiveness:            spotted.Float(cmd.Value("min-liveness").(float64)),
		MinLoudness:            spotted.Float(cmd.Value("min-loudness").(float64)),
		MinMode:                spotted.Int(cmd.Value("min-mode").(int64)),
		MinPopularity:          spotted.Int(cmd.Value("min-popularity").(int64)),
		MinSpeechiness:         spotted.Float(cmd.Value("min-speechiness").(float64)),
		MinTempo:               spotted.Float(cmd.Value("min-tempo").(float64)),
		MinTimeSignature:       spotted.Int(cmd.Value("min-time-signature").(int64)),
		MinValence:             spotted.Float(cmd.Value("min-valence").(float64)),
		SeedArtists:            spotted.String(cmd.Value("seed-artists").(string)),
		SeedGenres:             spotted.String(cmd.Value("seed-genres").(string)),
		SeedTracks:             spotted.String(cmd.Value("seed-tracks").(string)),
		TargetAcousticness:     spotted.Float(cmd.Value("target-acousticness").(float64)),
		TargetDanceability:     spotted.Float(cmd.Value("target-danceability").(float64)),
		TargetDurationMs:       spotted.Int(cmd.Value("target-duration-ms").(int64)),
		TargetEnergy:           spotted.Float(cmd.Value("target-energy").(float64)),
		TargetInstrumentalness: spotted.Float(cmd.Value("target-instrumentalness").(float64)),
		TargetKey:              spotted.Int(cmd.Value("target-key").(int64)),
		TargetLiveness:         spotted.Float(cmd.Value("target-liveness").(float64)),
		TargetLoudness:         spotted.Float(cmd.Value("target-loudness").(float64)),
		TargetMode:             spotted.Int(cmd.Value("target-mode").(int64)),
		TargetPopularity:       spotted.Int(cmd.Value("target-popularity").(int64)),
		TargetSpeechiness:      spotted.Float(cmd.Value("target-speechiness").(float64)),
		TargetTempo:            spotted.Float(cmd.Value("target-tempo").(float64)),
		TargetTimeSignature:    spotted.Int(cmd.Value("target-time-signature").(int64)),
		TargetValence:          spotted.Float(cmd.Value("target-valence").(float64)),
	}
	if cmd.IsSet("limit") {
		params.Limit = spotted.Opt(cmd.Value("limit").(int64))
	}
	var res []byte
	_, err := client.Recommendations.Get(
		ctx,
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
	return ShowJSON("recommendations get", json, format, transform)
}

func handleRecommendationsListAvailableGenreSeeds(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := client.Recommendations.ListAvailableGenreSeeds(
		ctx,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
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
