package roadmap

// Leetcode Problem: https://leetcode.com/problems/climbing-stairs/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	// dp[3] = dp[2] + dp[1]
	// dp[n] = dp[n-1] + dp[n-2]
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
