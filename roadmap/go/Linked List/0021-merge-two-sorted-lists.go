package roadmap

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy ListNode
	// 重要技巧, 可以避免需要判断 tail 是否是 nil
	tail := &dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil {
		tail.Next = list2
	}

	return dummy.Next
}
