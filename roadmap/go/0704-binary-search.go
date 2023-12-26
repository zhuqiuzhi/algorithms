package roadmap

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	// 循环不变式: a[lo, hi] 包含 target
	for lo <= hi {
		// mid := lo + (hi-lo)/2
		mid := lo + (hi-lo)>>1
		if target < nums[mid] {
			hi = mid - 1
		} else if target > nums[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
