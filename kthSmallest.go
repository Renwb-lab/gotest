package main

// https://leetcode.cn/leetbook/read/top-interview-questions/xazo8d/
// 二叉搜索树中第K小的元素

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var dp func(p *TreeNode) bool
	idx, ret := 0, 0
	dp = func(p *TreeNode) bool {
		if p == nil {
			return false
		}
		if dp(p.Left) {
			return true
		}
		idx += 1
		if idx == k {
			ret = p.Val
			return true
		}
		return dp(p.Right)
	}
	dp(root)
	return ret
}
