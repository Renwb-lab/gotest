package main

import "fmt"

// https://lelleetcode.cn/problems/longest-substring-without-repeating-characters/description/
// 3. 无重复字符的最长子串

func lengthOfLongestSubstring(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	res := 0
	windows := make(map[byte]int)
	left, right, length := 0, 0, len(s)
	for right < length {
		windows[s[right]] += 1

		for windows[s[right]] > 1 {
			windows[s[left]] -= 1
			left += 1
		}

		res = max(res, right-left+1)
		right += 1
	}
	return res
}

func main119() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
