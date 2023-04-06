package main

//  https://leetcode.cn/problems/bst-sequences-lcci/?favorite=xb9lfcwi

func BSTSequences(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{{}}
	}
	ret := [][]int{}
	queue := []*TreeNode{root}
	line := make([]int, 0)

	var dp func(queue []*TreeNode)
	dp = func(queue []*TreeNode) {
		if len(queue) == 0 {
			// 需要拷贝下，非常重要
			ret = append(ret, append([]int{}, line...))
			return
		}
		size := len(queue)
		for i := 0; i < size; i += 1 {
			// 这个思路好，计算出来长度，每次取第一个，省的拼接
			p := queue[0]
			queue = queue[1:]
			line = append(line, p.Val)

			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
			dp(queue)
			line = line[:len(line)-1]
			queue = queue[0 : size-1] // 重复使用切片，多次使用。否则会出现超出最大内存的情况
			queue = append(queue, p)
		}
	}
	dp(queue)
	return ret
}
