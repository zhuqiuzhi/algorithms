package roadmap

// TwoSum return indices of the two numbers such that they add up to a specific target.
// Approach: One-pass Hash Table
// https://leetcode-cn.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	preMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := preMap[complement]; ok {
			return []int{j, i}
		}
		preMap[num] = i
	}
	return nil
}
