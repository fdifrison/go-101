package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Giovanni")
	want := "Hello, Giovanni"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
