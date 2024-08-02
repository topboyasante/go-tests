package smi

import "testing"

//SMI : Structs, Models and Interfaces

func TestArea(t *testing.T) {

	// Table Driven Tests
	// They are used when we want to run a set of similar tests.
	// In our case, these tests all check the area of a shape
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, testCase := range areaTests {
		got := testCase.shape.Area()
		if got != testCase.hasArea {
			t.Errorf("%#v got %g want %g", testCase.shape, got, testCase.hasArea)
		}
	}
}


func BenchmarkArea(b *testing.B) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _,testCase := range areaTests{
		b.Run(testCase.name, func(b *testing.B){
			for i := 0; i < b.N; i++ {
				testCase.shape.Area()
			}
		})
	}
}