package roadmap

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		// 先将后面结点的指针保存下来
		next := cur.Next
		cur.Next = prev

		// update prev, cur
		prev = cur
		cur = next
	}

	return prev
}

func recursionReverseList(head *ListNode) *ListNode {
	// base case
	if head == nil {
		return nil
	}

	// 只有一个结点, 直接返回 head
	if head.Next == nil {
		return head
	}

	newHead := recursionReverseList(head.Next)
	// 难点是这里, 通过当前结点的 next 设置剩余链表指向自己
	head.Next.Next = head
	head.Next = nil

	return newHead
}

// reverseListUseStack 使用栈来反转链表
func reverseListUseStack(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var stack []*ListNode
	cur := head
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Next
	}

	// 难点是首尾需要单独处理, 不能在循环中进行
	head = stack[len(stack)-1]
	for i := len(stack) - 1; i > 0; i-- {
		stack[i].Next = stack[i-1]
	}
	stack[0].Next = nil

	return head
}
