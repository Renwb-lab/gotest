package main

// https://leetcode.cn/problems/minimum-height-trees/
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	parents := make([]int, n)
	bfs := func(start int) (x int) {
		visit := make([]bool, n)
		visit[start] = true
		queue := []int{start}
		for len(queue) > 0 {
			x = queue[0]
			queue = queue[1:]
			for _, y := range g[x] {
				if !visit[y] {
					visit[y] = true
					queue = append(queue, y)
					// y的父节点为x, 后续需要从y向上遍历
					parents[y] = x
				}
			}
		}
		return
	}
	x := bfs(0) // 找到与节点 0 最远的节点 x
	y := bfs(x) // 找到与节点 x 最远的节点 y

	// 从y出发找最长的路径 path
	path := []int{}
	parents[x] = -1
	for y != -1 {
		path = append(path, y)
		y = parents[y]
	}
	m := len(path)
	if m%2 == 0 {
		return []int{path[m/2-1], path[m/2]}
	}
	return []int{path[m/2]}
}
