// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestBrowseGetFeaturedPlaylists(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "browse", "get-featured-playlists",
			"--access-token", "string",
			"--limit", "10",
			"--locale", "sv_SE",
			"--offset", "5",
		)
	})
}

func TestBrowseGetNewReleases(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "browse", "get-new-releases",
			"--access-token", "string",
			"--limit", "10",
			"--offset", "5",
		)
	})
}
