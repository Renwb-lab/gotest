package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/x2we65/
// 阶乘后的零

func trailingZeroes(n int) int {
	ret := 0
	for n > 0 { // 主要25中有两个5
		n = n / 5
		// 特别重要，
		ret += n
	}
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25
	// 0, 0, 0, 0, 1, 0, 0, 0, 0, 1,  0,  0,   0,  0, 1,   0,  0,  0,  0, 1,  0,  0,   0, 0,  2,
	// 0, 0, 0, 0, 1, 1, 1, 1, 1, 2,  2,  2,  2,  2,  3,  3,  3,  3,  3,  4,  4,  4,  4,  4,  6
	return ret
}

func main71() {
	fmt.Println(trailingZeroes(3))
}
