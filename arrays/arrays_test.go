package main

import "testing"

func TestArrays(t *testing.T) {
	t.Run("summing array of numbers", func(t *testing.T){
		array := []int{1, 2, 3, 4, 5}
		got := sumNumbers(array)
		want := 15
		assertEquals(t, got, want, array)
	})
	//t.Run("", func(t *testing.T){
	//	got := add(-1, 1)
	//	want := 0
	//	assertCorrectBehaviour(t, got, want)
	//})
}

func assertEquals[T comparable](t testing.TB, got, want T, array []T){
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v, list: %v", got, want, array)
	}
}
