package algorithms

import (
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tcs := []sort.Interface{
		sort.IntSlice{3, 1, 4, 1, 5, 9, 2, 6, 5},
		sort.StringSlice{"base", "ball", "mound", "bat", "glove", "batter"},
	}

	for _, tc := range tcs {
		SelectionSort(tc)
		if !sort.IsSorted(tc) {
			t.Errorf("SelectionSort(%v) failed", tc)
		}
	}
}
