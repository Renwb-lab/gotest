package main

import "math"

const inf = math.MaxInt / 2

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type Graph [][]int

func Constructor13(n int, edges [][]int) Graph {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, e := range edges {
		g[e[0]][e[1]] = e[2]
	}
	return g
}

func (g Graph) AddEdge(e []int) {
	g[e[0]][e[1]] = e[2]
}

func (g Graph) ShortestPath(start, end int) int {
	n := len(g)
	dis := make([]int, n) // 从 start 出发，到各个点的最短路，如果不存在则为无穷大
	vis := make([]bool, n)
	for i := range dis {
		dis[i], vis[i] = inf, false
	}
	dis[start] = 0
	for {
		// 查找距离当前节点中最近的节点x
		x := -1
		for i, b := range vis {
			if !b && (x != -1 && dis[i] < dis[x]) {
				x = i
			}
		}
		// 如果所有的节点都已经更新
		if x < 0 || dis[x] == inf {
			return -1
		}
		// 如果最近的节点是end的话，则期间推出
		if x == end {
			return dis[x]
		}
		// 使用最近的节点更新最近的节点距离
		vis[x] = true
		for y, w := range g[x] {
			dis[y] = min(dis[y], dis[x]+w)
		}
	}
}
