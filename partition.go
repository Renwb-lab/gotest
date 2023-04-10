package main

import (
	"fmt"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/xaxi62/
// 分割回文串

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
// 回文串 是正着读和反着读都一样的字符串。

// 示例 1：

// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
// 示例 2：

// 输入：s = "a"
// 输出：[["a"]]

func partition(s string) [][]string {
	ret, line := make([][]string, 0), make([]string, 0)
	size := len(s)
	helper := make([][]bool, size)
	for i := 0; i < size; i += 1 {
		helper[i] = make([]bool, size)
	}
	for i := 0; i < size; i += 1 {
		helper[i][i] = true
	}
	for i := size - 1; i >= 0; i -= 1 {
		for j := i + 1; j < size; j += 1 {
			if j-i == 1 {
				helper[i][j] = s[i] == s[j]
			} else {
				helper[i][j] = (s[i] == s[j]) && helper[i+1][j-1]
			}
		}
	}

	var dp func(idx int)
	dp = func(idx int) {
		if idx >= size {
			ret = append(ret, append([]string{}, line...))
			return
		}
		for i := idx + 1; i <= size; i += 1 {
			if helper[idx][i-1] {
				line = append(line, s[idx:i])
				dp(i)
				line = line[:len(line)-1]
			}
		}
	}

	dp(0)
	return ret
}

func main44() {
	s := "aab"
	ret := partition(s)
	for _, line := range ret {
		for _, item := range line {
			fmt.Printf("%s\t", item)
		}
		fmt.Println()
	}
}
