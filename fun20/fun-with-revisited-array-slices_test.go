package fun20

import (
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})
}
