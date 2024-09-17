package main

import (
	"fmt"
	"sort"
)

func findLUSlength(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})
	isSubStr := func(s, t string) bool {
		si := 0
		for ti := range t {
			if s[si] == t[ti] {
				si += 1
				if si == len(s) {
					return true
				}
			}
		}
		return false
	}
	for i, s := range strs {
		flag := true
		for j, t := range strs {
			if j != i && isSubStr(s, t) {
				flag = false
				break
			}
		}
		if flag {
			return len(strs[i])
		}
	}
	return -1
}

func main071032() {
	fmt.Println(findLUSlength([]string{"aaa", "aaa", "aa"}))
}
