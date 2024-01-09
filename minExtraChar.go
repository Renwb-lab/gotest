package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/extra-characters-in-a-string/description/
func minExtraChar(s string, dictionary []string) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	n := len(s)
	// 注意是n+1, 有以下两个用途：
	// 1. helper[i] = helper[i-1] + 1
	// 2. i := 1 时，s[0:i]刚好可以去到s[0]的数据
	// helper[i] 表示s[0,i)时的最少数量
	// helper[0] 没有含义
	// helper[1] 表示只有s[0]的最少数量
	helper := make([]int, n+1)
	for i := 1; i <= n; i += 1 {
		helper[i] = math.MaxInt
	}
	mp := make(map[string]int)
	for _, e := range dictionary {
		mp[e] += 1
	}
	for i := 1; i <= n; i += 1 {
		helper[i] = helper[i-1] + 1
		for j := i - 1; j >= 0; j -= 1 {
			if _, ok := mp[s[j:i]]; ok {
				helper[i] = min(helper[j], helper[i])
			}
		}
	}
	return helper[n]
}

func main() {
	fmt.Println(minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
}
