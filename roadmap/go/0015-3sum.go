package roadmap

import "sort"

// Leetcode Problem: https://leetcode.com/problems/3sum/
// 难点: 排除重复的组合
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 2 {
		return nil
	}

	sort.Ints(nums)

	var res [][]int
	for iNum1 := 0; iNum1 < n-2; iNum1++ {
		// 第一个数字和前面数字重复, 可以跳过, 否则会产生重复的组合
		// 例如: [-1,0,1,2,-1,-4] 排序后是 [-4, -1, -1, 0, 1, 2]
		if iNum1 > 0 && nums[iNum1] == nums[iNum1-1] {
			continue
		}

		target := 0 - nums[iNum1]

		iNum2 := iNum1 + 1
		iNum3 := n - 1
		for iNum2 < iNum3 {
			sum := nums[iNum2] + nums[iNum3]
			if sum == target {
				res = append(res, []int{nums[iNum1], nums[iNum2], nums[iNum3]})

				// 例如: [-2,0,0,2,2]
				iNum3--
				for iNum3 > iNum2 && nums[iNum3] == nums[iNum3+1] {
					iNum3--
				}
				// 例如: [-2,0,0,1,1,1,2]
				iNum2++
				for iNum2 < iNum3 && nums[iNum2] == nums[iNum2-1] {
					iNum2++
				}
			} else if sum < target {
				iNum2++
			} else {
				iNum3--
			}
		}
	}

	return res
}
