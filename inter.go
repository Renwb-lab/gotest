package main

import "fmt"

// 打印回旋矩阵
// 示例：

// 输入 2，输出：
// 1   2
// 4   3

func printCycle(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i += 1 {
		ret[i] = make([]int, n)
	}
	x, y := 0, -1
	idx := 0
	minx, maxx := 0, n-1
	miny, maxy := 0, n-1
	for idx < n*n {
		// 输入 3，输出：
		// 1   2    3
		// 8   9    4
		// 7   6    5

		// right
		for y += 1; y <= maxy && idx <= n*n; y += 1 {
			idx += 1
			ret[x][y] = idx
		}
		minx += 1
		y = maxy
		// down
		for x += 1; x <= maxx && idx <= n*n; x += 1 {
			idx += 1
			ret[x][y] = idx
		}
		maxy -= 1
		x = maxx
		// left
		y -= 1
		for ; y >= miny && idx <= n*n; y -= 1 {
			idx += 1
			ret[x][y] = idx
		}
		maxx -= 1
		y = miny
		// up
		x -= 1
		for ; x >= minx && idx <= n*n; x -= 1 {
			idx += 1
			ret[x][y] = idx
		}
		miny += 1
		x = minx
	}
	return ret
}

func main80() {
	ret := printCycle(4)
	for _, line := range ret {
		for _, v := range line {
			fmt.Printf("%d\t", v)
		}
		fmt.Println()
	}
}
