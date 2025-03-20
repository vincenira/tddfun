package main

import "testing"

func TestJambo(t *testing.T) {
	got := Jambo()
	want := "Jambo, wetu"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
