// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
)

func TestMeFollowingBulkRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:following", "bulk-retrieve",
		"--type", "artist",
		"--after", "0I2XqVXqHScXjHhk6AYYRe",
		"--limit", "10",
	)
}

func TestMeFollowingCheck(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:following", "check",
		"--ids", "2CIMQHirSU0MQqyYHq0eOx,57dN52uHvrHOxijzpIgu3E,1vCWHaC5f2uS3yhpwWbIA6",
		"--type", "artist",
	)
}

func TestMeFollowingFollow(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:following", "follow",
		"--id", "string",
		"--published",
	)
}

func TestMeFollowingUnfollow(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"me:following", "unfollow",
		"--id", "string",
		"--published",
	)
}
