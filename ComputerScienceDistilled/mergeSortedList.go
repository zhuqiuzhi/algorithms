// Merging two sorted lists
package algorithms

// MergeSortedList merge sea and fresh into a sorted list
// Source: Book Computer science distilled: learn the art of solving computational
// Chapter 3.1 Iteration, Merging two sorted lists into a third sorted list:
// FISH REUNION: You’re given a list of saltwater fish and a list of freshwater fish,
// both in alphabetical order.How do you create a list featuring all the fish in alphabetical order?
func MergeSortedList(sea, fresh []string) []string {
	var results = make([]string, 0, len(sea)+len(fresh))
	var fish string
	for len(sea) != 0 || len(fresh) != 0 {
		if len(sea) == 0 {
			results = append(results, fresh...)
			break
		}
		if len(fresh) == 0 {
			results = append(results, sea...)
			break
		}

		if sea[0] < fresh[0] {
			fish = sea[0]
			// pop sea
			sea = sea[1:]
		} else {
			fish = fresh[0]
			// pop fresh
			fresh = fresh[1:]
		}
		results = append(results, fish)
	}
	return results
}

// MergeSortedListInPlace merge two sorted integer arrays nums1 and nums2 into sorted array using
// no auxiliary data structure.
// Source: https://leetcode-cn.com/problems/merge-sorted-array/
func MergeSortedListInPlace(nums1 []int, m int, nums2 []int, n int) {
	// TODO
}
