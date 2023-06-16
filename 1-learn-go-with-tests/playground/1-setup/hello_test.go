package main

import (
	"testing"
)

func testEquals(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestHello(t *testing.T) {
	t.Run("hello with an empty string", func(t *testing.T) {
		testEquals(t, Hello("", ""), "Hello, world!")
	})

	t.Run("hello to people", func(t *testing.T) {
		testEquals(t, Hello("Nancy", ""), "Hello, Nancy!")
	})

	t.Run("hello in a different language", func(t *testing.T) {
		testEquals(t, Hello("John", "Spanish"), "Hola, John!")
	})

	t.Run("hello in a different language", func(t *testing.T) {
		testEquals(t, Hello("John", "Japanese"), "Ohaiyoo, John!")
	})

	t.Run("hello in a different language", func(t *testing.T) {
		testEquals(t, Hello("John", "Chinese"), "Nihao, John!")
	})
}
