package main

import (
	"fmt"
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	h := make([][]int, n)
	for i := 0; i < n; i += 1 {
		h[i] = []int{profit[i], difficulty[i]}
	}
	sort.Slice(h, func(i, j int) bool {
		return h[i][0] > h[j][0]
	})
	ans := 0
	for _, w := range worker {
		idx := sort.Search(n, func(i int) bool {
			return w >= h[i][1]
		})
		if idx < n && idx >= 0 {
			ans += h[idx][0]
		}
	}
	return ans
}

func main262() {
	fmt.Println(maxProfitAssignment(
		[]int{85, 47, 57},
		[]int{24, 66, 99},
		[]int{40, 25, 25},
	))
}
