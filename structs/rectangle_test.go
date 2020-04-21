package structs

import "testing"

func TestRectangleArea(t *testing.T) {
	cases := []struct {
		in   Rectangle
		want float64
	}{
		{Rectangle{L: 0, W: 0}, 0},
		{Rectangle{L: 1, W: 3}, 3},
		{Rectangle{L: 2, W: 6}, 12},
		{Rectangle{L: 3, W: 9}, 27},
		{Rectangle{L: 4, W: 12}, 48},
		{Rectangle{L: 5, W: 15}, 75},
	}
	for _, c := range cases {
		got := c.in.Area()
		if got != c.want {
			t.Errorf("Rectangle.Area() function of %+v returned %g, want %g", c.in, got, c.want)
		}
	}
}
