package algorithms

// TwoSum return indices of the two numbers such that they add up to a specific target.
// https://leetcode-cn.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	v := make(map[int]int)
	for i, x := range nums {
		y := target - x
		if j, ok := v[y]; ok {
			return []int{j, i}
		}
		v[x] = i
	}
	return nil
}
