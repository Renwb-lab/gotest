package main

import "fmt"

// https://leetcode.cn/problems/ransom-note/
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int, 0)
	for i := range magazine {
		c := magazine[i]
		m[c] += 1
	}
	for i := range ransomNote {
		c := ransomNote[i]
		if n, ok := m[c]; ok && n > 0 {
			m[c] -= 1
		} else {
			return false
		}

	}
	return true
}

func main07105() {
	fmt.Println(canConstruct("aab", "baa"))
}
