package algorithms

import "sort"

// SelectionSort 通过选择排序算法排序 data
// 选择排序是每次从 a[i, n-1] 找出最小的数的下标 small，然后将 a[small] 和 a[i] 交换
// i 从 0 开始，直到 n-2
func SelectionSort(data sort.Interface) {
	n := data.Len()
	var i, j, small int
	for i = 0; i < n-1; i++ {
		small = i
		for j = i + 1; j < n; j++ {
			if data.Less(j, small) {
				small = j
			}
		}
		data.Swap(i, small)
	}
}
