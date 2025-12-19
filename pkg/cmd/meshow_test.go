// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestMeShowsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:shows", "list",
		"--limit", "10",
		"--offset", "5",
	)
}

func TestMeShowsCheck(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:shows", "check",
		"--ids", "5CfCWKI5pZ28U0uOzXkDHe,5as3aKmN2k11yfDDDSrvaZ",
	)
}

func TestMeShowsRemove(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:shows", "remove",
		"--id", "string",
		"--published",
	)
}

func TestMeShowsSave(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:shows", "save",
		"--id", "string",
		"--published",
	)
}
