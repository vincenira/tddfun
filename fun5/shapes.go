package fun5

// initial commit with test case and the minimal amount of code for test to run
// and check the failing test output
/*
func Perimeter(x, y float64) float64 {
	return 0.0
}
*/

// refactor the code

func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

func Area(width, height float64) float64 {
	return width * height
}
