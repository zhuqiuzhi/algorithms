package algorithms

import "testing"

func TestSliceEqual(t *testing.T) {
	tcs := []struct {
		x    [][]string
		y    [][]string
		want bool
	}{
		{
			[][]string{{"test", "slice"}},
			[][]string{{"test", "slice"}},
			true,
		},
		{
			[][]string{{"test", "slice"}},
			[][]string{{"test"}},
			false,
		},
		{
			[][]string{{"test", "slice"}},
			[][]string{{"test", "slice2"}},
			false,
		},
		{
			[][]string{{"test", "slice"}},
			[][]string{{"test"}, {"slice"}},
			false,
		},
	}

	for _, tc := range tcs {
		got := slicesEqual(tc.x, tc.y)
		if got != tc.want {
			t.Errorf("sliceEqual(%v, %v) got %t want %t", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestEqual(t *testing.T) {
	tcs := []struct {
		x    []string
		y    []string
		want bool
	}{
		{
			[]string{"test", "slice"},
			[]string{"test", "slice"},
			true,
		},
		{
			[]string{"test", "slice"},
			[]string{"test"},
			false,
		},
		{
			[]string{"test", "slice"},
			[]string{"test", "slice2"},
			false,
		},
	}

	for _, tc := range tcs {
		got := equal(tc.x, tc.y)
		if got != tc.want {
			t.Errorf("sliceEqual(%v, %v) got %t want %t", tc.x, tc.y, got, tc.want)
		}
	}
}
