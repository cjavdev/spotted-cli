// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestBrowseGetFeaturedPlaylists(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"browse", "get-featured-playlists",
		"--limit", "10",
		"--locale", "sv_SE",
		"--offset", "5",
	)
}

func TestBrowseGetNewReleases(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"browse", "get-new-releases",
		"--limit", "10",
		"--offset", "5",
	)
}
