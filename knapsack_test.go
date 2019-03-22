package algorithms

import (
	"reflect"
	"testing"
)

func TestBruteForceKnapsack(t *testing.T) {
	testKnapsack(t, BruteForceKnapsack)
}

//     10,[2,3,5,7],[1,5,2,4]
func testKnapsack(t *testing.T, testFunc func(items []product, maxWeight float64) ([]product, float64)) {
	tcs := []struct {
		items     []product
		maxWeight float64
		want      []product
		wantValue float64
	}{
		{
			items: []product{
				{"apple", 2, 1}, {"pen", 3, 5},
				{"orange", 5, 2}, {"book", 7, 4},
			},
			maxWeight: 10,
			want: []product{
				{"pen", 3, 5}, {"book", 7, 4},
			},
			wantValue: 9,
		},
	}

	for _, tc := range tcs {
		got, bestValue := testFunc(tc.items, tc.maxWeight)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("BruteForceKnapsack(%v, %v) got product %v, want product %v", tc.items, tc.maxWeight, got, tc.want)
		}
		if bestValue != tc.wantValue {
			t.Errorf("BruteForceKnapsack(%v, %v) got value %f, want value %f", tc.items, tc.maxWeight, bestValue, tc.wantValue)
		}
	}
}

func TestPowerSet(t *testing.T) {
	testPowerSet(t, PowerSet)
}

func TestRecursivePowerSet(t *testing.T) {
	// TODO
	t.Skip("RecursivePowerSet is incomplete")
	testPowerSet(t, RecursivePowerSet)
}

func testPowerSet(t *testing.T, testFunc func(ps []product) []products) {
	tcs := []struct {
		in     []product
		number int
	}{
		{
			[]product{
				{"apple", 2, 1}, {"pen", 3, 5},
				{"orange", 5, 2}, {"book", 7, 4},
			}, 16,
		},
	}
	for _, tc := range tcs {
		got := testFunc(tc.in)
		if len(got) != tc.number {
			t.Errorf("product set is %v, got the number of set in power set is %d, want %d", tc.in, len(got), tc.number)
		}
	}

}
