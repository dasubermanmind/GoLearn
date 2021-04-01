package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "HelloChris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("","")
		want := "HelloWorld"
		assertCorrectMessage(t, got, want)
	})


	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Eliot", "Spanish")
		want := "HolaEliot"
		assertCorrectMessage(t, got, want)
	})
}


