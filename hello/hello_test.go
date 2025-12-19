package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Someone", "")
		want := "Hello, Someone"
		assertCorrectBehaviour(t, got, want)
	})
	t.Run("saying hello to everyone, no argument supplied", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectBehaviour(t, got, want)
	})
	t.Run("saying hello to people in Portuguese", func(t *testing.T){
		got := Hello("Someone", "Portuguese")
		want := "Ola, Someone"
		assertCorrectBehaviour(t, got, want)
	})
	t.Run("saying hello to everyone in Portuguese", func(t *testing.T){
		got := Hello("", "Portuguese")
		want := "Ola, World"
		assertCorrectBehaviour(t, got, want)
	})
}

func assertCorrectBehaviour(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
