package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	line := []int{}
	var dp func(idx int, target int)
	dp = func(idx, target int) {
		if target == 0 {
			res = append(res, append([]int{}, line...))
			return
		}
		if idx == len(candidates) {
			return
		}
		if candidates[idx] <= target {
			line = append(line, candidates[idx])
			dp(idx, target-candidates[idx])
			line = line[:len(line)-1]
		}
		dp(idx+1, target)
	}
	dp(0, target)
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
