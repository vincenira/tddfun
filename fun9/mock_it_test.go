package main

import (
	"bytes"
	"testing"
)

/*
We know we want our Countdown function to write data somewhere and io.Writer is the de-facto
way of capturing that as an interface in Go.
- In main we will send to os.Stdout so our users see the countdown printed to the terminal.
- In test we will send to bytes.Buffer so our tests can capture what data is being generated.
*/
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
