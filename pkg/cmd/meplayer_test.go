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

func TestMePlayerPausePlayback(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "pause-playback",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerSeekToPosition(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "seek-to-position",
		"--position-ms", "25000",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerSetRepeatMode(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "set-repeat-mode",
		"--state", "context",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerSetVolume(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "set-volume",
		"--volume-percent", "50",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerSkipNext(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "skip-next",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerSkipPrevious(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "skip-previous",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerStartPlayback(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "start-playback",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
		"--context-uri", "spotify:album:5ht7ItJgpBH7W6vJ5BqpPr",
		"--offset", "{position: bar}",
		"--position-ms", "0",
		"--published",
		"--uris", "string",
	)
}

func TestMePlayerToggleShuffle(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "toggle-shuffle",
		"--state",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerTransfer(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player", "transfer",
		"--device-id", "74ASZWbe4lXaubB36ztrGX",
		"--play",
		"--published",
	)
}
