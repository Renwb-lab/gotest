package main

func minimumArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	left, right := n, 0
	top, botton := m, 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			if i < top {
				top = i
			}
			if i > botton {
				botton = i
			}
			if j < left {
				left = j
			}
			if j > right {
				right = j
			}
		}
	}
	return (right - left + 1) * (botton - top + 1)
}
