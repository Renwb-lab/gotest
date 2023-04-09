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
func reorderList(head *ListNode) {
	arr := make([]*ListNode, 0)
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p)
	}
	p, l := head, len(arr)
	for i := 1; i < l/2; i += 1 {
		pn := p.Next
		p.Next = arr[l-1-i]
		p.Next.Next = pn
		p = pn
	}
}
