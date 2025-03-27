/*
go mod init main you will presented with an error _testmain.go:13.2: cannot import "main".
package main will only contain integration of other packages and not unit-testable code
hence Go will not allow you to import a package with name main
*/
package fun4

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
