// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestPlaylistsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"playlists", "retrieve",
		"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		"--additional-types", "additional_types",
		"--fields", "items(added_by.id,track(name,href,album(name,href)))",
		"--market", "ES",
	)
}

func TestPlaylistsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"playlists", "update",
		"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		"--collaborative",
		"--description", "Updated playlist description",
		"--name", "Updated Playlist Name",
		"--published",
	)
}
