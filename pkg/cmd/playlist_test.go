// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestPlaylistsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "playlists", "retrieve",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
			"--additional-types", "additional_types",
			"--fields", "items(added_by.id,track(name,href,album(name,href)))",
			"--market", "ES",
		)
	})
}

func TestPlaylistsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "playlists", "update",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
			"--collaborative=true",
			"--description", "Updated playlist description",
			"--name", "Updated Playlist Name",
			"--published=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"collaborative: true\n" +
			"description: Updated playlist description\n" +
			"name: Updated Playlist Name\n" +
			"published: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "playlists", "update",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		)
	})
}
