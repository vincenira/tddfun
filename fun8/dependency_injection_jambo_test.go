package fun8

import (
	"bytes"
	"testing"
)

func TestMusalimie(t *testing.T) {
	buffer := bytes.Buffer{}
	Musalimie(&buffer, "Chris")

	got := buffer.String()
	want := "Jambo, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
