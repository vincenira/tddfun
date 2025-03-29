// fun5: struct, methods and interfaces
package fun5

import "testing"

/*
func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
*/
//refactored testing code for the rectangle type

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	// got := Perimeter(rectangle.width, rectangle.height)
	// refactor to pass the object Rectangle as an argument
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	// got := Area(rectangle.width, rectangle.height)
	// refactor to pass the object Rectangle as an argument
	got := Area(rectangle)
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Adding the circle calculation

func TestGeometryShapes(t *testing.T) {
	t.Run("Calculation the Area of Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("Calculation the Area of Circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793
		// f format string was replaced by g which is more precise
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

/*
Notice how our helper does not need to concern itself with whether the shape is a Rectangle or a Circle or a Triangle. By declaring an interface,
the helper is decoupled from the concrete types and only has the method it needs to do its job.
*/

func TestShapeArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Calculate Area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})
	t.Run("Calculate Area of a circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
