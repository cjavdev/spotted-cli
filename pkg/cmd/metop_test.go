// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestMeTopListTopArtists(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--access-token", "string",
			"me:top", "list-top-artists",
			"--max-items", "10",
			"--limit", "10",
			"--offset", "5",
			"--time-range", "medium_term",
		)
	})
}

func TestMeTopListTopTracks(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--access-token", "string",
			"me:top", "list-top-tracks",
			"--max-items", "10",
			"--limit", "10",
			"--offset", "5",
			"--time-range", "medium_term",
		)
	})
}
