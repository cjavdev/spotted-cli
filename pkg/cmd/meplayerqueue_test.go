// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestMePlayerQueueAdd(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player:queue", "add",
		"--uri", "spotify:track:4iV5W9uYEdYUVa79Axb7Rh",
		"--device-id", "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	)
}

func TestMePlayerQueueGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:player:queue", "get",
	)
}
