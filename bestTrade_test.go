package algorithms

import (
	"reflect"
	"testing"
)

func TestTrade(t *testing.T) {
	t.Skip("trade is incomplete")
	testTrade(t, trade, "trade")
}

func TestTradeDP(t *testing.T) {
	testTrade(t, tradeDP, "trade_dp")
}

func TestTradeKadane(t *testing.T) {
	testTrade(t, tradeKadane, "testKnapsack")
}

func testTrade(t *testing.T, testFunc func([]int) []int, funcName string) {
	tcs := []struct {
		in   []int
		want []int
	}{
		{
			[]int{7},
			[]int{0, 0},
		},
		{
			[]int{8, 9},
			[]int{1, 0},
		},
		{
			[]int{8, 7},
			[]int{0, 0},
		},
		{
			[]int{9, 9},
			[]int{0, 0},
		},
		{
			[]int{27, 53, 7, 25, 33, 2, 32, 47, 43},
			[]int{7, 5},
		},
	}

	for _, tc := range tcs {
		got := testFunc(tc.in)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s(%v) got %v, want %v", funcName, tc.in, got, tc.want)
		}
	}

}
