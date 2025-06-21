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

	return Reduce(numbers, sumTail, []int{})
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

func SumAllTails(sliceNumbers ...[]int) []int {
	var sums []int
	var sum int
	for _, sliceNumber := range sliceNumbers {
		sum = 0
		for i, number := range sliceNumber {
			if i == 0 {
				continue
			}
			sum += number
		}
		sums = append(sums, sum)
	}
	return sums
}
