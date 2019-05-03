package algorithms

import (
	"reflect"
	"testing"
)

func TestTwoSum3(t *testing.T) {
	tcs := []struct {
		in     []int
		target int
		want   [][]int
	}{
		{[]int{}, 9, nil},
		{[]int{1}, 9, nil},
		{[]int{2, 7, 11, 15}, 9, [][]int{{2, 7}}},
		{[]int{2, 7, 5, 4, 11, 15}, 9, [][]int{{2, 7}, {4, 5}}},
		{[]int{2, 7, 11, 15, 5, 4}, 9, [][]int{{2, 7}, {4, 5}}},
	}

	for _, tc := range tcs {
		if got := TwoSum3(tc.in, tc.target); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("TwoSum(%v,%d) got %v, want %v", tc.in, tc.target, got, tc.want)
		}
	}

}
