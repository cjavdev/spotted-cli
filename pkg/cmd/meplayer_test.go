// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestMePlayerGetCurrentlyPlaying(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "get-currently-playing",
		"--access-token", "string",
		"--additional-types", "additional_types",
		"--market", "ES",
	)
}

func TestMePlayerGetDevices(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "get-devices",
		"--access-token", "string",
	)
}

func TestMePlayerGetState(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "get-state",
		"--access-token", "string",
		"--additional-types", "additional_types",
		"--market", "ES",
	)
}

func TestMePlayerListRecentlyPlayed(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "list-recently-played",
		"--access-token", "string",
		"--after", "1484811043508",
		"--before", "0",
		"--limit", "10",
	)
}

func TestMePlayerPausePlayback(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "pause-playback",
		"--access-token", "string",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerSeekToPosition(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "seek-to-position",
		"--access-token", "string",
		"--position-ms", "25000",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerSetRepeatMode(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "set-repeat-mode",
		"--access-token", "string",
		"--state", "context",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerSetVolume(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "set-volume",
		"--access-token", "string",
		"--volume-percent", "50",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerSkipNext(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "skip-next",
		"--access-token", "string",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerSkipPrevious(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "skip-previous",
		"--access-token", "string",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerStartPlayback(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "start-playback",
		"--access-token", "string",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
		"--context-uri", "spotify:album:5ht7ItJgpBH7W6vJ5BqpPr",
		"--offset", "{position: bar}",
		"--position-ms", "0",
		"--published=true",
		"--uris", "string",
	)
}

func TestMePlayerToggleShuffle(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "toggle-shuffle",
		"--access-token", "string",
		"--state=true",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerTransfer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "transfer",
		"--access-token", "string",
		"--device-id", "74ASZWbe4lXaubB36ztrGX",
		"--play=true",
		"--published=true",
	)
}
