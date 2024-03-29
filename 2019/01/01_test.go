package main

import "testing"

func TestFuel(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Test fuel formula with 12", func(t *testing.T) {
		got := CalculateFuel(12)
		want := 2
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test fuel formula with 14", func(t *testing.T) {
		got := CalculateFuel(14)
		want := 2
		assertCorrectMessage(t, got, want)
	})

	t.Run("Tests fuel formula with 1969", func(t *testing.T) {
		got := CalculateFuel(1969)
		want := 966
		assertCorrectMessage(t, got, want)
	})

	t.Run("Tests fuel formula with 100756", func(t *testing.T) {
		got := CalculateFuel(100756)
		want := 50346
		assertCorrectMessage(t, got, want)
	})
}
