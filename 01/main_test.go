package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, c := range cases {
		got := CalculateFuel(c.in)
		if got != c.want {
			t.Errorf("CaluculateFuel(%f) == %f, want %f\n", c.in, got, c.want)
		}
	}
}
