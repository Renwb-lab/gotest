package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.cn/problems/check-subtree-lcci/?favorite=xb9lfcwi
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}
	if t1 == nil {
		return false
	}
	if checkSubTree(t1.Left, t2) {
		return true
	}
	if checkSubTree(t1.Right, t2) {
		return true
	}

	return theSameTree(t1, t2)
}

func theSameTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil {
		return false
	}
	if t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return theSameTree(t1.Left, t2.Left) && theSameTree(t1.Right, t2.Right)
}
