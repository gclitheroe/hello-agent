package main

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "hello world"
	if got != want {
		t.Fatalf("hello() = %q, want %q", got, want)
	}
}
