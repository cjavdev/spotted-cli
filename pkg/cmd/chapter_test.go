// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestChaptersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"chapters", "retrieve",
		"--access-token", "string",
		"--id", "0D5wENdkdwbqlrHoaJ9g29",
		"--market", "ES",
	)
}

func TestChaptersBulkRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"chapters", "bulk-retrieve",
		"--access-token", "string",
		"--ids", "0IsXVP0JmcB2adSE338GkK,3ZXb8FKZGU0EHALYX6uCzU,0D5wENdkdwbqlrHoaJ9g29",
		"--market", "ES",
	)
}
