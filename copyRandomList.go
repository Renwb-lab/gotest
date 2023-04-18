package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xam1wr/
// 复制带随机指针的链表

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 构建next信息
	for p1 := head; p1 != nil; {
		p1Next := p1.Next
		p1.Next = &Node{Val: p1.Val, Next: p1Next, Random: nil}
		p1 = p1Next
	}
	// 构建random信息
	for p1 := head; p1 != nil; {
		p1Next := p1.Next
		if p1.Random != nil {
			p1Next.Random = p1.Random.Next
		}
		p1 = p1Next.Next
	}
	// 将链表拆分为两条
	fakeNode := &Node{}
	p2 := fakeNode
	for p1 := head; p1 != nil; {
		p1Next := p1.Next
		p1.Next = p1Next.Next
		p1 = p1.Next

		p2.Next = p1Next
		p2 = p2.Next
	}
	ret := fakeNode.Next
	fakeNode.Next = nil
	return ret
}

func main82() {
	n7 := &Node{Val: 7}
	n13 := &Node{Val: 13}
	n11 := &Node{Val: 11}
	n10 := &Node{Val: 10}
	n1 := &Node{Val: 1}
	n7.Next = n13
	n13.Next = n11
	n11.Next = n10
	n10.Next = n1

	n7.Random = nil
	n13.Random = n7
	n11.Random = n1
	n10.Random = n11
	n1.Random = n7

	ret := copyRandomList(n7)
	for p1 := ret; p1 != nil; p1 = p1.Next {
		fmt.Printf("[%d, %v]\n", p1.Val, p1.Random)
	}
	for p1 := n7; p1 != nil; p1 = p1.Next {
		fmt.Printf("[%d, %v]\n", p1.Val, p1.Random)
	}
}
