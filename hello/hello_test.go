package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Yury")
	want := "Hello, Yury"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
