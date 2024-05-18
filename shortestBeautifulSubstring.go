package main

import (
	"fmt"
)

func shortestBeautifulSubstring(s string, k int) string {
	better := func(t, r string) bool {
		if len(r) == 0 {
			return true
		}
		if len(t) < len(r) {
			return true
		}
		if len(t) > len(r) {
			return false
		}
		l := len(t)
		for i := 0; i < l; i += 1 {
			if t[i] == r[i] {
				continue
			}
			if t[i] == '1' && r[i] == '0' {
				return false
			}
			if t[i] == '0' && r[i] == '1' {
				return true
			}
		}
		return false
	}

	left, right, l := 0, 0, len(s)
	count := 0
	r := ""
	// 滑动窗口
	for right < l {
		if s[right] == '0' {
			right += 1
			continue
		}
		count += 1
		right += 1

		for count == k {
			t := s[left:right]
			fmt.Println(t)
			if better(t, r) {
				r = t
			}
			if s[left] == '1' {
				count -= 1
			}
			left += 1
		}
	}
	return r
}

func main() {
	fmt.Println(shortestBeautifulSubstring("001110101101101111", 10))
}
