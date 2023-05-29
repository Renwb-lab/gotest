package main

import (
	"fmt"
)

// https://leetcode.cn/problems/maximum-profit-from-trading-stocks/
// 2291. 最大股票收益

func maximumProfit(present []int, future []int, budget int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(present)
	helper := make([][]int, n+1)
	for i := 0; i <= n; i += 1 {
		helper[i] = make([]int, budget+1)
	}
	for i := 1; i <= n; i += 1 {
		for j := 0; j < budget; j += 1 {
			if j < present[i-1] {
				helper[i][j] = helper[i-1][j]
			} else {
				helper[i][j] = max(helper[i-1][j], helper[i-1][j-present[i-1]]+future[i-1]-present[i-1])
			}
		}
	}
	return helper[n][budget]
}

func main115() {
	present := []int{5, 4, 6, 2, 3}
	future := []int{8, 5, 4, 3, 5}
	fmt.Println(maximumProfit(present, future, 10))
}
