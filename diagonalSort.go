package main

import (
	"fmt"
	"sort"
)

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	r := make([][]int, m)
	for i := 0; i < m; i += 1 {
		r[i] = make([]int, n)
	}
	get := func(x, y int) (res []int) {
		for x < m && y < n {
			res = append(res, mat[x][y])
			x, y = x+1, y+1
		}
		return
	}
	set := func(res []int, x, y int) {
		for _, v := range res {
			r[x][y] = v
			x, y = x+1, y+1
		}
	}

	for i := 0; i < m; i += 1 {
		h := get(i, 0)
		sort.Slice(h, func(i, j int) bool { return h[i] < h[j] })
		set(h, i, 0)
	}
	for i := 1; i < n; i += 1 {
		h := get(0, i)
		sort.Slice(h, func(i, j int) bool { return h[i] < h[j] })
		set(h, 0, i)
	}
	return r
}

func main071030() {
	fmt.Println(diagonalSort([][]int{
		{3, 3, 1, 1},
		{2, 2, 1, 2},
		{1, 1, 1, 2},
	}))
}
