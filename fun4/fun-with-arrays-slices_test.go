// fun4: arrays and slices
/*
go mod init main you will presented with an error _testmain.go:13.2: cannot import "main".
package main will only contain integration of other packages and not unit-testable code
hence Go will not allow you to import a package with name main
*/
package fun4

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumSlice(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := SumSlice(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := SumSlice(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	// you can compare a slice using the equality sign we need the reflect.DeepEqual
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

/*
refactor the repeated parts of the testing code with a checksum function helper
*/
func TestSumAllTails(t *testing.T) {
	/*
		It's not shown here, but this technique can be useful when you want to bind a function to other
		local variables in "scope" (e.g between some {}). It also allows you to reduce the surface area of your API.
	*/
	checksum := func(t testing.TB, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Make the sums of some integer slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		// you can compare a slice using the equality sign we need the reflect.DeepEqual
		checksum(t, got, want)
	})
	t.Run("Make the sums of an empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checksum(t, got, want)
	})
}
