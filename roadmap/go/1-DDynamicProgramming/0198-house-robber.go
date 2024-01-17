package roadmap

// Leetcode Problem: https://leetcode.com/problems/house-robber/
// 难点: 当抢第 i 个房子时, 需要计算此时的最大价值
func rob(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	var dp = make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 2; i < len(dp); i++ {
		// dp[i] 是在第 i 个房子(i 从 1 开始)时做出能不触发警报能抢到最大价值
		// 选择不抢, 则经过第 i -1 个房子的最大价值
		// 选择抢, 则是经过第 i-2 个房子抢到的最大价值加上第 i 个房子的价值
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[len(nums)]
}

// 节约空间, 使用 pre 和 prepre 记录 i-1 和 i-2 的数据
// 边界case: nums=[1]
func robO1(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	var ret int
	prepre, pre := 0, 0
	for i := 0; i < len(nums); i++ {
		// ret 是在第 i 个房子(i 从 0 开始)时做出能不触发警报能抢到最大价值
		// 选择不抢, 则经过第 i -1 个房子的最大价值
		// 选择抢, 则是经过第 i-2 个房子抢到的最大价值加上第 i 个房子的价值
		ret = max(pre, prepre+nums[i])
		// update
		prepre = pre
		pre = ret
	}

	return ret
}
