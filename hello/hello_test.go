package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying learn-go-with-tests to people", func(t *testing.T) {
		got := Hello("Marcello", "English")
		want := "Hello, Marcello!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when an empty name string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "¡Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("José", "Portuguese")
		want := "Olá, José!"
		assertCorrectMessage(t, got, want)
	})
}

func ExampleHello() {
	fmt.Println(Hello("John", "English"))
	// Output: Hello, John!
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
