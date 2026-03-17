// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestArtistsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--access-token", "string",
			"artists", "retrieve",
			"--id", "0TnOYISbd1XYRBk9myaseg",
		)
	})
}

func TestArtistsBulkRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--access-token", "string",
			"artists", "bulk-retrieve",
			"--ids", "2CIMQHirSU0MQqyYHq0eOx,57dN52uHvrHOxijzpIgu3E,1vCWHaC5f2uS3yhpwWbIA6",
		)
	})
}

func TestArtistsListAlbums(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--access-token", "string",
			"artists", "list-albums",
			"--max-items", "10",
			"--id", "0TnOYISbd1XYRBk9myaseg",
			"--include-groups", "single,appears_on",
			"--limit", "5",
			"--market", "ES",
			"--offset", "5",
		)
	})
}

func TestArtistsListRelatedArtists(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--access-token", "string",
			"artists", "list-related-artists",
			"--id", "0TnOYISbd1XYRBk9myaseg",
		)
	})
}

func TestArtistsTopTracks(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--access-token", "string",
			"artists", "top-tracks",
			"--id", "0TnOYISbd1XYRBk9myaseg",
			"--market", "ES",
		)
	})
}
