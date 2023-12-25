package algorithms

import "sort"

// 字节跳动第一面面试题, 要求算法时间复杂度少于 O(n^2), 空间复杂度是 O(1)
// TwoSum3 返回 a 中两数之和等于 target 的所有组合
// a 中没有重复的数字
func TwoSum3(a []int, target int) [][]int {
	sort.Ints(a)

	var result [][]int

	i, j := 0, len(a)-1
	for i < j {
		got := a[i] + a[j]

		if got == target {
			result = append(result, []int{a[i], a[j]})
			// 注意这里要缩小范围
			i++
			j--
		} else if got < target {
			i++
		} else {
			j--
		}
	}

	return result
}
