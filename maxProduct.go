package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xmk3rv/

// 乘积最大子数组

func maxProduct(nums []int) int {
	ret := nums[0]
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	maxv, minv := nums[0], nums[0]
	for i := 1; i < len(nums); i += 1 {
		tmp1, tmp2 := maxv*nums[i], minv*nums[i]
		maxv = max(nums[i], max(tmp1, tmp2))
		minv = min(nums[i], min(tmp1, tmp2))
		ret = max(ret, maxv)
	}
	return ret
}

func main50() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}
