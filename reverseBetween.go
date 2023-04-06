package main

// https://leetcode.cn/problems/reverse-linked-list-ii/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}

	fakeHead := &ListNode{Next: head}

	getNode := func(idx int) *ListNode {
		p, t := fakeHead, 1
		for p != nil {
			if t == idx {
				break
			}
			p, t = p.Next, t+1
		}
		return p
	}

	p1, p2 := getNode(left), getNode(right+1)
	if p1 == nil {
		return head
	}
	for p1.Next != p2 {
		p1Next, p2Next := p1.Next, p2.Next

		p1.Next = p1.Next.Next
		p2.Next = p1Next
		p1Next.Next = p2Next
	}

	ret := fakeHead.Next
	fakeHead.Next = nil
	return ret
}
