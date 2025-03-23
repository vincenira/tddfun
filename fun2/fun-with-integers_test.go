package fun2

import (
	"testing"
)

func AdderTest(t *testing.T) {
	got := Adder(2, 2)
	want := 4
	t.Errorf("got %d want %d", got, want)
}
