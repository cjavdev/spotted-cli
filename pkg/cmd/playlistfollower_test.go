// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestPlaylistsFollowersCheck(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "playlists:followers", "check",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
			"--ids", "jmperezperez",
		)
	})
}

func TestPlaylistsFollowersFollow(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "playlists:followers", "follow",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
			"--published=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("published: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "playlists:followers", "follow",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		)
	})
}

func TestPlaylistsFollowersUnfollow(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "playlists:followers", "unfollow",
			"--access-token", "string",
			"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		)
	})
}
