// Given an integer n, return all distinct solutions to the n-queens puzzle.
// Each solution contains a distinct board configuration of the n-queens' placement,
// where 'Q' and '.' both indicate a queen and an empty space respectively.
// Leetcode: https://leetcode-cn.com/problems/n-queens/
// The queen is the most powerful piece: it can attack pieces that occupy its row,
// its column, or its diagonals. How do you place eight queens on the board
// Demo: https://code.energy/8-queens-problem
// Solution:
// That’s the essence of the backtracking strategy: keep on placing queens in valid positions.
// Once we get stuck, roll back the placement of the last queen and carry on.
// The process can be streamlined using recursion.
// 位于（i,j） 的皇后能攻击第 i 行、第 j 列的棋子、行数和列数之和等于 i+j 的棋子(正斜线上的)、行数和列数之差等于 i-j
// 的棋子. 左上角的第一个格子是(0，0),解中的皇后的位置只要记录列，因为皇后肯定分布在每一行中
package algorithms

import (
	"fmt"
	"strings"
)

func String(n, col int) string {
	return fmt.Sprintf("%sQ%s", strings.Repeat(".", col), strings.Repeat(".", n-col-1))
}

type position struct {
	row, col int
}

// solveNQueensOneSolution 找出 N-皇后一个解
// 比较好的讲解寻找一个解的视频教程：https://www.youtube.com/watch?v=xouin83ebxE, 它通过遍历 position 对象数组
// 判断皇后是否会被已经放置在 position 位置的皇后攻击
func SolveNQueensOneSolution(n int) []string {
	var queens = make([]position, 0, n)
	hasSolution := solveNQueensOneSolutionUtil(n, 0, &queens)

	var result []string
	if hasSolution {
		for _, q := range queens {
			result = append(result, String(n, q.col))
		}
	}

	return result
}

// solveNQueensOneSolutionUtil 找出第 row 行的第一个可以安全放置皇后的列追加到 colsOfQueen
// 当 row 大于或等于 n 说明找到解的所有的皇后的位置
// 因为递归的过程中会修改 queues，所以必须返回修改后的值
func solveNQueensOneSolutionUtil(n, row int, queens *[]position) bool {
	if row >= n {
		return true
	}

	for col := 0; col < n; col++ {
		var foundSafe = true
		for _, q := range *queens {
			if col == q.col || row+col == q.col+q.row || row-col == q.row-q.col {
				// 会被已放置的皇后攻击
				foundSafe = false
				break
			}
		}
		if foundSafe {
			// 递归到下一行
			var ok bool
			*queens = append(*queens, position{row: row, col: col})
			ok = solveNQueensOneSolutionUtil(n, row+1, queens)
			if ok {
				// 只有当递归完到最后一行，solveNQueensOneSolutionUtil 才会返回 true，此时说明找到了一个解
				return true
			} else {
				// Youtube 中视频没有做这步操作是因为他每次都是按行来放置皇后的，而不是使用追加的方式
				// 下一行找不到安全的位置，要回溯上一次的操作,即把上一次放置的皇后移除
				*queens = (*queens)[:len(*queens)-1]
			}
		}
	}
	// 枚举完当前行的列, 没有找到安全的位置
	return false
}

// 此处仍使用类似上面找一个解的方式来找到每一行安全的位置
func solveNQueens(n int) [][]string {
	var queens = make([]position, 0, n)
	// 所有的可行解的集合
	var allSolutions [][]position

	solveNQueensUtil(n, 0, queens, &allSolutions)
	var result [][]string
	for _, s := range allSolutions {
		var solution []string
		for _, q := range s {
			solution = append(solution, String(n, q.col))
		}
		result = append(result, solution)
	}

	return result
}

// solveNQueensOneSolutionUtil 找出第 row 行的第一个可以安全放置皇后的列追加到 colsOfQueen
// 当 row 大于或等于 n 说明找到解的所有的皇后的位置
// 因为递归的过程中会修改 queues，所以必须返回修改后的值
func solveNQueensUtil(n, row int, queens []position, allSolutions *[][]position) {
	if row >= n {
		*allSolutions = append(*allSolutions, queens)
		return
	}

	for col := 0; col < n; col++ {
		var foundSafe = true
		for _, q := range queens {
			if col == q.col || row+col == q.col+q.row || row-col == q.row-q.col {
				// 会被已放置的皇后攻击
				foundSafe = false
				break
			}
		}
		if foundSafe {
			queens = append(queens, position{row: row, col: col})

			// 注意此种解法，必须拷贝 slice, 因为递归结束时，如果行数已经等于 n 时，会将 slice 添加到 allSolutions 里
			var copyQueens = make([]position, len(queens))
			copy(copyQueens, queens)
			solveNQueensUtil(n, row+1, copyQueens, allSolutions)

			// 回溯
			queens = queens[0 : len(queens)-1]
		}
	}
}

type board struct {
	cols      map[int]struct{}
	sum       map[int]struct{}
	diff      map[int]struct{}
	solutions [][]int
}

func (b *board) AddQueen(row, col int) {
	b.cols[col] = struct{}{}
	b.sum[row+col] = struct{}{}
	b.diff[row-col] = struct{}{}
}

func (b *board) RemoveQueen(row, col int) {
	delete(b.cols, col)
	delete(b.sum, row+col)
	delete(b.diff, row-col)
}

func (b *board) IsSafe(row, col int) bool {
	if _, ok := b.cols[col]; ok {
		return false
	}

	if _, ok := b.sum[row+col]; ok {
		return false
	}

	if _, ok := b.diff[row-col]; ok {
		return false
	}

	return true
}

// 视频 https://time.geekbang.org/course/detail/130-67638  中使用三个 set 来记录可以被皇后攻击的位置,
// 然后回溯时将之前的放置的皇后可以攻击的位置从 set 中移除，最后找到所有解
func solveNQueensUseSet(n int) [][]string {
	var b = &board{
		cols: make(map[int]struct{}),
		sum:  make(map[int]struct{}),
		diff: make(map[int]struct{}),
	}

	var cols []int
	b.solveNQueensUtilUseSet(n, 0, cols)

	var result [][]string
	for _, s := range b.solutions {
		var solution []string
		for _, col := range s {
			solution = append(solution, String(n, col))
		}
		result = append(result, solution)
	}

	return result
}

func (b *board) solveNQueensUtilUseSet(n, row int, cols []int) {
	if row >= n {
		b.solutions = append(b.solutions, cols)
		return
	}

	for col := 0; col < n; col++ {
		if b.IsSafe(row, col) {
			// 添加皇后
			b.AddQueen(row, col)

			// 拷贝记录已添加皇后的列的数组
			cols = append(cols, col)
			var copyCols = make([]int, len(cols))
			copy(copyCols, cols)

			// 搜索下一行
			b.solveNQueensUtilUseSet(n, row+1, copyCols)

			// 回溯, 注意别忘了从 cols 移除当前行的皇后
			cols = cols[0 : len(cols)-1]
			b.RemoveQueen(row, col)
		}
	}
}
