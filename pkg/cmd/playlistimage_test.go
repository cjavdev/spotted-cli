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
			t,
			"--access-token", "string",
			"playlists:images", "update",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
			"--body", mocktest.TestFile(t, "Example data"),
			"--output", "/dev/null",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("Example data")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--access-token", "string",
			"playlists:images", "update",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
			"--output", "/dev/null",
		)
	})
}

func TestPlaylistsImagesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--access-token", "string",
			"playlists:images", "list",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		)
	})
}
