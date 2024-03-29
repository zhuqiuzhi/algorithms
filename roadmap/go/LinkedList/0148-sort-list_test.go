package roadmap

import (
	"flag"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func getInts(start, end int) []int {
	size := end - start + 1
	var a = make([]int, size)

	for i, j := 0, start; i < size; i++ {
		a[i] = j
		j++
	}
	return a
}

var isRand bool

func init() {
	flag.BoolVar(&isRand, "rand", false, "pseudo-randomizes the order of elements")
}

func TestSortList(t *testing.T) {
	tcs := []struct {
		in   []int
		want []int
		f    func(*ListNode) *ListNode
	}{
		{in: []int{4, 2, 1, 3}, want: []int{1, 2, 3, 4}, f: sortList},
		{in: []int{-1, 5, 3, 4, 0}, want: []int{-1, 0, 3, 4, 5}, f: sortList},
		{in: []int{2, 3, 4, 5, 6, 7, 8, 9, 1}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, f: sortList},
		{in: getInts(1, 50000), want: getInts(1, 50000), f: sortList},
		{in: append(getInts(2, 50000), 1), want: getInts(1, 50000), f: sortList},
		{in: append(getInts(2, 100000), 1), want: getInts(1, 100000), f: sortList},

		{in: []int{4, 2, 1, 3}, want: []int{1, 2, 3, 4}, f: mergeSortList},
		{in: []int{-1, 5, 3, 4, 0}, want: []int{-1, 0, 3, 4, 5}, f: mergeSortList},
		{in: []int{2, 3, 4, 5, 6, 7, 8, 9, 1}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, f: mergeSortList},
		{in: getInts(1, 50000), want: getInts(1, 50000), f: mergeSortList},
		{in: append(getInts(2, 50000), 1), want: getInts(1, 50000), f: mergeSortList},
		{in: append(getInts(2, 100000), 1), want: getInts(1, 100000), f: mergeSortList},
	}

	if isRand {
		t.Logf("isRand: %t", isRand)
	}

	for i, tc := range tcs {
		begin := time.Now().UnixMilli()
		if isRand {
			rand.Shuffle(len(tc.in), func(i, j int) {
				tc.in[i], tc.in[j] = tc.in[j], tc.in[i]
			})
		}

		list := NewList(tc.in)
		sortedList := tc.f(list)
		values := Values(sortedList)
		if !reflect.DeepEqual(tc.want, values) {
			t.Errorf("origin: %v, got: %v", tc.in, values)
		}
		t.Logf("case %d: %d milli seconds", i, time.Now().UnixMilli()-begin)
	}
}
