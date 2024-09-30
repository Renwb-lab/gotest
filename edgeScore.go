package main

// https://leetcode.cn/problems/node-with-highest-edge-score
func edgeScore(edges []int) int {
	n := len(edges)
	degree := make([]int, n)
	for i, e := range edges {
		degree[e] += i
	}
	res, deg := 0, degree[0]
	for i := 1; i < n; i += 1 {
		if degree[i] > deg {
			res, deg = i, degree[i]
		}
	}
	return res
}
