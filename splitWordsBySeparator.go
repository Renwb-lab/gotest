package main

import "fmt"

// https://leetcode.cn/problems/split-strings-by-separator/
func splitWordsBySeparator(words []string, separator byte) []string {
	split := func(s string) []string {
		byteArr := []byte(s)
		ans := []string{}
		b, e, l := 0, 0, len(s)
		for b < l {
			for b < l && s[b] == separator {
				b += 1
			}
			if b == l {
				break
			}
			e = b
			for e < l && s[e] != separator {
				e += 1
			}
			ans = append(ans, string(byteArr[b:e]))
			b = e + 1
		}
		return ans
	}
	ans := make([]string, 0)
	for i := range words {
		ans = append(ans, split(words[i])...)
	}
	return ans
}

func main310() {
	fmt.Println(splitWordsBySeparator([]string{"one.two.three", "four.five", "six"}, '.'))
}
