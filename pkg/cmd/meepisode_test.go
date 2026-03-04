// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestMeEpisodesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:episodes", "list",
		"--access-token", "string",
		"--limit", "10",
		"--market", "ES",
		"--offset", "5",
	)
}

func TestMeEpisodesCheck(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:episodes", "check",
		"--access-token", "string",
		"--ids", "77o6BIVlYM3msb4MMIL1jH,0Q86acNRm6V9GYx55SXKwf",
	)
}

func TestMeEpisodesRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:episodes", "remove",
		"--access-token", "string",
		"--id", "string",
		"--published=true",
	)
}

func TestMeEpisodesSave(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:episodes", "save",
		"--access-token", "string",
		"--id", "string",
		"--published=true",
	)
}
