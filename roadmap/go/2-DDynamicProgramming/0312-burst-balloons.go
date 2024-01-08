package roadmap

// Leetcode Problem: https://leetcode.com/problems/burst-balloons/
func maxCoins(nums []int) int {
	l := len(nums) + 2

	fullNums := make([]int, l)
	fullNums[0], fullNums[l-1] = 1, 1

	for i := range nums {
		fullNums[i+1] = nums[i]
	}

	store := make([][]int, l)
	for i := 1; i < l; i++ {
		store[i] = make([]int, l)
		for j := i; j < l; j++ {
			store[i][j] = -1
		}
	}

	return getStore(1, l-2, store, fullNums)
}

// getStore 获取 [begin, end] 区间最大得分
// 矩阵 store 用于记录 [begin,end] 区间的最大得分 , 值为 -1 时，表示还未计算出该区间的最大得分
func getStore(begin, end int, store [][]int, nums []int) int {
	if begin > end {
		return 0
	}
	if store[begin][end] != -1 {
		return store[begin][end]
	}

	for pos := begin; pos <= end; pos++ {
		left := getStore(begin, pos-1, store, nums)

		// 注意 middle 不是 nums[pos-1] * nums[pos] * nums[pos+1]
		middle := nums[begin-1] * nums[pos] * nums[end+1]

		right := getStore(pos+1, end, store, nums)

		coins := left + middle + right

		if coins > store[begin][end] {
			store[begin][end] = coins
		}
	}
	return store[begin][end]
}
