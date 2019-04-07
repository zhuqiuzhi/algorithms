package sort

import (
	"sort"
	"testing"
)

func testSortFunc(t *testing.T, sortFunc func(sort.Interface), name string) {
	tcs := []sort.Interface{
		sort.IntSlice{},
		sort.IntSlice{3},
		sort.IntSlice{3, 3, 3, 3},
		sort.IntSlice{3, 3, 3, 3, 3},
		sort.IntSlice{3, 1, 4, 1, 5, 9, 2, 6, 5},
		sort.IntSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		sort.IntSlice{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		sort.StringSlice{"base", "ball", "mound", "bat", "glove", "batter"},
	}

	for _, tc := range tcs {
		sortFunc(tc)
		if !sort.IsSorted(tc) {
			t.Errorf("%s(%v) failed", name, tc)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	testSortFunc(t, SelectionSort, "SelectionSort")
}

func TestRecSelectionSort(t *testing.T) {
	testSortFunc(t, RecSelectionSort, "RecSelectionSort")
}

func TestQuickSort(t *testing.T) {
	testSortFunc(t, QuickSort, "QuickSort")
}
