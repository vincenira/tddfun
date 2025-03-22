package main

import "testing"

func TestJambo(t *testing.T) {
	// %q placeholder wrap your value into double quotes
	// Reference from Go in fmt printing symbol https://pkg.go.dev/fmt#hdr-Printing
	t.Run("Default: Greetings in swahili", func(t *testing.T) {
		got := Jambo("")
		want := "Jambo, wetu"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("English: Greeting", func(t *testing.T) {
		got := Jambo("wetu")
		want := "Hello, wetu"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Spanish: Greetings", func(t *testing.T) {
		got := Jambo("wetu")
		want := "Hola, wetu"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("French: Greetings", func(t *testing.T) {
		got := Jambo("wetu")
		want := "Salut, wetu"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Chinese: Greetings", func(t *testing.T) {
		got := Jambo("wetu")
		want := "Nǐ hǎo, Wetu"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
