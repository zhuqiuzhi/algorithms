from typing import Optional

# Definition for singly-linked list.


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    # Use Optional[X] for a value that could be None
    # Optional[X] is the same as X | None
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev, cur = None, head

        while cur:
            next = cur.next
            cur.next = prev
            # update prev, cur
            prev = cur
            cur = next

        return prev
