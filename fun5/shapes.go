package fun5

import "math"

// initial commit with test case and the minimal amount of code for test to run
// and check the failing test output
/*
func Perimeter(x, y float64) float64 {
	return 0.0
}
*/
// let's tell Go what a Shape is using an interface

// In go interface resolution is implicit. if the ype you pass in matches what the inteface
// is asking for, it will compile
type Shape interface {
	Area() float64
}

// refactor the code
// since nothing indicates, it is a rectangle. Let's create our own type using struct

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// it is a convention in Go to have the receiver variable be the first letter of the type
// for example: r Rectange(since the letter r is the first letter of the Type Rectangle)
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

/*
There is some duplication in our tests
all we want to do is take a collection of shapes, call the Area() method on them and then
check the result
*/
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Add a new shape Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
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
	return rect.Width * rect.Height
}
