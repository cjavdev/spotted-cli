// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestMeTracksList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:tracks", "list",
		"--limit", "10",
		"--market", "ES",
		"--offset", "5",
	)
}

func TestMeTracksCheck(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:tracks", "check",
		"--ids", "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B",
	)
}

func TestMeTracksRemove(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:tracks", "remove",
		"--id", "string",
		"--published",
	)
}

func TestMeTracksSave(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:tracks", "save",
		"--id", "string",
		"--published",
		"--timestamped-id", "{id: id, added_at: '2019-12-27T18:11:19.117Z'}\n",
	)
}
