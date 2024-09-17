package main

// https://leetcode.cn/problems/bus-routes/

import (
	"fmt"
	"math"
)

func numBusesToDestination(routes [][]int, source int, target int) int {
	if target == source {
		return 0
	}

	// 公交车线路为节点， 如果存在一个站点有多个公交车都可以达到，则认为两个路线相互连通，距离为1
	// 问题转换为：source节点所在路线和target节点所在路线之间的最短距离
	// 从source到target之间的距离可以转换为source节点到其他节点到最短距离
	// target可能在多个路线上，所以需要求到其他节点到最短距离
	n := len(routes)

	// 构建图
	edges := make([][]bool, n)
	for i := range edges {
		edges[i] = make([]bool, n)
	}
	// 站点到节点的映射关系
	rec := make(map[int][]int)
	for i, route := range routes {
		for _, site := range route {
			rec[site] = append(rec[site], i)
			for _, j := range rec[site] {
				edges[i][j] = true
				edges[j][i] = true
			}
		}
	}

	// dis 存放source到每个节点到最近距离
	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	// 将source所在的车辆存放当前的队列，并标记距离
	q := make([]int, 0)
	for _, r := range rec[source] {
		q = append(q, r)
		dis[r] = 1
	}
	// 从队列中取出车辆，访问从该车辆可以达到的另一车辆
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		// 从x出发，判断是否能够到达y
		for y := 0; y < n; y += 1 {
			// 如果可以达到，并且距离更新的话，则更新距离，并将y放到q中
			if edges[x][y] && (dis[y] == -1 || dis[y] > dis[x]+1) {
				dis[y] = dis[x] + 1
				q = append(q, y)
			}
		}
	}
	// 判断到达targe的最近距离，可能有多个
	ans := math.MaxInt32
	for _, r := range rec[target] {
		if dis[r] != -1 && dis[r] < ans {
			ans = dis[r]
		}
	}
	if ans < math.MaxInt32 {
		return ans
	}
	return -1
}

func main285() {
	fmt.Println(numBusesToDestination([][]int{
		{1, 2, 7}, {3, 6, 7},
	}, 1, 6))
}
