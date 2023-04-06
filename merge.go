package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/merge-intervals/description/?favorite=2ckc81c

// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

func merge(intervals [][]int) [][]int {
	ret := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] ||
			(intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	l, preStart, preEnd := len(intervals), intervals[0][0], intervals[0][1]
	for i := 1; i < l; i += 1 {
		if intervals[i][0] > preEnd {
			ret = append(ret, []int{preStart, preEnd})
			preStart, preEnd = intervals[i][0], intervals[i][1]
		} else {
			// 特别注意
			if intervals[i][0] > preEnd {
				preEnd = intervals[i][1]
			}
		}
	}
	if len(ret) == 0 || ret[len(ret)-1][1] != preEnd {
		ret = append(ret, []int{preStart, preEnd})
	}
	return ret
}

func main23() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	ret := merge(intervals)
	for _, line := range ret {
		for _, i := range line {
			fmt.Printf("%d\t", i)
		}
		fmt.Println()
	}
}
