// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestBrowseCategoriesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"browse:categories", "retrieve",
		"--access-token", "string",
		"--category-id", "dinner",
		"--locale", "sv_SE",
	)
}

func TestBrowseCategoriesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"browse:categories", "list",
		"--access-token", "string",
		"--limit", "10",
		"--locale", "sv_SE",
		"--offset", "5",
	)
}

func TestBrowseCategoriesGetPlaylists(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"browse:categories", "get-playlists",
		"--access-token", "string",
		"--category-id", "dinner",
		"--limit", "10",
		"--offset", "5",
	)
}
