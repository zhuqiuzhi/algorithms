package algorithms

import (
	"testing"
)

func TestSolveNQueensOneSolution(t *testing.T) {
	tcs := []struct {
		n    int
		want []string
	}{
		{1, []string{"Q"}},
		{2, []string{}},
		{3, []string{}},
		{4, []string{
			".Q..",
			"...Q",
			"Q...",
			"..Q.",
		}},
	}
	for _, tc := range tcs {
		got := SolveNQueensOneSolution(tc.n)
		if !equal(got, tc.want) {
			t.Errorf("solveNQueensOneSolution(%d) got %v want %v", tc.n, got, tc.want)
		}
	}
}

func TestSolveNQueens(t *testing.T) {
	testSolveNQueens(t, solveNQueens)
}

func TestSolveNQueensUseSet(t *testing.T) {
	testSolveNQueens(t, solveNQueensUseSet)
}

func testSolveNQueens(t *testing.T, testFunc func(n int) [][]string) {
	tcs := []struct {
		n    int
		want [][]string
	}{
		{1, [][]string{{"Q"}}},
		{2, [][]string{}},
		{3, [][]string{}},
		{4, [][]string{
			{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			},
			{
				"..Q.",
				"Q...",
				"...Q",
				".Q..",
			},
		}},
	}
	for _, tc := range tcs {
		got := testFunc(tc.n)
		if !slicesEqual(got, tc.want) {
			t.Errorf("solutions to the %d-queens puzzle got %v want %v", tc.n, got, tc.want)
		}
	}
}
