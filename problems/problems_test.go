package problems

import "testing"

func TestFibonacciSequencer(t *testing.T) {
	cases := []struct {
		in, want uint
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
	}
	for _, c := range cases {
		got := FibonacciSequencer(c.in)
		if got != c.want {
			t.Errorf("FibonacciSequencer(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
