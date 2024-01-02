package backtracking

import (
	"fmt"
	"strings"
)

// Leetcode Problem: https://leetcode.com/problems/n-queens/
// 按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

func String(n, col int) string {
	return fmt.Sprintf("%sQ%s", strings.Repeat(".", col), strings.Repeat(".", n-col-1))
}

// 解法 1:
// 比较好的讲解寻找一个解的视频教程：https://www.youtube.com/watch?v=xouin83ebxE,
func solveNQueens(n int) [][]string {
	// 所有的可行解的集合
	var allSolutions [][]int

	var dfs func(row int)
	// 当 row 大于或等于 n 说明找到解的所有的皇后的位置
	// row 是已经处理了几行, queens 是已经确定的皇后放置的位置(行和列)
	var queens = make([]int, 0, n)
	dfs = func(row int) {
		// base case
		if row >= n {
			var copyQueens = make([]int, len(queens))
			copy(copyQueens, queens)
			allSolutions = append(allSolutions, copyQueens)
			return
		}

		// 在 row 行寻找安全的列 col
		for col := 0; col < n; col++ {
			var foundSafe = true
			// 难点: 怎么判断放在(row, col)是否安全: 遍历已经放置的皇后位置, 判断是否在同一列、反对角线、正对角线
			for qRow, q := range queens {
				if col == q || row+col == q+qRow || row-col == qRow-q {
					// 会被已放置的皇后攻击
					foundSafe = false
					break
				}
			}
			if !foundSafe {
				continue
			}

			queens = append(queens, col)
			// 注意此种解法，必须拷贝 slice, 因为 slice 是共用底层数组的
			dfs(row + 1)
			// 回溯, 将已经放置的皇后去掉, 继续尝试下一列
			queens = queens[0 : len(queens)-1]
		}
	}
	dfs(0)

	// 将所有解决方案转换成 .Q.. 输出
	var result [][]string
	for _, s := range allSolutions {
		var solution []string
		for _, q := range s {
			solution = append(solution, String(n, q))
		}
		result = append(result, solution)
	}

	return result
}
