// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestRecommendationsGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recommendations", "get",
		"--limit", "10",
		"--market", "ES",
		"--max-acousticness", "0",
		"--max-danceability", "0",
		"--max-duration-ms", "0",
		"--max-energy", "0",
		"--max-instrumentalness", "0",
		"--max-key", "0",
		"--max-liveness", "0",
		"--max-loudness", "0",
		"--max-mode", "0",
		"--max-popularity", "0",
		"--max-speechiness", "0",
		"--max-tempo", "0",
		"--max-time-signature", "0",
		"--max-valence", "0",
		"--min-acousticness", "0",
		"--min-danceability", "0",
		"--min-duration-ms", "0",
		"--min-energy", "0",
		"--min-instrumentalness", "0",
		"--min-key", "0",
		"--min-liveness", "0",
		"--min-loudness", "0",
		"--min-mode", "0",
		"--min-popularity", "0",
		"--min-speechiness", "0",
		"--min-tempo", "0",
		"--min-time-signature", "11",
		"--min-valence", "0",
		"--seed-artists", "4NHQUGzhtTLFvgF5SZesLK",
		"--seed-genres", "classical,country",
		"--seed-tracks", "0c6xIDDpzE81m2q797ordA",
		"--target-acousticness", "0",
		"--target-danceability", "0",
		"--target-duration-ms", "0",
		"--target-energy", "0",
		"--target-instrumentalness", "0",
		"--target-key", "0",
		"--target-liveness", "0",
		"--target-loudness", "0",
		"--target-mode", "0",
		"--target-popularity", "0",
		"--target-speechiness", "0",
		"--target-tempo", "0",
		"--target-time-signature", "0",
		"--target-valence", "0",
	)
}

func TestRecommendationsListAvailableGenreSeeds(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recommendations", "list-available-genre-seeds",
	)
}
