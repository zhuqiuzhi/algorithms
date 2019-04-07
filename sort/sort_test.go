package sort

import (
	"sort"
	"testing"
)

func testSortFunc(t *testing.T, sortFunc func(sort.Interface)) {
	tcs := []sort.Interface{
		sort.IntSlice{},
		sort.IntSlice{3},
		sort.IntSlice{3, 1, 4, 1, 5, 9, 2, 6, 5},
		sort.StringSlice{"base", "ball", "mound", "bat", "glove", "batter"},
	}

	for _, tc := range tcs {
		sortFunc(tc)
		if !sort.IsSorted(tc) {
			t.Errorf("SelectionSort(%v) failed", tc)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	testSortFunc(t, SelectionSort)
}

func TestRecSelectionSort(t *testing.T) {
	testSortFunc(t, RecSelectionSort)
}
