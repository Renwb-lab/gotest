package main

import "fmt"

// 100033. 最大合金数
// https://leetcode.cn/contest/weekly-contest-363/problems/maximum-number-of-alloys/

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	// idx,  budget, stock
	minVal := stock[0]
	for i := 1; i < n; i += 1 {
		if stock[i] < minVal {
			minVal = stock[i]
		}
	}

	ans := 0
	for _, com := range composition {

		check := func(num int) bool {
			money := 0
			for i := 0; i < n; i += 1 {
				st, ct, co := stock[i], cost[i], com[i]
				if st < co*num {
					money += (co*num - st) * ct
					if money > budget {
						return false
					}
				}
			}
			return true
		}

		left, right := 0, budget+minVal+1
		for left < right {
			mid := left + (right-left)/2
			if check(mid) {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left-1 > ans {
			ans = left - 1
		}
	}
	return ans
}

func main137() {
	fmt.Print(maxNumberOfAlloys(7, 10, 177,
		[][]int{
			{1, 8, 4, 1, 9, 7, 4},
			{8, 1, 3, 3, 9, 4, 5},
			{8, 5, 4, 2, 9, 9, 10},
			{2, 10, 3, 3, 3, 10, 8},
			{6, 3, 1, 3, 7, 1, 7},
			{3, 5, 7, 6, 8, 10, 10},
			{2, 10, 10, 2, 2, 7, 7},
			{3, 2, 10, 9, 4, 1, 2},
			{2, 7, 1, 8, 2, 7, 10},
			{10, 9, 2, 8, 10, 1, 4},
		},
		[]int{0, 3, 5, 10, 0, 9, 9}, []int{3, 5, 4, 8, 10, 1, 2}),
	)
}
