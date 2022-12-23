package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("In spanish", func(t *testing.T) {
		got := Hello("Abdessamed", "spanish")
		want := "Hola, Abdessamed"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Abdessamed", "english")
		want := "Hello, Abdessamed"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}