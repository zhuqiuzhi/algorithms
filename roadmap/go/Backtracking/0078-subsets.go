package backtracking

// Leetcode Problem: https://leetcode.com/problems/subsets/
// Given an integer array nums of unique elements, return all possible
// subsets (the power set).
// In mathematics, the power set of a set S is the set of all subsets of S,
// including the empty set and S itself.

// The solution set must not contain duplicate subsets. Return the solution in any order.
// 限制: 不能重复, 和排列有区别
// 技巧: 对 n 个数字依次做 0 和 1 选择
// 可以作为回溯算法的模板
func subsets(nums []int) [][]int {
	var result = make([][]int, 0, 2^len(nums))

	var dfs func(i int)
	var subset []int
	// i 是已进行选择的元素个数
	dfs = func(i int) {
		// base case
		if i >= len(nums) {
			newSubset := make([]int, len(subset))
			copy(newSubset, subset)
			result = append(result, newSubset)
			return
		}

		// 选择添加 num[i]
		subset = append(subset, nums[i])
		dfs(i + 1)

		// 选择不添加 num[i]
		// pop
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}

	dfs(0)
	return result
}
