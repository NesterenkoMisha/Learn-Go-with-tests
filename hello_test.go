package main

import "testing"

func assertMethodsOutput(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf("Actual: %q ; But expected: %q", actual, expected)
	}
}

// Using sub-tests for checking all possible scenarios for the Hello(sting, string) method
func TestHelloNameAgain(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello("Chris", "")
		expected := "Hello, Chris"

		assertMethodsOutput(t, actual, expected)
	})
	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, world"

		assertMethodsOutput(t, actual, expected)
	})
	t.Run("in Spanish", func(t *testing.T) {
		actual := Hello("Elide", "Spanish")
		expected := "Hola, Elide"

		assertMethodsOutput(t, actual, expected)
	})
}
