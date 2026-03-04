// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestPlaylistsImagesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"playlists:images", "update",
		"--access-token", "string",
		"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		"--body", mocktest.TestFile(t, "..."),
		"--output", "/dev/null",
	)
}

func TestPlaylistsImagesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"playlists:images", "list",
		"--access-token", "string",
		"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
	)
}
