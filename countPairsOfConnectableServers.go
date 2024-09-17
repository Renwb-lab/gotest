package main

import "fmt"

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
	}

	var cnt int
	var dfs func(int, int, int)
	// fa <- x 中的满足【被 signalSpeed 整除】的数量
	dfs = func(x, fa, sum int) {
		if sum%signalSpeed == 0 {
			cnt += 1
		}
		for _, e := range g[x] {
			if e.to != fa {
				dfs(e.to, x, sum+e.wt)
			}
		}
	}

	ans := make([]int, n)
	for i, gi := range g {
		sum := 0
		for _, e := range gi {
			cnt = 0
			// 计算 i 节点 从 to 分支上满足【被 signalSpeed 整除】的数量
			dfs(e.to, i, e.wt) // (to, i) 的数据
			ans[i] += cnt * sum
			sum += cnt
		}
	}
	return ans
}

func main071036() {
	fmt.Println(countPairsOfConnectableServers([][]int{
		{0, 1, 1},
		{1, 2, 5},
		{2, 3, 13},
		{3, 4, 9},
		{4, 5, 2},
	}, 1))
}
