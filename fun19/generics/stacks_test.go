package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(StacksOfInts)
		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())
	})
}
