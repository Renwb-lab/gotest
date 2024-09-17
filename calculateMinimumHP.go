package main

import "fmt"

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	if dungeon[m-1][n-1] <= 0 {
		dungeon[m-1][n-1] = 1 - dungeon[m-1][n-1]
	} else {
		dungeon[m-1][n-1] = 1
	}
	update := func(xa, ya, xb, yb int) {
		if dungeon[xa][ya] > 0 {
			if dungeon[xb][yb] > dungeon[xa][ya] {
				dungeon[xa][ya] = dungeon[xb][yb] - dungeon[xa][ya]
			} else {
				dungeon[xa][ya] = 1
			}
		} else {
			dungeon[xa][ya] = dungeon[xb][yb] - dungeon[xa][ya]
		}
	}
	// 最后一行
	for j := n - 2; j >= 0; j -= 1 {
		update(m-1, j, m-1, j+1)
	}
	// 最后一列
	for i := m - 2; i >= 0; i -= 1 {
		update(i, n-1, i+1, n-1)
	}

	for i := m - 2; i >= 0; i -= 1 {
		for j := n - 2; j >= 0; j -= 1 {
			a, b := 0, 0
			if dungeon[i][j+1] < dungeon[i+1][j] {
				a, b = i, j+1
			} else {
				a, b = i+1, j
			}
			update(i, j, a, b)
		}
	}
	return dungeon[0][0]
}

func main07106() {
	fmt.Println(calculateMinimumHP([][]int{
		{1, -2, 3},
		{2, -2, -2},
	}))
}
