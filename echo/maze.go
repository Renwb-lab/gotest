package main

import (
	"fmt"
)

func compute(h [][]int, m, n int) [][]int {
	res := make([][]int, 0)
	forwards := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	var dp func(x, y int)
	dp = func(x, y int) {
		res = append(res, []int{x, y})
		if x == m-1 && y == n-1 {
			print(res)
			return
		}
		for _, f := range forwards {
			x1, y1 := x+f[0], y+f[1]
			if x1 >= 0 && x1 < m && y1 >= 0 && y1 < n && h[x1][y1] == 0 {
				h[x1][y1] = 2
				dp(x1, y1)
				h[x1][y1] = 0
			}
		}
		res = res[:len(res)-1]
	}
	dp(0, 0)
	return res
}

func print(res [][]int) {
	for i := 0; i < len(res); i += 1 {
		fmt.Printf("(%d,%d), ", res[i][0], res[i][1])
	}
	fmt.Println()
}

func main() {
	// m, n := 0, 0
	// fmt.Scan(&m, &n)
	// h := make([][]int, m, m)
	// for i := 0; i < m; i += 1 {
	//     h[i] = make([]int, n, n)
	// }
	// for i := 0; i < m; i += 1 {
	//     for j := 0; j < n; j += 1 {
	//         fmt.Scan(&h[i][j])
	//     }
	// }
	m, n := 5, 5
	h := [][]int{
		[]int{0, 1, 0, 0, 0},
		[]int{0, 1, 1, 1, 0},
		[]int{0, 0, 0, 0, 0},
		[]int{0, 1, 1, 1, 0},
		[]int{0, 0, 0, 1, 0},
	}
	compute(h, m, n)
	// print(res)
}
