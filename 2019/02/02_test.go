package main

import "testing"

func TestRunInstructions(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("First element changes", func(t *testing.T) {
		instructions := []int{1, 0, 0, 0, 99}
		got := RunInstructions(instructions)
		want := 2
		assertCorrectMessage(t, got, want)
	})

	t.Run("No changes", func(t *testing.T) {
		instructions := []int{2, 3, 0, 3, 99}
		got := RunInstructions(instructions)
		want := 2
		assertCorrectMessage(t, got, want)
	})

	t.Run("Longer Intcode", func(t *testing.T) {
		instructions := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
		got := RunInstructions(instructions)
		want := 30
		assertCorrectMessage(t, got, want)
	})
}
