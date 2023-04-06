package main

import "fmt"

// https://leetcode.cn/problems/spiral-matrix/

// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]

func spiralOrder(matrix [][]int) []int {
	l, r := len(matrix), len(matrix[0])
	minx, maxx := 0, l-1
	miny, maxy := 0, r-1
	x, y := 0, -1
	ret := make([]int, 0)
	for len(ret) < l*r {
		// right
		for y += 1; len(ret) < l*r && y <= maxy; y += 1 {
			ret = append(ret, matrix[x][y])
		}
		y, minx = maxy, minx+1
		// down
		for x += 1; len(ret) < l*r && x <= maxx; x += 1 {
			ret = append(ret, matrix[x][y])
		}
		x, maxy = maxx, maxy-1
		// left
		for y -= 1; len(ret) < l*r && y >= miny; y -= 1 {
			ret = append(ret, matrix[x][y])
		}
		y, maxx = miny, maxx-1
		// up
		for x -= 1; len(ret) < l*r && x >= minx; x -= 1 {
			ret = append(ret, matrix[x][y])
		}
		x, miny = minx, miny+1
	}
	return ret
}

func main25() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	ret := spiralOrder(matrix)
	for _, v := range ret {
		fmt.Printf("%d\t", v)
	}
}
