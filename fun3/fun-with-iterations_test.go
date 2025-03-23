package fun3

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

/*
the benchmark feature is provided by the standard library in Golang
Benchmark library manages for you, the number of times to run your code
through the variable b.N
And to run the benchmark function "go test -bench=."
Note:
- that the benchmark and testing functions can coexist in the same _test.go file
- By default benchmarks are run sequentially
*/
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

/*
we are going to append all the code with duplicated function with a prefix EX
To do the practical exercises
*/

func TestExRepeat(t *testing.T) {
	repeated := ExRepeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("Expected %q, but got %q", expected, repeated)
	}
}
