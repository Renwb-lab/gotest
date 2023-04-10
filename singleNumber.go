package main

import "fmt"

//	https://leetcode.cn/leetbook/read/top-interview-questions/xm0u83/
//
// 只出现一次的数字
func singleNumber(nums []int) int {
	ret := 0
	for _, val := range nums {
		ret ^= val
	}
	return ret
}

func main41() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
