package main

// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ret := [][]int{}
	queue, newQueue := []*TreeNode{root}, make([]*TreeNode, 0)
	layer := 1
	getLine := func() []int {
		line := make([]int, 0)
		if layer%2 == 1 {
			for i := 0; i < len(queue); i += 1 {
				line = append(line, queue[i].Val)
			}
		} else {
			for i := len(queue) - 1; i >= 0; i -= 1 {
				line = append(line, queue[i].Val)
			}
		}
		return line
	}
	for len(queue) > 0 {
		for _, p := range queue {
			if p.Left != nil {
				newQueue = append(newQueue, p.Left)
			}
			if p.Right != nil {
				newQueue = append(newQueue, p.Right)
			}
		}
		line := getLine()
		ret = append(ret, line)
		queue = append([]*TreeNode{}, newQueue...)
		newQueue = make([]*TreeNode, 0)
		layer += 1
	}
	return ret
}
