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
