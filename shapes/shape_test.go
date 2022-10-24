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
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36},
	}

	for _, ex := range areaTests {
		got := ex.shape.Area()
		if got != ex.want {
			t.Errorf("got %g want %g", got, ex.want)
		}
	}
}
