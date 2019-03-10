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
		// 循环不变式： small 是 A[i, k-1] 中最小元素的下标, k 是到达表达式 j < n 时 j 的值, k >= i+1
		// 归纳依据: j 是 i + 1 时, 此时 small 是 i，所以 small 是 [i, i] 中最小元素的下标
		// 归纳假设: small 已经是 A[i..k-1]中最小元素的下标
		for j = i + 1; j < n; j++ {
			if data.Less(j, small) {
				small = j
			}
		}
		data.Swap(i, small)
	}
}

// RecSelectionSort 是递归版本的选择排序
// 假设要排序的数据是在数组A[0..n-1]中。
// (1) 从数组A的尾部，也就是从A[i..n-1]中，选出最小的元素。
// (2) 将步骤(1)中选出的元素与A[i]互换。
// (3) 将剩下的数组A[i+1..n-1]进行排序。
func RecSelectionSort(data sort.Interface) {
	recSelectionSort(data, 0)
}

func recSelectionSort(data sort.Interface, i int) {
	n := data.Len()
	if i < n {
		var small = i
		for j := i + 1; j < n; j++ {
			if data.Less(j, small) {
				small = j
			}
		}
		data.Swap(i, small)
		recSelectionSort(data, i+1)
	}
}

type student struct {
	ID    int
	Name  string
	Grade string
}

type students []*student

func (s students) Len() int { return len(s) }

func (s students) Less(i, j int) bool { return s[i].ID < s[j].ID }

func (s students) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func SortStudents(s students) {
	SelectionSort(s)
}
