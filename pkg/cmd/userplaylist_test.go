// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestUsersPlaylistsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:playlists", "create",
		"--access-token", "string",
		"--user-id", "smedjan",
		"--name", "New Playlist",
		"--collaborative=true",
		"--description", "New playlist description",
		"--published=true",
	)
}

func TestUsersPlaylistsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:playlists", "list",
		"--access-token", "string",
		"--user-id", "smedjan",
		"--limit", "10",
		"--offset", "5",
	)
}
