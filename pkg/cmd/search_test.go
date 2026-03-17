// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestSearchQuery(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--access-token", "string",
			"search", "query",
			"--q", "remaster%20track:Doxy%20artist:Miles%20Davis",
			"--type", "album",
			"--include-external", "audio",
			"--limit", "10",
			"--market", "ES",
			"--offset", "5",
		)
	})
}
