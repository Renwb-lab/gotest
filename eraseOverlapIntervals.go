package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/non-overlapping-intervals/description/

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return (intervals[i][0] < intervals[j][0])
	})
	n := len(intervals)
	// dp[i]表示截止到第i个下标的地方能够形成的最大不重复数组的数量
	dp := make([]int, n, n)
	for i := 0; i < n; i += 1 {
		dp[i] = 1
	}
	max := func(args ...int) int {
		res := args[0]
		for _, n := range args[1:] {
			if n > res {
				res = n
			}
		}
		return res
	}
	for i := 1; i < n; i += 1 {
		for j := 0; j < i; j += 1 {
			if intervals[j][1] <= intervals[i][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return n - max(dp...)
}

func eraseOverlapIntervals2(intervals [][]int) int {
	// 计算最多有多少个不重叠的数量，使用总数减去该值，即为最终结果
	sort.Slice(intervals, func(i, j int) bool {
		return (intervals[i][1] < intervals[j][1])
	})
	res, n := 1, len(intervals)
	right := intervals[0][1]
	for i := 1; i < n; i += 1 {
		if intervals[i][0] >= right {
			res += 1
			right = intervals[i][1]
		}
	}
	return n - res
}

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{
		{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26},
	}))
	fmt.Println(eraseOverlapIntervals2([][]int{
		{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26},
	}))
}
