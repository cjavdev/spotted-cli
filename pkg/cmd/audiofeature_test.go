// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestAudioFeaturesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"audio-features", "retrieve",
		"--id", "11dFghVXANMlKmJXsNCbNl",
	)
}

func TestAudioFeaturesBulkRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"audio-features", "bulk-retrieve",
		"--ids", "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B",
	)
}
