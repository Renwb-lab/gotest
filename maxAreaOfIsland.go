package main

import "fmt"

// https://leetcode.cn/problems/max-area-of-island/?company_slug=ubiquanthr

// 给你一个大小为 m x n 的二进制矩阵 grid 。
//
// 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
//
// 岛屿的面积是岛上值为 1 的单元格的数目。
//
// 计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

func maxAreaOfIsland(grid [][]int) int {
	forwards := [][]int{
		{0, 1}, {0, -1}, {-1, 0}, {1, 0},
	}
	l, r := len(grid), len(grid[0])
	var dp func(x, y int) int
	dp = func(x, y int) int {
		grid[x][y] = 2
		ret := 1
		for i := 0; i < 4; i += 1 {
			xx, yy := x+forwards[i][0], y+forwards[i][1]
			if xx >= 0 && xx < l && yy >= 0 && yy < r && grid[xx][yy] == 1 {
				ret += dp(xx, yy)
			}
		}
		return ret
	}
	ret := 0
	for i := 0; i < l; i += 1 {
		for j := 0; j < r; j += 1 {
			if grid[i][j] == 1 {
				tr := dp(i, j)
				if tr > ret {
					ret = tr
				}
			}
		}
	}
	return ret
}

func main34() {
	// grid := [][]int{
	// 	{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	// 	{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
	// 	{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	// }
	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}
	fmt.Println(maxAreaOfIsland(grid))
}
