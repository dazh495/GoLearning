package main

import "testing"

func TestHello(t *testing.T) {
	//testing TB is the interface that testing T and B inherit from
	//t.Helper lets the function know that this is a Helper function. This way,
	//the function will debug when the function is CALLED. not within the function itself.
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Daniel")
		want := "Hello, Daniel"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello by default", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
