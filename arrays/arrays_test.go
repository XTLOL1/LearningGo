package main

import "testing"
import "slices"

func TestArrays(t *testing.T) {
	t.Run("summing slice of numbers", func(t *testing.T){
		array := []int{1, 2, 3, 4, 5}
		got := sumNumbers(array)
		want := 15
		assertEquals(t, got, want)
	})
	t.Run("summing slices of numbers", func(t *testing.T){
		got := sumAllNumbers([]int{2, 4, 5, 3}, []int{5, 5})
		want := []int{14, 10}
		assertEquals(t, slices.Equal(got, want), true)
	})

	//this can be done through always returning float64 values, since it's the highest level of number possible, but requires the got variable to be created as such:
	//got := []float64{sumAllNumbers([]int{1, 2, 3, 4}), sumAllNumbers([]float32{1.0, 2.0, 3.0, 4.0}), ...}
	//which is not the end goal i wanted so i will refrain from it
	//t.Run("summing slices of numbers with various types", func(t *testing.T){
	//	got := sumAllNumbers([]int{1, 2, 3, 4}, []float32{4.5, 5.5})
	//	want := []float32{10.0, 10.0}
	//	assertEquals(t, slices.Equal(got, want), true)
	//})
}

func assertEquals[T comparable](t testing.TB, got, want T){
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
