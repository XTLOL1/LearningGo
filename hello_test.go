package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Someone")
	want := "Hello, Someone"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
