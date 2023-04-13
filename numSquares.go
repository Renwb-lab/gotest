package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/x2959v/
// 完全平方数

func numSquares(n int) int {
	helper := make([]int, n+1)
	for i := 1; i*i <= n; i += 1 {
		helper[i*i] = 1
	}
	if helper[n] != 0 {
		return 1
	}

	for i := 2; i <= n; i += 1 {
		if helper[i] != 0 {
			continue
		}
		t := math.MaxInt
		for j := 1; j < i; j += 1 {
			if helper[i-j]+helper[j] < t {
				t = helper[i-j] + helper[j]
			}
		}
		helper[i] = t
	}
	return helper[n]
}

func main58() {
	fmt.Println(numSquares(12))
}
