package main

import "fmt"

// 100041. 可以到达每一个节点的最少边反转次数
// https://leetcode.cn/problems/minimum-edge-reversals-so-every-node-is-reachable/

func minEdgeReversals(n int, edges [][]int) []int {
	type pair struct{ to, dir int }
	g := make([][]pair, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], pair{v, 1})
		g[v] = append(g[v], pair{u, -1})
	}

	ans := make([]int, n)
	var dfs func(int, int)
	dfs = func(u, fu int) {
		for _, e := range g[u] {
			v, d := e.to, e.dir
			if v != fu {
				if d < 0 {
					ans[0] += 1
				}
				dfs(v, u)
			}
		}
	}
	dfs(0, -1)

	var reroot func(int, int)
	reroot = func(u, fu int) {
		for _, e := range g[u] {
			v, d := e.to, e.dir
			if v != fu {
				ans[v] = ans[u] + d
				reroot(v, u)
			}
		}
	}
	reroot(0, -1)

	return ans
}

func main135() {
	fmt.Printf("%+#v\n",
		minEdgeReversals(4, [][]int{
			{2, 0},
			{2, 1},
			{1, 3},
		},
		))
}
