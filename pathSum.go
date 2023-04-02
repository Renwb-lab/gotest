package main

// https://leetcode.cn/problems/paths-with-sum-lcci/solutions/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	st, ret := make([]*TreeNode, 0), 0
	preMap := map[int]int{
		0: 1,
	}
	visited, preSum := &TreeNode{}, 0
	pushLeftBranch := func(p *TreeNode) {
		for p != nil {
			st = append(st, p)
			preSum += p.Val
			if r, ok := preMap[preSum-sum]; ok {
				ret += r
			}
			if _, ok := preMap[preSum]; ok {
				preMap[preSum] += 1
			} else {
				preMap[preSum] = 1
			}
			p = p.Left
		}
	}

	pushLeftBranch(root)
	for len(st) > 0 {
		p := st[len(st)-1]
		if (p.Left == nil || p.Left == visited) && p.Right != visited {
			// 左孩子遍历结束
			pushLeftBranch(p.Right)
		}
		if p.Right == nil || p.Right == visited {
			// 右孩子遍历结束
			visited = p
			st = st[:len(st)-1]
			preMap[preSum] -= 1
			preSum -= p.Val
		}
	}
	return ret
}
