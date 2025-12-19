// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestPlaylistsTracksUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"playlists:tracks", "update",
		"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		"--insert-before", "3",
		"--published",
		"--range-length", "2",
		"--range-start", "1",
		"--snapshot-id", "snapshot_id",
		"--uris", "string",
	)
}

func TestPlaylistsTracksList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"playlists:tracks", "list",
		"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		"--additional-types", "additional_types",
		"--fields", "items(added_by.id,track(name,href,album(name,href)))",
		"--limit", "10",
		"--market", "ES",
		"--offset", "5",
	)
}

func TestPlaylistsTracksAdd(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"playlists:tracks", "add",
		"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		"--position", "0",
		"--published",
		"--uris", "string",
	)
}

func TestPlaylistsTracksRemove(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"playlists:tracks", "remove",
		"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		"--track", "{uri: uri}\n",
		"--published",
		"--snapshot-id", "snapshot_id",
	)
}
