package roadmap

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(a []int) *ListNode {
	var head, tail *ListNode
	for i, value := range a {
		if i == 0 {
			head = &ListNode{
				Val: value,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: value,
			}
			tail = tail.Next
		}
	}

	return head
}

func Values(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}
