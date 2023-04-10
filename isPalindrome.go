package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xah8k6/
// 验证回文串

// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
// 字母和数字都属于字母数字字符。
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。

func isPalindrome(s string) bool {
	helper := make([]rune, 0)
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			helper = append(helper, c)
		} else if c >= 'A' && c <= 'Z' {
			helper = append(helper, c+'a'-'A')
		} else if c >= '0' && c <= '9' {
			helper = append(helper, c)
		}
	}
	i, j := 0, len(helper)-1
	for i <= j {
		if helper[i] != helper[j] {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}

func main43() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("0P"))
}
