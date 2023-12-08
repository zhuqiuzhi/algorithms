package dynamic

import "testing"

func TestChange2(t *testing.T) {
	testChange2(t, change2, "change2")
}

func testChange2(t *testing.T, testFunc func(int, []int) int, name string) {
	tcs := []struct {
		amount int
		coins  []int
		want   int
	}{
		{5, []int{1, 2, 5}, 4},
		{5, []int{1, 2, 3}, 5},
		{3, []int{2}, 0},
		{10, []int{10}, 1},
	}
	for _, tc := range tcs {
		got := testFunc(tc.amount, tc.coins)
		if got != tc.want {
			t.Errorf("%s(%d,%v) got %d want %d", name, tc.amount, tc.coins, got, tc.want)
		}
	}

}
