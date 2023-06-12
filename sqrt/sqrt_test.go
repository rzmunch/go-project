package main

import "testing"

func TestSqrt(t *testing.T) {
	t.Run("Test sqrt 4", func(t *testing.T) {
		got := Sqrt(4)
		var want float64 = 2
		assertCorrectMessageFloat(t, got, want)
	})
	t.Run("Test sqrt 25", func(t *testing.T) {
		got := Sqrt(25)
		var want float64 = 5
		assertCorrectMessageFloat(t, got, want)
	})
	t.Run("Test sqrt 10000", func(t *testing.T) {
		got := Sqrt(10000)
		var want float64 = 100
		assertCorrectMessageFloat(t, got, want)
	})
	t.Run("Test sqrt 1000000", func(t *testing.T) {
		got := Sqrt(1000000)
		var want float64 = 1000
		assertCorrectMessageFloat(t, got, want)
	})
}

func assertCorrectMessageFloat(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
