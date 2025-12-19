// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestAudiobooksRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"audiobooks", "retrieve",
		"--id", "7iHfbu1YPACw6oZPAFJtqe",
		"--market", "ES",
	)
}

func TestAudiobooksBulkRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"audiobooks", "bulk-retrieve",
		"--ids", "18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ,7iHfbu1YPACw6oZPAFJtqe",
		"--market", "ES",
	)
}

func TestAudiobooksListChapters(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"audiobooks", "list-chapters",
		"--id", "7iHfbu1YPACw6oZPAFJtqe",
		"--limit", "10",
		"--market", "ES",
		"--offset", "5",
	)
}
