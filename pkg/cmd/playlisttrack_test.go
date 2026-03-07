// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
	"github.com/cjavdev/spotted-cli/internal/requestflag"
)

func TestPlaylistsTracksUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "playlists:tracks", "update",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
			"--insert-before", "3",
			"--published=true",
			"--range-length", "2",
			"--range-start", "1",
			"--snapshot-id", "snapshot_id",
			"--uris", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"insert_before: 3\n" +
			"published: true\n" +
			"range_length: 2\n" +
			"range_start: 1\n" +
			"snapshot_id: snapshot_id\n" +
			"uris:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "playlists:tracks", "update",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		)
	})
}

func TestPlaylistsTracksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "playlists:tracks", "list",
			"--access-token", "string",
			"--max-items", "10",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
			"--additional-types", "additional_types",
			"--fields", "items(added_by.id,track(name,href,album(name,href)))",
			"--limit", "10",
			"--market", "ES",
			"--offset", "5",
		)
	})
}

func TestPlaylistsTracksAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "playlists:tracks", "add",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
			"--position", "0",
			"--published=true",
			"--uris", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"position: 0\n" +
			"published: true\n" +
			"uris:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "playlists:tracks", "add",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		)
	})
}

func TestPlaylistsTracksRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "playlists:tracks", "remove",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
			"--track", "{uri: uri}",
			"--published=true",
			"--snapshot-id", "snapshot_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(playlistsTracksRemove)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "playlists:tracks", "remove",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
			"--track.uri", "uri",
			"--published=true",
			"--snapshot-id", "snapshot_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"tracks:\n" +
			"  - uri: uri\n" +
			"published: true\n" +
			"snapshot_id: snapshot_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "playlists:tracks", "remove",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		)
	})
}
