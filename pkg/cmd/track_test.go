// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestTracksRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"tracks", "retrieve",
		"--access-token", "string",
		"--id", "11dFghVXANMlKmJXsNCbNl",
		"--market", "ES",
	)
}

func TestTracksBulkRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"tracks", "bulk-retrieve",
		"--access-token", "string",
		"--ids", "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B",
		"--market", "ES",
	)
}
