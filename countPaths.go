package main

import "fmt"

const mx = 1e5 + 1

var np = [mx]bool{1: true} // true: 非质数, false: 质数

func init() {
	for i := 2; i*i < mx; i += 1 {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func countPaths(n int, edges [][]int) int64 {
	g := make([][]int, n+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	size := make([]int, n+1)
	var nodes []int
	var dfs func(int, int)
	dfs = func(x, fa int) {
		nodes = append(nodes, x)
		for _, y := range g[x] {
			if y != fa && np[y] {
				dfs(y, x)
			}
		}
	}
	res := int64(0)
	for x := 1; x <= n; x += 1 {
		if np[x] { // 跳过非质数
			continue
		}
		sum := 0
		for _, y := range g[x] { // 质数x把这棵树分成了若干个连通块
			if !np[y] { // 如果y为质数，则跳过
				continue
			}
			if size[y] == 0 { // 尚未计算过
				nodes = []int{} // 记录从y出发，能够找到的非质数，包括y
				dfs(y, -1)
				// fmt.Printf("%d, %d, %+#v\n", x, y, nodes)
				for _, z := range nodes {
					size[z] = len(nodes) // 这些非质数都是相互连接的，所以从任意一个点出发，他们的数量都是整体集合的数量
				}
			}
			// 这 size[y] 个非质数与之前遍历到的 sum 个非质数，两两之间的路径只包含质数 x
			res += int64(size[y]) * int64(sum)
			sum += size[y]
		}
		res += int64(sum) // 从 x 出发的路径
	}
	return res
}

func main() {
	fmt.Printf("%d\n", countPaths(5, [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}))
}
