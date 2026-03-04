// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestMeAlbumsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:albums", "list",
		"--access-token", "string",
		"--limit", "10",
		"--market", "ES",
		"--offset", "5",
	)
}

func TestMeAlbumsCheck(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:albums", "check",
		"--access-token", "string",
		"--ids", "382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc",
	)
}

func TestMeAlbumsRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:albums", "remove",
		"--access-token", "string",
		"--id", "string",
		"--published=true",
	)
}

func TestMeAlbumsSave(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:albums", "save",
		"--access-token", "string",
		"--id", "string",
		"--published=true",
	)
}
