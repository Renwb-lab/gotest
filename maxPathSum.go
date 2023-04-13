package main

import "math"

// https://leetcode.cn/leetbook/read/top-interview-questions/x2hnpi/
// 二叉树中的最大路径和

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxfunc := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	result := math.MinInt32 // 最终的数据肯定是来源于某一个节点
	var dfs func(root *TreeNode) int
	// 包含当前节点的一边路径的长度
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftValue := dfs(root.Left)
		rightValue := dfs(root.Right)
		// 包含当前节点的一边路径的长度
		thisResult := root.Val + maxfunc(maxfunc(0, leftValue), maxfunc(0, rightValue))
		// 包含当前节点的两边路径的长度
		sumResult := root.Val + maxfunc(0, leftValue) + maxfunc(0, rightValue)
		if sumResult > result {
			result = sumResult
		}
		return thisResult
	}
	_ = dfs(root)
	return result
}
