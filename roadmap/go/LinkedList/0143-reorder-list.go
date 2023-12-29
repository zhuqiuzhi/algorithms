package roadmap

// Leetcode Problem: https://leetcode.com/problems/reorder-list/
// 难点: 如何控制间隔放置首尾结点: 将链表分成两半, 然后将后半部分链表反转, 然后再合并
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// find middle
	slow, fast := head, head.Next
	// 当链表只有一个结点时, fast 是 nil
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// second half head
	secondHalf := slow.Next
	// 将中间结点和第二个链表切断
	slow.Next = nil

	// reverse second half linked list
	var prev *ListNode
	for secondHalf != nil {
		secondHalf.Next, secondHalf, prev = prev, secondHalf.Next, secondHalf
	}

	// merge two linked list
	firstHalf := head
	secondHalf = prev
	for secondHalf != nil {
		// 因为后面要修改 firstHalf 和 secondHalf 的指针,所以要先保存起来
		firstTemp, secondTemp := firstHalf.Next, secondHalf.Next

		firstHalf.Next = secondHalf
		secondHalf.Next = firstTemp

		// update firstHalt, secondHalf, 指向链表的下一个结点
		firstHalf = firstTemp
		secondHalf = secondTemp
	}
}
