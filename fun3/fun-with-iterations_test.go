package fun3

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
