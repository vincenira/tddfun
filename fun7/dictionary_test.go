package fun7

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "This is just a test"}

	got := Search(dictionary, "test")
	want := "This is just a test"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
