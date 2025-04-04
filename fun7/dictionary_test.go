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
		_, got := dictionary.Search("unknown")
		if got == nil {
			t.Fatal("Expected to get an error")
		}
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "This is just a test")

	want := "This is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word", err)
	}
	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
