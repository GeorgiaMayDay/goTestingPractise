package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}

	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36},
	}

	for _, ex := range areaTests {
		t.Run(ex.name, func(t *testing.T) {
			got := ex.shape.Area()
			if got != ex.want {
				t.Errorf("In case %v, got %g want %g", ex.shape, got, ex.want)
			}
		})
	}
}
