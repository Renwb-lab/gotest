package main

import "fmt"

func maximumPoints2(edges [][]int, coins []int, k int) int {
	l := len(coins)
	g := make([][]int, l)
	for i := 0; i < len(edges); i += 1 {
		x, y := edges[i][0], edges[i][1]
		if len(g[x]) == 0 {
			g[x] = make([]int, 0)
		}
		g[x] = append(g[x], y)

		if len(g[y]) == 0 {
			g[y] = make([]int, 0)
		}
		g[y] = append(g[y], x)
	}

	var dfs func(i, j, fa int) int
	m := make(map[string]int)
	dfs = func(i, j, fa int) int {
		key := fmt.Sprintf("%d_%d_%d", i, j, fa)
		if val, ok := m[key]; ok {
			// fmt.Printf("%s: %d\n", key, val)
			return val
		}
		res1 := (coins[i] >> j) - k
		res2 := coins[i] >> (j + 1)
		for _, c := range g[i] {
			if c != fa {
				res1 += dfs(c, j, i)
				if j < 13 {
					res2 += dfs(c, j+1, i)
				}
			}
		}
		res := res2
		if res1 > res2 {
			res = res1
		}
		m[key] = res
		return res
	}
	return dfs(0, 0, -1)
}

func main159() {
	fmt.Println(maximumPoints2([][]int{
		{0, 1},
		{1, 2},
		{2, 3},
	}, []int{10, 10, 3, 3}, 5))
}
