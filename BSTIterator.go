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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	root *TreeNode
	c1   chan *TreeNode
	c2   []int
}

func Constructor7(root *TreeNode) BSTIterator {
	c := make(chan *TreeNode)
	bst := BSTIterator{
		root: root,
		c1:   c,
		c2:   make([]int, 0),
	}
	go func(root *TreeNode) {
		defer func() {
			close(bst.c1)
		}()
		var dp func(root *TreeNode)
		dp = func(root *TreeNode) {
			if root == nil {
				return
			}
			if root.Left != nil {
				dp(root.Left)
			}
			bst.c1 <- root
			if root.Right != nil {
				dp(root.Right)
			}
		}
		fmt.Println("begin")
		dp(root)
		fmt.Println("end")
	}(root)
	return bst
}

func (this *BSTIterator) Next() int {
	if len(this.c2) == 0 {
		this.HasNext()
	}
	t := this.c2[0]
	this.c2 = this.c2[1:]
	return t
}

func (this *BSTIterator) HasNext() bool {
	if len(this.c2) > 0 {
		return true
	}
	t, ok := <-this.c1
	if ok {
		this.c2 = append(this.c2, t.Val)
	}
	return ok
}

func main07101() {
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
	bSTIterator := Constructor7(root)
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
