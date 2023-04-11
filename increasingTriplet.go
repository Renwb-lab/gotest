package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/xmb141/
// 递增的三元子序列

func increasingTriplet(nums []int) bool {
	size := len(nums)
	if size < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt32
	for i := 1; i < size; i += 1 {
		if nums[i] > second {
			return true
		} else if nums[i] > first {
			second = nums[i]
		} else {
			first = nums[i]
		}
	}
	return false
}

func main54() {
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}
