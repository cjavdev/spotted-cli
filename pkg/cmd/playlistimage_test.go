// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestPlaylistsImagesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "playlists:images", "update",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
			"--body", mocktest.TestFile(t, "..."),
			"--output", "/dev/null",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("...")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "playlists:images", "update",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
			"--output", "/dev/null",
		)
	})
}

func TestPlaylistsImagesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "playlists:images", "list",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		)
	})
}
