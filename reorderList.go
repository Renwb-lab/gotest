package main

// https://leetcode.cn/problems/reorder-list/

// 输入：head = [1,2,3,4]
// 输出：[1,4,2,3]

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

func reorderList(head *ListNode) {
	var dp func(inp *ListNode, in int)
	outp := head
	out := 0
	dp = func(inp *ListNode, in int) {
		if inp == nil {
			return
		}
		dp(inp.Next, in+1)
		if in <= out {
			return
		}
		p1 := outp.Next
		outp.Next = inp
		outp.Next.Next = p1
		outp = p1

		out += 1
	}
	dp(head, 0)
	if outp != nil {
		outp.Next = nil
	}
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	reorderList(head)
}
