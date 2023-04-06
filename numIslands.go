package main

import "fmt"

// https://leetcode.cn/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	l, r := len(grid), len(grid[0])
	ret := 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		grid[i][j] = '#'
		if i-1 >= 0 && grid[i-1][j] == '1' {
			dfs(i-1, j)
		}
		if i+1 < l && grid[i+1][j] == '1' {
			dfs(i+1, j)
		}
		if j-1 >= 0 && grid[i][j-1] == '1' {
			dfs(i, j-1)
		}
		if j+1 < r && grid[i][j+1] == '1' {
			dfs(i, j+1)
		}
	}

	for i := 0; i < l; i += 1 {
		for j := 0; j < r; j += 1 {
			if grid[i][j] == '1' {
				dfs(i, j)
				ret += 1
			}
		}
	}
	return ret
}

func main21() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}
