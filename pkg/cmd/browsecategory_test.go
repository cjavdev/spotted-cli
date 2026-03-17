// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestBrowseCategoriesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--access-token", "string",
			"browse:categories", "retrieve",
			"--category-id", "dinner",
			"--locale", "sv_SE",
		)
	})
}

func TestBrowseCategoriesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--access-token", "string",
			"browse:categories", "list",
			"--max-items", "10",
			"--limit", "10",
			"--locale", "sv_SE",
			"--offset", "5",
		)
	})
}

func TestBrowseCategoriesGetPlaylists(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--access-token", "string",
			"browse:categories", "get-playlists",
			"--category-id", "dinner",
			"--limit", "10",
			"--offset", "5",
		)
	})
}
