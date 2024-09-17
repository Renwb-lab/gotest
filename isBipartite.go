package main

import "fmt"

func isBipartite(graph [][]int) bool {
	// 遍历一遍
	// a 输入A
	// 那么他的相邻边 输入B
	// 如果出现在A中，则返回失败
	// 否则就加入到B中
	n := len(graph)
	partA, partB := make([]int, n, n), make([]int, n, n)
	// 判断当前节点是否满足要求
	feasible := func(idx int, flag int) bool {
		if flag == 1 {
			if partB[idx] == 1 {
				return false
			}
			partA[idx] = 1
		} else {
			if partA[idx] == 1 {
				return false
			}
			partB[idx] = 1
		}
		return true
	}
	var dfs func(idx int, flag int) bool
	dfs = func(idx int, flag int) bool {
		// flag: 1 放入到A中
		// flag: 0 放入到B中
		if idx == n {
			return true
		}
		if !feasible(idx, flag) {
			return false
		}
		for _, nex := range graph[idx] {
			// 已经访问过
			if partA[nex] == 1 || partB[nex] == 1 {
				if !feasible(nex, 1-flag) {
					return false
				}
				continue
			}
			// 如果某个next值不符合要求，直接返回错误
			if !dfs(nex, 1-flag) {
				return false
			}
		}
		return true
	}
	for i := 0; i < n; i += 1 {
		// 是否存在单独的点
		if partA[i] == 1 || partB[i] == 1 {
			continue
		}
		if !dfs(i, 1) {
			return false
		}
	}
	return true
}

func main209() {
	fmt.Println(isBipartite([][]int{
		{},
		{2, 4, 6},
		{1, 4, 8, 9},
		{7, 8},
		{1, 2, 8, 9},
		{6, 9},
		{1, 5, 7, 8, 9},
		{3, 6, 9},
		{2, 3, 4, 6, 9},
		{2, 4, 5, 6, 7, 8},
	}))
}
