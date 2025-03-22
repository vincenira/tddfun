package main

import "testing"

func TestJambo(t *testing.T) {
	// %q placeholder wrap your value into double quotes
	// Reference from Go in fmt printing symbol https://pkg.go.dev/fmt#hdr-Printing
	t.Run("Default: Greetings in swahili", func(t *testing.T) {
		got := Jambo("", "")
		want := "Jambo, wetu"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Swahili: Greetings in swahili", func(t *testing.T) {
		got := Jambo("wetu", "Swahili")
		want := "Jambo, wetu"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("English: Greeting", func(t *testing.T) {
		got := Jambo("wetu", "English")
		want := "Hello, wetu"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("Spanish: Greetings", func(t *testing.T) {
		got := Jambo("wetu", "Spanish")
		want := "Hola, wetu"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("French: Greetings", func(t *testing.T) {
		got := Jambo("wetu", "French")
		want := "Salut, wetu"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("Chinese: Greetings", func(t *testing.T) {
		got := Jambo("wetu", "Chinese")
		want := "Nǐ hǎo, wetu"
		assertCorrectGreeting(t, got, want)
	})
}

// we pass as argument t with type testing.TB to allow the assertCorrectGreeting()
// to access error information from the entire testing code when it fails
func assertCorrectGreeting(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
