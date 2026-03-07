// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestAlbumsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "albums", "retrieve",
			"--access-token", "string",
			"--id", "4aawyAB9vmqN3uQ7FjRGTy",
			"--market", "ES",
		)
	})
}

func TestAlbumsBulkRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "albums", "bulk-retrieve",
			"--access-token", "string",
			"--ids", "382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc",
			"--market", "ES",
		)
	})
}

func TestAlbumsListTracks(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "albums", "list-tracks",
			"--access-token", "string",
			"--max-items", "10",
			"--id", "4aawyAB9vmqN3uQ7FjRGTy",
			"--limit", "10",
			"--market", "ES",
			"--offset", "5",
		)
	})
}
