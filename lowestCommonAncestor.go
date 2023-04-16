package main

// https://leetcode.cn/leetbook/read/top-interview-questions/xas5th/
// 二叉树的最近公共祖先

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 从下向上，判断当前节点是否包含两个节点中的一个节点，如果包含的话，则返回当前节点
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 如果左右孩子分别包含一个节点的话，则说明根节点是最低的公共祖先
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
