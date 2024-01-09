package roadmap

// Leetcode Problem: https://leetcode.com/problems/min-cost-climbing-stairs/
func minCostClimbingStairs(cost []int) int {
	// minCost(0) = 0
	// minCost(1) = 0
	// minCost(2) = min(0+cost[1], 0+cost(0))
	// 递推公式: minCost(n)= min(minCost(n-1)+cost[n-1], minCost(n-2)+cost[n-2])
	var dp = make([]int, len(cost)+1)

	dp[0] = 0
	dp[1] = 0
	for i := 2; i < len(dp); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(dp)-1]
}
