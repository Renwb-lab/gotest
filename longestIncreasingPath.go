package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/x2osfr/
// 矩阵中的最长递增路径

func longestIncreasingPath(matrix [][]int) int {
	l, r := len(matrix), len(matrix[0])
	helper := make([][]int, l)
	for i := 0; i < l; i += 1 {
		helper[i] = make([]int, r)
	}
	for i := 0; i < l; i += 1 {
		for j := 0; j < r; j += 1 {
			helper[i][j] = -1
		}
	}

	forwards := [][]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if helper[x][y] != -1 {
			return helper[x][y]
		}
		ret := 0
		for _, f := range forwards {
			xx, yy := x+f[0], y+f[1]
			if xx >= 0 && xx < l && yy >= 0 && yy < r && matrix[xx][yy] > matrix[x][y] {
				t := dfs(xx, yy)
				if ret < t {
					ret = t
				}
			}
		}
		helper[x][y] = ret + 1
		return helper[x][y]
	}

	ret := 0
	for i := 0; i < l; i += 1 {
		for j := 0; j < r; j += 1 {
			if helper[i][j] == -1 {
				_ = dfs(i, j)
			}
			if helper[i][j] > ret {
				ret = helper[i][j]
			}
		}
	}
	return ret
}

func main61() {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	fmt.Println(longestIncreasingPath(matrix))
}
