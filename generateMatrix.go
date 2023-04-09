package main

import "fmt"

// https://leetcode.cn/problems/spiral-matrix-ii/?company_slug=ubiquanthr

// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

func generateMatrix(n int) [][]int {
	ret := make([][]int, 0)
	for i := 0; i < n; i += 1 {
		ret = append(ret, make([]int, n))
	}
	cur, sum := 0, n*n
	minx, maxx := 0, n-1
	miny, maxy := 0, n-1
	x, y := 0, -1
	for cur < sum {
		// left
		for y += 1; y <= maxy; y += 1 {
			cur += 1
			ret[x][y] = cur
		}
		y, minx = maxy, minx+1
		// down
		for x += 1; x <= maxx; x += 1 {
			cur += 1
			ret[x][y] = cur
		}
		x, maxy = maxx, maxy-1
		// left
		for y -= 1; y >= miny; y -= 1 {
			cur += 1
			ret[x][y] = cur
		}
		y, maxx = miny, maxx-1
		// up
		for x -= 1; x >= minx; x -= 1 {
			cur += 1
			ret[x][y] = cur
		}
		x, miny = minx, miny+1
	}

	return ret
}

func main31() {
	ret := generateMatrix(3)
	for _, line := range ret {
		for _, v := range line {
			fmt.Printf("%d\t", v)
		}
		fmt.Println()
	}
}
