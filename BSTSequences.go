package main

//  https://leetcode.cn/problems/bst-sequences-lcci/?favorite=xb9lfcwi

func BSTSequences(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	if root.Left == nil && root.Right == nil {
		ret = append(ret, []int{root.Val})
		return ret
	}
	leftRet := BSTSequences(root.Left)
	rightRet := BSTSequences(root.Right)
	for _, lt := range leftRet {
		for _, rt := range rightRet {
			line := []int{root.Val}
			line = append(line, lt...)
			line = append(line, rt...)
			ret = append(ret, line)
		}
	}
	for _, rt := range rightRet {
		for _, lt := range leftRet {
			line := []int{root.Val}
			line = append(line, rt...)
			line = append(line, lt...)
			ret = append(ret, line)
		}
	}
	return ret
}
