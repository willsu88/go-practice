package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Chris", "")
		want := "Hello Chris"

		assertCorrectMessage(got, want, t)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola Elodie"

		assertCorrectMessage(got, want, t)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour Elodie"

		assertCorrectMessage(got, want, t)
	})

	t.Run("saying hello to noone", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(got, want, t)
	})

}

func assertCorrectMessage(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("Error: Expected %q, Got %q", want, got)
	}
}
