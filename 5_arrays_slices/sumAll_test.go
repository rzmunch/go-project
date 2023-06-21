package main

import (
	"reflect"
	"testing"
)

func TestSumaAll(t *testing.T) {

	got := SumAll([]int{2, 5}, []int{3, 8})
	want := []int{7, 11}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumaAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{2, 5}, []int{3, 8})
		want := []int{5, 8}

		checkSums(t, got, want)

	})

	t.Run("make the sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 5})
		want := []int{0, 8}

		checkSums(t, got, want)
	})
}
