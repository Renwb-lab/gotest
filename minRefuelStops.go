package main

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/minimum-number-of-refueling-stops/submissions/?envType=daily-question&envId=2024-10-07
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	l := len(stations)
	// 保存从i到j节点到距离
	graph := make([][]int, l+1)
	for i := 0; i <= l; i += 1 {
		graph[i] = make([]int, l+1)
	}
	for i := 0; i < l; i += 1 {
		graph[0][i+1] = stations[i][0]
		for j := i + 1; j < l; j += 1 {
			graph[i+1][j+1] = stations[j][0] - stations[i][0]
		}
	}
	nodes := make([]int, l+1)
	nodes[0] = startFuel
	for i := 0; i < l; i += 1 {
		nodes[i+1] = stations[i][1]
	}

	ans := -1
	var dfs func(cur int, curFuel int, curTimes int, curDistance int)
	dfs = func(cur, curFuel, curTimes, curDistance int) {
		if curDistance >= target {
			if ans == -1 || ans > curTimes {
				ans = curTimes
			}
			return
		}
		if curFuel == 0 {
			return
		}
		for i := cur + 1; i <= l; i += 1 {
			if graph[cur][i] > 0 && graph[cur][i] <= curFuel {
				dfs(i, curFuel-graph[cur][i]+nodes[i], curTimes+1, curDistance+graph[cur][i])
			}
		}
		dfs(cur, 0, curTimes, curDistance+curFuel)
	}
	dfs(0, startFuel, 0, 0)
	return ans
}

func minRefuelStops2(target, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int, n+1)
	// 用 dp[i] 表示加油 i 次的最大行驶英里数。
	dp[0] = startFuel
	for i, s := range stations {
		for j := i; j >= 0; j -= 1 {
			if dp[j] >= s[0] {
				dp[j+1] = max(dp[j+1], dp[j]+s[1])
			}
		}
	}
	for i, d := range dp {
		if d >= target {
			return i
		}
	}
	return -1
}

func minRefuelStops3(target, startFuel int, stations [][]int) (ans int) {
	stations = append(stations, []int{target, 0})
	prePosition, curFuel := 0, startFuel
	fuelHeap := &hp3{}
	for _, station := range stations {
		position, fuel := station[0], station[1]
		curFuel -= position - prePosition       // 每行驶 1 英里用掉 1 升汽油
		for fuelHeap.Len() > 0 && curFuel < 0 { // 没油了
			curFuel += heap.Pop(fuelHeap).(int) // 选油量最多的油桶
			ans++
		}
		if curFuel < 0 { // 无法到达
			return -1
		}
		heap.Push(fuelHeap, fuel) // 留着后面加油
		prePosition = position
	}
	return
}

type hp3 struct{ sort.IntSlice }

func (h hp3) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (h *hp3) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp3) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
