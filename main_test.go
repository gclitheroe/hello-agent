package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	t.Parallel()

	var output bytes.Buffer
	run(&output)

	const want = "Hello Team\n"
	if got := output.String(); got != want {
		t.Errorf("output = %q, want %q", got, want)
	}
}
