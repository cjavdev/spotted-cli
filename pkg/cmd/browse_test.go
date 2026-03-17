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
			t,
			"--access-token", "string",
			"browse", "get-featured-playlists",
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
			t,
			"--access-token", "string",
			"browse", "get-new-releases",
			"--limit", "10",
			"--offset", "5",
		)
	})
}
