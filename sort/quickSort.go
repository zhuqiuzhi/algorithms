package sort

import "sort"

// QuickSort 它将一个数组分成两个子数组，将两部分独立地排序
// 划分的数组满足以下条件:
// 1. 对于某个 j，a[j] 已经排定;
// 2. a[lo] 到 a[j-1] 中的所有元素都不大于 a[j]
// 3. a[j+1] 到 a[hi] 中的所有元素都不小于 a[j]
func QuickSort(data sort.Interface) {
	quickSort(data, 0, data.Len()-1)
}

func quickSort(data sort.Interface, low, high int) {
	if high <= low {
		return
	}
	j := partition(data, low, high)
	quickSort(data, low, j-1)
	quickSort(data, j+1, high)
}

func partition(data sort.Interface, low, high int) int {
	i, j := low+1, high

	for {
		for ; i < high && data.Less(i, low); i++ {
		}
		// data[i] >= data[low] or i = high

		for ; j > low && data.Less(low, j); j-- {
		}
		// data[j] <= data[low] or j = low

		if i >= j {
			break
		}

		data.Swap(i, j)
		i++
		j--
		// data[i] <= data[low]
		// data[j] >= data[low]
	}

	// j <= i => data[j] <= data[low]
	data.Swap(low, j)
	return j
}
