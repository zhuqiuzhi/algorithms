package roadmap

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// sortList 使用快速排序算法将链表排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	return sortListRange(head, getTail(head))
}

func getTail(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var tail = head
	for tail.Next != nil {
		tail = tail.Next
	}
	return tail
}

func sortListRange(head, tail *ListNode) *ListNode {
	// base case
	if head == nil || head == tail {
		return head
	}

	var smallHead, smallTail, largeHead, largeTail *ListNode
	// use tail as pivot node
	for head != tail {
		if head.Val <= tail.Val { // add head to small linked list
			if smallHead == nil {
				smallHead = head
				smallTail = head
			} else {
				smallTail.Next = head
				smallTail = head // 易错点: 忘记更新 smallTail 指针
			}
		} else {
			if largeHead == nil { // add head into large linked list
				largeHead = head
				largeTail = head
			} else {
				largeTail.Next = head
				largeTail = head
			}
		}
		head = head.Next
	}

	var newHead *ListNode
	// small linked list is not empty
	if smallTail != nil {
		smallTail.Next = nil
		newHead = sortListRange(smallHead, smallTail)
		getTail(newHead).Next = tail
	}

	// large linked list is not empty
	if largeTail != nil {
		largeTail.Next = nil
		tail.Next = sortListRange(largeHead, largeTail)
	}

	if newHead == nil {
		return tail
	}

	return newHead
}

// mergeSortLlist 使用归并排序算法将链表排序
// 难点: 将链表分成两个单独的链表: 找到中间的结点, 然后将结点的 Next 指针设置为 nil
func mergeSortList(head *ListNode) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}

	// div into two linked list
	left := head
	mid := getMid(head)
	right := mid.Next
	mid.Next = nil

	left = mergeSortList(left)
	right = mergeSortList(right)
	return mergeTwoLists(left, right)
}

// 重要技巧: 快慢指针, 当快指针到达表尾时, 慢指针指向的就是中间结点
// getMid 返回两个结点和两个结点以上的链表的中间结点
// 1-2
// 1->2->3
// 1->2->3->4
func getMid(head *ListNode) *ListNode {
	// error: slow, fast := head, head
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
