package main

// https://fleetcode.cn/leetbook/read/top-interview-questions/xa0a97/
// 奇偶链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	otherHead := &ListNode{}
	p1, p2 := head, otherHead
	for p1 != nil {
		p1Next := p1.Next
		p2.Next = p1Next
		p2 = p2.Next
		if p1Next == nil {
			break
		}
		p1NextNext := p1Next.Next
		p1.Next = p1NextNext
		p1 = p1.Next
		p2.Next = nil

	}
	for p1 = head; p1.Next != nil; p1 = p1.Next {
	}
	p1.Next = otherHead.Next

	return head
}
