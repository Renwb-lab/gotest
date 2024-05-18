package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	visited := [9]bool{false}
	var dp func(idx int, target int)
	dp = func(idx, target int) {
		if target == 0 {
			line := []int{}
			for i, b := range visited {
				if b {
					line = append(line, i+1)
				}
			}
			if len(line) == k {
				res = append(res, line)
			}
			return
		}
		if idx >= 9 {
			return
		}
		if target >= idx+1 {
			visited[idx] = true
			dp(idx+1, target-idx-1)
			visited[idx] = false
		}
		dp(idx+1, target)
	}
	dp(0, n)
	return res
}

func main() {
	fmt.Print(combinationSum3(9, 45))
}
