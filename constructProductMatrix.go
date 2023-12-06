package main

import "fmt"

func constructProductMatrix(grid [][]int) [][]int {
	const mod = 12345
	n, m := len(grid), len(grid[0])
	res := make([][]int, n)
	suf := 1
	for i := n - 1; i >= 0; i -= 1 {
		res[i] = make([]int, m)
		for j := m - 1; j >= 0; j -= 1 {
			res[i][j] = suf
			suf = suf * grid[i][j] % mod
		}
	}
	pre := 1
	for i := 0; i < n; i += 1 {
		for j := 0; j < m; j += 1 {
			res[i][j] = res[i][j] * pre % mod
			pre = pre * grid[i][j] % mod
		}
	}
	return res
}

func main153() {
	fmt.Println(constructProductMatrix([][]int{
		{1, 2}, {3, 4}, {5, 6},
	}))
	fmt.Println(constructProductMatrix([][]int{
		{12345}, {2}, {1},
	}))
}
