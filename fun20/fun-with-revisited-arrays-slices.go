package fun20

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return ReduceR(numbers, sumTail, []int{})
}

/*
e.g: make([]int, 0, 5) creates a slice with length 0 and capacity 5.
*/
func SumAll(sliceNumbers ...[]int) []int {
	var sumSlice []int
	var sum int
	for _, sliceNumber := range sliceNumbers {
		sum = 0
		for _, number := range sliceNumber {
			sum += number
		}
		sumSlice = append(sumSlice, sum)
	}
	return sumSlice
}

func ReduceR[A any](collection []A, f func(A, A) A, initialValue A) A {
	result := initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	result := initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}
