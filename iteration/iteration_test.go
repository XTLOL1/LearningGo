package main

import "testing"

func TestIteration(t *testing.T) {
	t.Run("adding 1+1", func(t *testing.T){
		got := iterateString("a", 5)
		want := "aaaaa"
		assertEquals(t, got, want)
	})
	//t.Run("", func(t *testing.T){
	//	got := add(-1, 1)
	//	want := 0
	//	assertCorrectBehaviour(t, got, want)
	//})
}

func assertEquals[T comparable](t testing.TB, got, want T){
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
