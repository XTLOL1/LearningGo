package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("adding 1+1", func(t *testing.T){
		got := add(1, 1)
		want := 2
		assertCorrectBehaviour(t, got, want)
	})
	t.Run("", func(t *testing.T){
		got := add(-1, 1)
		want := 0
		assertCorrectBehaviour(t, got, want)
	})
}

func assertCorrectBehaviour(t testing.TB, got, want int){
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
