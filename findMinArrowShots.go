package main

import (
	"sort"
)

// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/

// 452. 用最少数量的箭引爆气球
// 有一些球形气球贴在一堵用 XY 平面表示的墙面上。墙面上的气球记录在整数数组 points ，其中points[i] = [xstart, xend] 表示水平直径在 xstart 和 xend之间的气球。你不知道气球的确切 y 坐标。
// 一支弓箭可以沿着 x 轴从不同点 完全垂直 地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被 引爆 。可以射出的弓箭的数量 没有限制 。 弓箭一旦被射出之后，可以无限地前进。
// 给你一个数组 points ，返回引爆所有气球所必须射出的 最小 弓箭数 。

func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	res, cur := 1, points[0][1]
	// [1, 10], [9, 11], [2, 3], [3, 4]
	// [2, 3], [3, 4], [1, 10], [9, 11]
	for i := 1; i < len(points); i += 1 {
		if points[i][0] <= cur {
			continue
		} else {
			res += 1
			cur = points[i][1]
		}
	}
	return res
}

func main131() {
	// [1, 6], [2, 8], [7, 12], [10, 16]
	println(findMinArrowShots([][]int{
		{10, 16},
		{2, 8},
		{1, 6},
		{7, 12},
	}))
}
