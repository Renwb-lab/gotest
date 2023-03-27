package main

import "fmt"

// https://leetcode.cn/problems/bracket-lcci/?favorite=xb9lfcwi

func recursion(line string, left int, right int, n int, ret *[]string) {
	if left == n && right == n {
		*ret = append(*ret, line)
	}
	if left < n {
		recursion(line+"(", left+1, right, n, ret)
	}
	if right < left {
		recursion(line+")", left, right+1, n, ret)
	}
}

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	recursion("", 0, 0, n, &ret)
	return ret
}

func main7() {
	ret := generateParenthesis(3)
	for _, line := range ret {
		fmt.Println(line)
	}
}
