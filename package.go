package main

import (
	"fmt"
)

func max(xs ...int) int {
	t := xs[0]
	for _, x := range xs[1:] {
		if x > t {
			t = x
		}
	}
	return t
}

func main287() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	n = n / 10

	helper := make([][3][3]int, m+1, m+1)
	for i := 1; i <= m; i += 1 {
		v, p, q := 0, 0, 0
		fmt.Scan(&v, &p, &q)
		v = v / 10
		if q == 0 {
			helper[i][0] = [3]int{v, p, v * p}
		} else {
			if helper[q][1][0] == 0 {
				helper[q][1] = [3]int{v, p, v * p}
			} else {
				helper[q][2] = [3]int{v, p, v * p}
			}
		}
	}
	h := make([][]int, n+1)
	for i := 0; i <= n; i += 1 {
		h[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i += 1 {
		for j := 1; j <= m; j += 1 {
			a, b := helper[j][0][0], helper[j][0][2]
			c, d := helper[j][1][0], helper[j][1][2]
			e, f := helper[j][2][0], helper[j][2][2]
			if i >= a {
				h[i][j] = max(b+h[i-a][j-1], h[i][j-1], h[i][j])
			} else {
				h[i][j] = h[i][j-1]
			}
			if i >= a+c {
				h[i][j] = max(b+d+h[i-a-c][j-1], h[i][j-1], h[i][j])
			}
			if i >= a+e {
				h[i][j] = max(b+f+h[i-a-e][j-1], h[i][j-1], h[i][j])
			}
			if i >= a+c+e {
				h[i][j] = max(b+d+f+h[i-a-c-e][j-1], h[i][j-1], h[i][j])
			}
		}
	}

	fmt.Println(h[n][m] * 10)
}
