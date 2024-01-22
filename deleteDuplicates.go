package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	helper := &ListNode{Val: 0, Next: head}
	cur := helper
	for cur != nil {
		p1 := cur.Next
		var p2 *ListNode = nil
		if p1 != nil {
			p2 = p1.Next
		}
		if p2 == nil {
			break
		}

		if p1.Val != p2.Val {
			cur = cur.Next
		} else {
			for p1.Val == p2.Val {
				p2 = p2.Next
				if p2 == nil {
					break
				}
			}
			cur.Next = p2
		}
	}
	return helper.Next
}
