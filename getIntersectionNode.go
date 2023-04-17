package main

// https://leetcode.cn/leetbook/read/top-interview-questions/xan8ah/
// 相交链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 * 	Val  int
 * 	Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	getLength := func(head *ListNode) int {
		ret := 0
		for p := head; p != nil; p = p.Next {
			ret += 1
		}
		return ret
	}
	n, m := getLength(headA), getLength(headB)
	longList, shortList := headA, headB
	if n < m {
		longList, shortList = shortList, longList
		n, m = m, n
	}
	k := n - m
	p1, p2 := longList, shortList
	for i := 0; i < k; i += 1 {
		p1 = p1.Next
	}
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
