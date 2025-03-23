package fun2

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Adder(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

/*
Extra mile, let's add an Example function which can be use for go doc
Example function begins with the a prefix name Example like test function begin with test
and reside in a package's _test.go
it requires the special comment
// Output: value_output
*/

func ExampleAdder() {
	sumOfTwoIntegers := Adder(3, 9)
	fmt.Printf("%d", sumOfTwoIntegers)
	// Output: 12
}
