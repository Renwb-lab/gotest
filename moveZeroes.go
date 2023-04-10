package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xmy9jh/
// 移动零

func moveZeroes(nums []int) {
	size := len(nums)
	b := -1
	for i := 0; i < size; i += 1 {
		if nums[i] != 0 {
			b += 1
			nums[i], nums[b] = nums[b], nums[i]
		}
	}
}

func main52() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	for _, n := range nums {
		fmt.Printf("%d\t", n)
	}
}
