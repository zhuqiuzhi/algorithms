package roadmap

// Leetcode Problem: https://leetcode.com/problems/number-of-ways-to-reach-a-position-after-exactly-k-steps/
// Since the answer may be very large, return it modulo 109 + 7.
// Constraints:
// 1 <= startPos, endPos, k <= 1000
var combinations [1000][]int32

func init() {
	// tabulation
	// 使用组合学来解决问题
	for i := 0; i < 1000; i++ {
		combinations[i] = make([]int32, i+1)
		combinations[i][0] = int32(i + 1)

		for k := 1; k < i; k++ {
			combinations[i][k] = (combinations[i-1][k-1] + combinations[i-1][k]) % 1000000007
		}

		combinations[i][i] = 1
	}
}

func numberOfWays(startPos int, endPos int, k int) int {
	diff := endPos - startPos
	if diff < 0 {
		diff = -diff
	}
	// 当距离大于 k 时, 或者 k-diff 是奇数, 因为先通过 diff 步数走到 endPos, 然后必须通过偶数步才能回到 endPos
	if k < diff || (k-diff)%2 == 1 {
		return 0
	}

	return int(combinations[k-1][(diff+k)>>1-1])
}

// 链接：https://leetcode.cn/problems/number-of-ways-to-reach-a-position-after-exactly-k-steps/solutions/2595035/er-wei-dong-tai-gui-zhi-yi-wei-yan-bian-cs3d0/
// 定义状态dp[i][j] 为移动 i 步从任意位置达到位置 j 的不同方法数目
// dp[i][j] = dp[i-1][j-1] + dp[i-1][j+1]
// 左边界递推公式是：dp[i][0] = dp[i-1][1]；使用 i 步移动到 0 的方法数等于使用 i-1 步移动到 1的方法数(因为再往左一步就能走到 0)
// 右边界递推公式是：dp[i][2k] = dp[i-1][2k-1]
func numberOfWaysDp(startPos int, endPos int, k int) int {
	diff := endPos - startPos
	if diff < 0 {
		diff = -diff
	}
	// 当距离大于 k 时, 或者 k-diff 是奇数, 因为先通过 diff 步数走到 endPos, 然后必须通过偶数步才能回到 endPos
	if k < diff || (k-diff)%2 == 1 {
		return 0
	}

	var dp [][]int = make([][]int, k+1)
	dp[0] = make([]int, 2*k+1)
	dp[0][k] = 1 // 其余都为 0

	for i := 1; i <= k; i++ {
		dp[i] = make([]int, 2*k+1)
		dp[i][0] = dp[i-1][1]
		for j := 1; j < 2*k; j++ {
			dp[i][j] = (dp[i-1][j-1] + dp[i-1][j+1]) % 1000000007
		}
		dp[i][2*k] = dp[i-1][2*k-1]
	}

	return dp[k][endPos+k-startPos]
}

// 使用深度遍历暴力破解, 某些情况下, 时间太久
func numberOfWaysDfs(startPos int, endPos int, k int) int {
	var ret int

	var dfs func(startPos, endPos, i int)
	dfs = func(startPos, endPos, i int) {
		if startPos == endPos && i == k {
			ret++
			return
		}
		if i > k {
			return
		}
		// left
		dfs(startPos-1, endPos, i+1)
		// right
		dfs(startPos+1, endPos, i+1)
	}
	dfs(startPos, endPos, 0)

	return ret
}
