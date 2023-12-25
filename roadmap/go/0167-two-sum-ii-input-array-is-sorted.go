package roadmap

// Leetcode Problem: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		got := numbers[i] + numbers[j]
		if got == target {
			return []int{i + 1, j + 1}
		} else if got < target {
			// 排除 numbers[i], 因为它和一个最大的数相加都比 target 小
			i++
		} else {
			// 排除 numbers[j] , 因为它和一个最小的数相加都比 target 大
			j--
		}
	}
	return nil
}
