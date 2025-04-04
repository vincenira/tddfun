package fun7

import "testing"

func TestSearch(t *testing.T) {
	/*
	   	dictionary := map[string]string{"test": "This is just a test"}
	     Refactor: dictionnary to have search as a method of dictionnary whereas dictionnary
	     is a type.
	*/
	dictionary := Dictionary{"test": "This is just a test"}

	got, _ := dictionary.Search("test")
	want := "This is just a test"
	assertStrings(t, got, want)
	/*
	   Implementing a logic when the word is not in the dictionary
	*/
	t.Run("Know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is just a test"
		assertStrings(t, got, want)
	})
	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "Could not find the word you were looking for"
		if err == nil {
			t.Fatal("Expected to get an error")
		}
		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
