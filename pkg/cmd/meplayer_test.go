// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestMePlayerGetCurrentlyPlaying(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "get-currently-playing",
		"--additional-types", "additional_types",
		"--market", "ES",
	)
}

func TestMePlayerGetDevices(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "get-devices",
	)
}

func TestMePlayerGetState(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "get-state",
		"--additional-types", "additional_types",
		"--market", "ES",
	)
}

func TestMePlayerListRecentlyPlayed(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "list-recently-played",
		"--after", "1484811043508",
		"--before", "0",
		"--limit", "10",
	)
}
