package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	cases := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 36.0},
		{Circle{1}, 2 * math.Pi},
	}

	for _, testcase := range cases {
		got := testcase.shape.Perimeter()

		if got != testcase.want {
			t.Errorf("case %#v: got %.2f want %.2f", testcase, got, testcase.want)
		}
	}
}
