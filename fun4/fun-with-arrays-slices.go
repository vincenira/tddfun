// fun4: arrays and slices
package fun4

func Sum(numbers [5]int) int {
	sum := 0
	/*
	   for i := 0; i < 5; i++ {
	   refactor using the for range
	*/
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumSlice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
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
