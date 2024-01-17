package roadmap

// Leetcode Problem: https://leetcode.com/problems/arithmetic-slices/
// 难点: 以 nums[i] 结尾的等差数组是 num[i-1] 结尾的等差数组都增加 num[i], 另外再加上最后两个数和 num[i]构成的等差数组
// 也就得到递推公式: dp[i] = dp[i-1]+1, 但最后需要求和
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	var sum int
	var dp = make([]int, len(nums))
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
			sum += dp[i]
		}
	}
	return sum
}
