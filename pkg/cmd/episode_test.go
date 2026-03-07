// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestEpisodesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "episodes", "retrieve",
			"--access-token", "string",
			"--id", "512ojhOuo1ktJprKbVcKyQ",
			"--market", "ES",
		)
	})
}

func TestEpisodesBulkRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "episodes", "bulk-retrieve",
			"--access-token", "string",
			"--ids", "77o6BIVlYM3msb4MMIL1jH,0Q86acNRm6V9GYx55SXKwf",
			"--market", "ES",
		)
	})
}
