package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/description/

// Floyd 算法
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	mp := make([][]int, n)
	for i := 0; i < n; i += 1 {
		mp[i] = make([]int, n)
		for j := 0; j < n; j += 1 {
			mp[i][j] = math.MaxInt32
		}
		mp[i][i] = 0
	}

	for i := range edges {
		from, to, weight := edges[i][0], edges[i][1], edges[i][2]
		mp[from][to], mp[to][from] = weight, weight
	}

	for k := 0; k < n; k += 1 {
		for i := 0; i < n; i += 1 {
			for j := 0; j < n; j += 1 {
				mp[i][j] = min(mp[i][j], mp[i][k]+mp[k][j])
			}
		}
	}

	ans := []int{math.MaxInt32, -1}
	for i := 0; i < n; i += 1 {
		cnt := 0
		for j := 0; j < n; j += 1 {
			if mp[i][j] <= distanceThreshold {
				cnt += 1
			}
		}
		if cnt <= ans[0] {
			ans[0], ans[1] = cnt, i
		}
	}

	return ans[1]
}

func findTheCity2(n int, edges [][]int, distanceThreshold int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	graph := make([][]int, n)
	dis := make([][]int, n)
	vis := make([][]bool, n)
	for i := 0; i < n; i += 1 {
		graph[i] = make([]int, n)
		dis[i] = make([]int, n)
		vis[i] = make([]bool, n)
		for j := 0; j < n; j += 1 {
			graph[i][j] = math.MaxInt32 / 2
			dis[i][j] = math.MaxInt32 / 2
			vis[i][j] = false
		}
	}

	for i := range edges {
		from, to, weight := edges[i][0], edges[i][1], edges[i][2]
		graph[from][to], graph[to][from] = weight, weight
	}

	// 对于每一个节点都使用dijkstra算法计算距离其他节点的最近算法
	for i := 0; i < n; i += 1 {
		// 每次只能查找一个节点，所以需要便利n次
		dis[i][i] = 0
		for j := 0; j < n; j += 1 {
			// 寻找离节点i最近的节点t
			t := -1
			for k := 0; k < n; k += 1 {
				if !vis[i][k] && (t == -1 || dis[i][k] < dis[i][t]) {
					t = k
				}
			}
			// 通过节点t,更新所有的节点，并将节点标记为已经访问
			for k := 0; k < n; k += 1 {
				dis[i][k] = min(dis[i][k], dis[i][t]+graph[t][k])
			}
			vis[i][t] = true
		}
	}

	ans := []int{math.MaxInt32 / 2, -1}
	for i := 0; i < n; i += 1 {
		cnt := 0
		for j := 0; j < n; j += 1 {
			if dis[i][j] <= distanceThreshold {
				cnt += 1
			}
		}
		if cnt <= ans[0] {
			ans[0], ans[1] = cnt, i
		}
	}

	return ans[1]
}

func main() {
	fmt.Println(findTheCity(4, [][]int{
		{0, 1, 3},
		{1, 2, 1},
		{1, 3, 4},
		{2, 3, 1},
	}, 4))
}
