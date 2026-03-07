// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/spotted-cli/internal/mocktest"
	"github.com/cjavdev/spotted-cli/internal/requestflag"
)

func TestMeTracksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "me:tracks", "list",
			"--access-token", "string",
			"--limit", "10",
			"--market", "ES",
			"--offset", "5",
		)
	})
}

func TestMeTracksCheck(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "me:tracks", "check",
			"--access-token", "string",
			"--ids", "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B",
		)
	})
}

func TestMeTracksRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "me:tracks", "remove",
			"--access-token", "string",
			"--id", "string",
			"--published=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ids:\n" +
			"  - string\n" +
			"published: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "me:tracks", "remove",
			"--access-token", "string",
		)
	})
}

func TestMeTracksSave(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "me:tracks", "save",
			"--access-token", "string",
			"--id", "string",
			"--published=true",
			"--timestamped-id", "{id: id, added_at: '2019-12-27T18:11:19.117Z'}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(meTracksSave)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "me:tracks", "save",
			"--access-token", "string",
			"--id", "string",
			"--published=true",
			"--timestamped-id.id", "id",
			"--timestamped-id.added-at", "2019-12-27T18:11:19.117Z",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ids:\n" +
			"  - string\n" +
			"published: true\n" +
			"timestamped_ids:\n" +
			"  - id: id\n" +
			"    added_at: '2019-12-27T18:11:19.117Z'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "me:tracks", "save",
			"--access-token", "string",
		)
	})
}
