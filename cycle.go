package main

import (
	"fmt"
)

func cycle(n int) [][]int {
	ret := make([][]int, 0)
	for i := 0; i < n; i += 1 {
		line := make([]int, 0)
		for j := 0; j < n; j += 1 {
			line = append(line, 0)
		}
		ret = append(ret, line)
	}

	x, y := n/2, n/2
	step, val := 0, 1
	ret[x][y] = val

	for val < n*n {
		step += 1
		// right
		for i := 0; i < step && val < n*n; i += 1 {
			y += 1
			val += 1
			ret[x][y] = val
		}
		// up
		for i := 0; i < step && val < n*n; i += 1 {
			x -= 1
			val += 1
			ret[x][y] = val
		}
		step += 1
		// left
		for i := 0; i < step && val < n*n; i += 1 {
			y -= 1
			val += 1
			ret[x][y] = val
		}
		// down
		for i := 0; i < step && val < n*n; i += 1 {
			x += 1
			val += 1
			ret[x][y] = val
		}
	}
	return ret
}

func main() {
	ret := cycle(9)
	for _, line := range ret {
		for _, v := range line {
			fmt.Printf("%d\t", v)
		}
		fmt.Printf("\n")
	}
}
