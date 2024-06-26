package main

import "fmt"

// 834. 树中距离之和
// https://leetcode.cn/problems/sum-of-distances-in-tree/description/

func sumOfDistancesInTree(n int, edges [][]int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	size := make([]int, n) // 每个节点作为根节点，包含的子节点数量
	ans := make([]int, n)  // 结果
	var dfs func(int, int, int)
	dfs = func(u, fu, h int) {
		ans[0] += h

		size[u] = 1
		for _, v := range g[u] {
			if v != fu {
				dfs(v, u, h+1)
				size[u] += size[v]
			}
		}
	}
	dfs(0, -1, 0)

	var reroot func(int, int)
	reroot = func(u, fu int) {
		for _, v := range g[u] {
			if v != fu {
				// u -> v
				// 距离v中的节点变近了（ - size[v]），不在v中的节点变远了(+ n - size[v])
				ans[v] += ans[u] + n - 2*size[v]
				reroot(v, u)
			}
		}
	}
	reroot(0, -1)
	return ans
}

// main135
func main136() {
	fmt.Printf("%+#v\n", sumOfDistancesInTree(6, [][]int{
		{0, 1},
		{0, 2},
		{2, 3},
		{2, 4},
		{2, 5},
	}))
}
