package main

import (
	"bytes"
	"testing"
)

func TestCLIVersion(t *testing.T) {
	oldVersion, oldCommit, oldDate := version, commit, date
	t.Cleanup(func() { version, commit, date = oldVersion, oldCommit, oldDate })
	version, commit, date = "2.2.2", "0123456789abcdef0123456789abcdef01234567", "2026-11-03T03:00:00Z"
	var output bytes.Buffer
	if err := run([]string{"--version"}, &output); err != nil {
		t.Fatal(err)
	}
	want := "training-go-releaser version 2.2.2 (commit: 0123456789abcdef0123456789abcdef01234567, date: 2026-11-03T03:00:00Z)\n"
	if output.String() != want {
		t.Fatalf("got %q, want %q", output.String(), want)
	}
}

func TestDefaultOutput(t *testing.T) {
	var output bytes.Buffer
	if err := run(nil, &output); err != nil {
		t.Fatal(err)
	}
	if output.String() != "test 1\n" {
		t.Fatalf("unexpected output: %q", output.String())
	}
}
