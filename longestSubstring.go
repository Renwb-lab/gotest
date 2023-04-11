package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xafdmc/
// 至少有K个重复字符的最长子串

// 给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。
//
//
// 示例 1：
//
// 输入：s = "aaabb", k = 3
// 输出：3
// 解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。
// 示例 2：
//
// 输入：s = "ababbc", k = 2
// 输出：5
// 解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。

func longestSubstring(s string, k int) int {
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if r-l+1 < k {
			return 0
		}
		m := map[rune]int{}
		for i := l; i <= r; i += 1 {
			c := rune(s[i])
			if _, ok := m[c]; !ok {
				m[c] = 1
			} else {
				m[c] += 1
			}
		}
		for r-l+1 >= k && m[rune(s[l])] < k {
			l += 1
		}
		for r-l+1 >= k && m[rune(s[r])] < k {
			r -= 1
		}
		if r-l+1 < k {
			return 0
		}
		for i := l; i <= r; i += 1 {
			if m[rune(s[i])] < k {
				left := dfs(l, i-1)
				right := dfs(i+1, r)
				if left > right {
					return left
				} else {
					return right
				}
			}
		}

		return r - l + 1
	}
	return dfs(0, len(s)-1)
}

func main() {
	s := "bbaaacbd"
	fmt.Println(longestSubstring(s, 3))
}
