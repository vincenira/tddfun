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

/*
Learning the table driven tests

table driven tests are used when you want to make all list of test cases which can be tested
in the same manner
*/

func TestTDTArea(t *testing.T) {
	// Anonymous struct which is a slice of structs by using []struct with two fields, the shape and the want
	AreaTests := []struct {
		shape Shape
		want  float64
	}{
		// composite literal declaration when need a comma for the last element
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		// Add a test for the new shape the triangle
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range AreaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

// The above struct needs some refactor because it is not immediately clear
// what all the numbers represent and we should be aiming for your tests to be
// easily understood
// In Test-Driven Development by Example Kent Beck refactors some tests to a point and asserts:
//    "The test speaks to us more clearly, as if it were an assertion of truth, not a
//     sequence of operations"
// Now our tests - Rather, the list of test cases - make assertions of truth about shapes
// and their areas.
/*
The below test function is a refacted code of the TestTDTArea to solve the following problems:
- To assert the truth, not to have a sequence of operations
- Allow the developer to know rapidly, the type of its fields when it fails
- Increase readability of the test cases
- By wrapping the testcases into t.Run we can run specify test
	you can run specific tests within your table with go test -run TestTDTRefactedArea/Rectangle
*/

func TestTDTRefactedArea(t *testing.T) {
	AreaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangletoo", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range AreaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				// %#v format stirng will print out our struct with the values in its fields
				// so the developer can see at glance the properties that are being tested
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
