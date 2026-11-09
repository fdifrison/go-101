package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("say hello", func(t *testing.T) {
		got := Hello("Giovanni")
		want := "Hello, Giovanni"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("say 'Hello World' when no name is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
