package algorithms

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tcs := []struct {
		in     []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{11, 7, 2, 15}, 9, []int{1, 2}},
	}

	for _, tc := range tcs {
		if got := TwoSum(tc.in, tc.target); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("TwoSum(%v,%d) got %v, want %v", tc.in, tc.target, got, tc.want)
		}
	}

}
