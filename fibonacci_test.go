package algorithms

import "testing"

func TestFibonacci(t *testing.T) {
	tcs := []struct {
		in   int
		want int
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
		{12, 144},
	}
	for _, tc := range tcs {
		got := Fibonacci(tc.in)
		if got != tc.want {
			t.Errorf("Fibonacci(%d) got %d want %d", tc.in, got, tc.want)
		}
	}

}
