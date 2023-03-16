package main

import "testing"

// Testing the Hello method that should return "Hello, world"
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Checking that HelloName is returning a "Hello, $NAME" message
func TestHelloName(t *testing.T) {
	got := HelloName("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}