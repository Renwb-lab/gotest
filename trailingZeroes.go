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
	return ret
}

func main71() {
	fmt.Println(trailingZeroes(3))
}
