// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestShowsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "shows", "retrieve",
			"--access-token", "string",
			"--id", "38bS44xjbVVZ3No3ByF1dJ",
			"--market", "ES",
		)
	})
}

func TestShowsBulkRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "shows", "bulk-retrieve",
			"--access-token", "string",
			"--ids", "5CfCWKI5pZ28U0uOzXkDHe,5as3aKmN2k11yfDDDSrvaZ",
			"--market", "ES",
		)
	})
}

func TestShowsListEpisodes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "shows", "list-episodes",
			"--access-token", "string",
			"--id", "38bS44xjbVVZ3No3ByF1dJ",
			"--limit", "10",
			"--market", "ES",
			"--offset", "5",
		)
	})
}
