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

	t.Run("say hello if name is input", func(t *testing.T) {
		got := Hello("Daniel", "")
		want := "Hello, Daniel"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello by default if input name is empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say Hola, if language specified is Spanish", func(t *testing.T) {
		got := Hello("Jose", "Spanish")
		want := "Hola, Jose"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say Bonhour, if language specified is French", func(t *testing.T) {
		got := Hello("Pierre", "French")
		want := "Bonjour, Pierre"
		assertCorrectMessage(t, got, want)
	})
}
