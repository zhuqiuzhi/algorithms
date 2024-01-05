package roadmap

// Leetcode Problem: https://leetcode.com/problems/trapping-rain-water/
// 难点1: 每个位置能存储的水量 = min(max(left), max(right))- height(i) 且如果是负数是 0
// 难点2: 怎么快速得到 max(left), max(right):
// 线性时间: 可以用空间换时间, 从做到右遍历得到 maxLeft 数组, 从右到左遍历得到 maxRight 数组
// 常量时间: 用双指针,
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	var ret int
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]

	for left < right {
		if leftMax < rightMax {
			// 虽然此时 reightMax 不一定是右边全部的最大值, 但是如果 leftMax 比右边的数小,
			// 则肯定是 leftMax 和 右边最大值的较小值
			left++
			// 此时 leftMax 是 [0...left) 中的最大值
			// if leftMax > height[left] {
			// 	ret += leftMax - height[left]
			// }
			leftMax = max(leftMax, height[left])
			// 此时 leftMax 是 [0...left] 中的最大值
			// 因为如果是负数, 则是 0, 所以可以先计算 leftMax 和 height 的较大值, 然后减去 height[left]
			ret += leftMax - height[left]
		} else {
			// rightMax <= leftMax
			right--
			rightMax = max(rightMax, height[right])
			ret += rightMax - height[right]
		}
	}

	return ret
}
