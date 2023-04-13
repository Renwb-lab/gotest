package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/x25oeg/
// 打家劫舍

func rob(nums []int) int {
	robCur, notRobCur := make([]int, len(nums)), make([]int, len(nums))
	robCur[0], notRobCur[0] = nums[0], 0
	maxFunc := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ret := nums[0]
	for i := 1; i < len(nums); i += 1 {
		robCur[i] = notRobCur[i-1] + nums[i]
		notRobCur[i] = maxFunc(notRobCur[i-1], robCur[i-1])
		ret = maxFunc(ret, maxFunc(robCur[i], notRobCur[i]))
	}
	return ret
}

func main57() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}
