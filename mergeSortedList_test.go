package algorithms

import (
	"sort"
	"testing"
)

func TestMergeSortedList(t *testing.T) {
	tcs := []struct {
		sea   []string
		fresh []string
		want  []string
	}{
		{[]string{"Cod", "Herring", "Marlin"}, []string{"Asp", "Carp", "Ide", "Trout"},
			[]string{"Asp", "Carp", "Cod", "Herring", "Ide", "Marlin", "Trout"},
		},
		{nil, []string{"Asp", "Carp", "Ide", "Trout"}, []string{"Asp", "Carp", "Ide", "Trout"}},
		{[]string{"Cod", "Herring", "Marlin"}, nil, []string{"Cod", "Herring", "Marlin"}},
	}

	for _, tc := range tcs {
		got := MergeSortedList(tc.sea, tc.fresh)
		if !sort.IsSorted(sort.StringSlice(got)) {
			t.Errorf("MergeSortedList(%v, %v) not sorted, got %v", tc.sea, tc.fresh, got)
		}
		if !equal(got, tc.want) {
			t.Errorf("MergeSortedList(%v, %v) got %v, want %v", tc.sea, tc.fresh, got, tc.want)
		}
	}
}
