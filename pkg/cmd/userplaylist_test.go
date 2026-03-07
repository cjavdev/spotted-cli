// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestUsersPlaylistsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:playlists", "create",
			"--access-token", "string",
			"--user-id", "smedjan",
			"--name", "New Playlist",
			"--collaborative=true",
			"--description", "New playlist description",
			"--published=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: New Playlist\n" +
			"collaborative: true\n" +
			"description: New playlist description\n" +
			"published: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "users:playlists", "create",
			"--access-token", "string",
			"--user-id", "smedjan",
		)
	})
}

func TestUsersPlaylistsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:playlists", "list",
			"--access-token", "string",
			"--user-id", "smedjan",
			"--limit", "10",
			"--offset", "5",
		)
	})
}
