package main

import (
	"os/exec"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"key":"value"}
	got := Search(dictionary, "key")

	want := "value"
	if got != want{
		t.Errorf("got %q want %q" , got , want)
	}
}

func assertString(t testing.TB, got, want string){
	t.Helper()

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case exec.ErrNotFound:
		d[word] = definition
	case nil:
		return exec.ErrNotFound
	default:
		return err
	}

	return nil
}