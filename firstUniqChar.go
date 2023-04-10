package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xaph0j/

func firstUniqChar(s string) int {
	m := map[rune]int{}
	for _, c := range s {
		if v, ok := m[c]; !ok {
			m[c] = 1
		} else {
			m[c] = v + 1
		}
	}
	for i, c := range s {
		if v, ok := m[c]; ok {
			if v == 1 {
				return i
			}
		}
	}
	return -1
}

func main49() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
