package main

import "fmt"

// https://leetcode.cn/problems/one-away-lcci/?favorite=xb9lfcwi

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if m < n {
		return oneEditAway(second, first)
	}
	if m-1 > n {
		return false
	}
	for i, ch := range second {
		if first[i] != byte(ch) {
			// 进行更改
			if m == n {
				return second[i+1:] == first[i+1:]
			}
			// 添加或者删除一个字符
			return second[i:] == first[i+1:]
		}
	}
	return true
}

func main12() {
	ret := oneEditAway("pale", "ple")
	fmt.Println(ret)
	ret = oneEditAway("pales", "pal")
	fmt.Println(ret)
}
