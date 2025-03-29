package fun5

// initial commit with test case and the minimal amount of code for test to run
// and check the failing test output
/*
func Perimeter(x, y float64) float64 {
	return 0.0
}
*/

// refactor the code
// since nothing indicates, it is a rectangle. Let's create our own type using struct

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.141592653589793
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.width + rect.height)
}

/*
since we can not have two Area function in the same package, we have to use the methods

We have two choices:

You can have functions with the same name declared in different packages.
So we could create our Area(Circle) in a new package, but that feels overkill here.
We can define methods on our newly defined types instead.

	func Area(rect Rectangle) float64 {
		return rect.width * rect.height
	}

	func Area(circ Circle) float64 {
		return 2 * Circ.Radius * 3.141592653589793
	}
*/
func Area(rect Rectangle) float64 {
	return rect.width * rect.height
}
