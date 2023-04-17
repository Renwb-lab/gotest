package main

// https://leetcode.cn/leetbook/read/top-interview-questions/xaw0rm/
// 回文链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindromeList(head *ListNode) bool {
	p1 := head
	var dp func(p2 *ListNode) bool
	dp = func(p2 *ListNode) bool {
		if p1 != nil && p2 != nil {
			ret := dp(p2.Next) && (p1.Val == p2.Val)
			p1 = p1.Next
			return ret
		}
		return true
	}
	return dp(head)
}

func isPalindromeList2(head *ListNode) bool {
	fakeHead := &ListNode{Val: -1}
	fakeHead.Next = head
	slow, fast := fakeHead, fakeHead
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	// 0 1 1	slow ->1 fast -> 1
	// 0 1 2 1  slow ->1 fast -> 2
	// 注意这里是slow.Next
	newHead := reverseList(slow.Next)
	slow.Next = nil
	ret := true
	p1, p2 := head, newHead
	// 如果链表长度是奇数的话，则一个链比另外一个长一个节点。
	// 对比时，需要注意
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			ret = false
			break
		}
		p1, p2 = p1.Next, p2.Next
	}
	slow.Next = reverseList(newHead)
	fakeHead.Next = nil
	return ret
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		curNext := cur.Next
		cur.Next = pre
		pre = cur
		cur = curNext
	}
	return pre
}
