package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Grace")
	})

	//	t.Run("asserting different types", func(t *testing.T) {
	//		AssertEqual(t, 1, "1")
	//	}) // uncomment to see the error
}

// This allows only integer to be tested
// because of the argument and %d format specifier with represent a decimal integer
func Oldv1AssertEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func Oldv1AssertNotEqual(t *testing.T, got, want int) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %d", got)
	}
}

// If your argument is of type interface{}, it means "anything" /any types can be passed
// The compiler can not help for type-safety with interface{}
// for instance AssertEqual(t, 1, "1") can pass the guard rails
func Oldv2AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func Oldv2AssertNotEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %+v", got)
	}
}

/*
To write generic functions in Go, you need to provide "type parameters" which is just a fancy
way of saying "describe your generic type and gitve it a label"
we used the comparable because we want to describe to the compiler that we wish to use the ==
and != operators on things of type T in our function, we want to compare
*/
func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}
