package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xar9lv/
// 有效的字母异位词

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

// 注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
//
// 示例 1:
//
// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 示例 2:
//
// 输入: s = "rat", t = "car"
// 输出: false

func isAnagram(s string, t string) bool {
	m := map[rune]int{}
	for _, c := range s {
		if _, ok := m[c]; !ok {
			m[c] = 0
		}
		m[c] += 1
	}
	vaild := 0
	for _, c := range t {
		if _, ok := m[c]; !ok {
			return false
		}
		m[c] -= 1
		if m[c] == 0 {
			vaild += 1
		} else if m[c] < 0 {
			return false
		}
	}
	return vaild == len(m)
}

func main48() {
	s, t := "anagram", "nagaram"
	fmt.Println(isAnagram(s, t))
}
