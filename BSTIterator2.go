package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type BSTIterator2 struct {
	cur   *TreeNode
	stack []*TreeNode
}

func Constructor2(root *TreeNode) BSTIterator2 {
	return BSTIterator2{
		cur:   root,
		stack: make([]*TreeNode, 0),
	}
}

func (this *BSTIterator2) pushLeftBranch() {
	for ; this.cur != nil; this.cur = this.cur.Left {
		this.stack = append(this.stack, this.cur)
	}
}

func (this *BSTIterator2) Next() int {
	this.pushLeftBranch()
	this.cur, this.stack = this.stack[len(this.stack)-1], this.stack[:len(this.stack)-1]
	val := this.cur.Val
	this.cur = this.cur.Right
	return val
}

func (this *BSTIterator2) HasNext() bool {
	return this.cur != nil || len(this.stack) > 0
}

func main07104() {
	root := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	bSTIterator := Constructor2(root)
	fmt.Println(bSTIterator.Next())    // 返回 3
	fmt.Println(bSTIterator.Next())    // 返回 7
	fmt.Println(bSTIterator.HasNext()) // 返回 True
	fmt.Println(bSTIterator.Next())    // 返回 9
	fmt.Println(bSTIterator.HasNext()) // 返回 True
	fmt.Println(bSTIterator.Next())    // 返回 15
	fmt.Println(bSTIterator.HasNext()) // 返回 True
	fmt.Println(bSTIterator.Next())    // 返回 20
	fmt.Println(bSTIterator.HasNext()) // 返回 False
}
