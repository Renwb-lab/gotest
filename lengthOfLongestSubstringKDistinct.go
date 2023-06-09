package main

import "fmt"

// https://leetcode.cn/problems/longest-substring-with-at-most-k-distinct-characters/description/
// 340. 至多包含 K 个不同字符的最长子串

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	left, right, size := 0, 0, len(s)
	windows, validNum := map[byte]int{}, 0
	ret := 0
	for right < size {
		c := s[right]
		windows[c] += 1
		if windows[c] == 1 {
			validNum += 1
		}
		for validNum > k {
			c := s[left]
			windows[c] -= 1
			if windows[c] == 0 {
				validNum -= 1
			}
			left += 1
		}
		ret = max(ret, right-left+1)
		right += 1
	}
	return ret
}

func main122() {
	fmt.Println(lengthOfLongestSubstringKDistinct("eceba", 2))
}
