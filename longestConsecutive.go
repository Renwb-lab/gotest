package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/x2xmre/

// 最长连续序列

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}

	ret := 0
	for n := range m {
		// 如果前面一个不存在的话，则表示已n开头
		if !m[n-1] {
			cur, curRet := n, 0
			for m[cur] {
				cur, curRet = cur+1, curRet+1
			}
			if curRet > ret {
				ret = curRet
			}
		}
	}

	return ret
}

func main56() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
