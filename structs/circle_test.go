package structs

import (
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {
	cases := []struct {
		in   Circle
		want float64
	}{
		{Circle{R: 0}, 0},
		{Circle{R: 1}, math.Pi},
		{Circle{R: 2}, (4 * math.Pi)},
		{Circle{R: 3}, (9 * math.Pi)},
		{Circle{R: 4}, (16 * math.Pi)},
	}
	for _, c := range cases {
		got := c.in.Area()
		if got != c.want {
			t.Errorf("Circle.Area() function of %+v returned %g, want %g", c.in, got, c.want)
		}
	}
}
