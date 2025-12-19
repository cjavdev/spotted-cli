// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestSearchQuery(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"search", "query",
		"--q", "remaster%20track:Doxy%20artist:Miles%20Davis",
		"--type", "album",
		"--include-external", "audio",
		"--limit", "10",
		"--market", "ES",
		"--offset", "5",
	)
}
