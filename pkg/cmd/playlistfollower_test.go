// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestPlaylistsFollowersCheck(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"playlists:followers", "check",
		"--playlist-id", "3cEYpjA9oz9GiPac4AsH4n",
		"--ids", "jmperezperez",
	)
}
