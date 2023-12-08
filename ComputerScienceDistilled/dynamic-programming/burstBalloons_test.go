package dynamic

import "testing"

func TestMaxCoins(t *testing.T) {
	tcs := []struct {
		in   []int
		want int
	}{
		{[]int{3, 1, 5, 8}, 167},
		{[]int{1, 4, 9, 2, 7, 3}, 482},
	}
	for _, tc := range tcs {
		got := maxCoins(tc.in)
		if got != tc.want {
			t.Errorf("maxCoins(%v) got %d want %d", tc.in, got, tc.want)
		}
	}

}
