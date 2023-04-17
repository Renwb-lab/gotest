package main

// https://leetcode.cn/leetbook/read/top-interview-questions/xa262d/
// 排序链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var preSlow *ListNode = nil
	slow, fast := head, head
	for fast != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	newHead := preSlow.Next
	preSlow.Next = nil
	oneList := sortList(head)
	otherList := sortList(newHead)

	fakeNode := ListNode{}
	p1, p2, p3 := oneList, otherList, &fakeNode
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p3.Next = p1
			p1 = p1.Next
		} else {
			p3.Next = p2
			p2 = p2.Next
		}
		p3 = p3.Next
		p3.Next = nil
	}
	if p1 != nil {
		p3.Next = p1
	}
	if p2 != nil {
		p3.Next = p2
	}
	ret := fakeNode.Next
	fakeNode.Next = nil
	return ret
}

func main() {
	head := &ListNode{Val: -1}
	head.Next = &ListNode{Val: 5}
	head.Next.Next = &ListNode{Val: 3}
	// head.Next.Next.Next = &ListNode{Val: 4}
	// head.Next.Next.Next.Next = &ListNode{Val: 0}
	sortList(head)
}
