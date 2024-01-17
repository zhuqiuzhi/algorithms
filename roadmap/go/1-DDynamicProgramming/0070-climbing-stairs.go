package roadmap

// Leetcode Problem: https://leetcode.com/problems/climbing-stairs/
// 空间复杂度是 O(n)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 空间复杂度: O(1)
// 节约空间, 使用两个变量 pre 和 prepre 记录 i-1 和 i-2
func climbStairsO1(n int) int {
	if n <= 2 {
		return n
	}

	var ret int
	prepre := 1
	pre := 2
	// 1, 2, 3, 5
	for i := 3; i <= n; i++ {
		ret = pre + prepre
		// update
		prepre = pre
		pre = ret
	}

	return ret
}
