// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestArtistsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"artists", "retrieve",
		"--access-token", "string",
		"--id", "0TnOYISbd1XYRBk9myaseg",
	)
}

func TestArtistsBulkRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"artists", "bulk-retrieve",
		"--access-token", "string",
		"--ids", "2CIMQHirSU0MQqyYHq0eOx,57dN52uHvrHOxijzpIgu3E,1vCWHaC5f2uS3yhpwWbIA6",
	)
}

func TestArtistsListAlbums(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"artists", "list-albums",
		"--access-token", "string",
		"--id", "0TnOYISbd1XYRBk9myaseg",
		"--include-groups", "single,appears_on",
		"--limit", "5",
		"--market", "ES",
		"--offset", "5",
	)
}

func TestArtistsListRelatedArtists(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"artists", "list-related-artists",
		"--access-token", "string",
		"--id", "0TnOYISbd1XYRBk9myaseg",
	)
}

func TestArtistsTopTracks(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"artists", "top-tracks",
		"--access-token", "string",
		"--id", "0TnOYISbd1XYRBk9myaseg",
		"--market", "ES",
	)
}
