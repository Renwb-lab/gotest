package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xabtn6/
// 寻找重复数

func findDuplicate(nums []int) int {
	first, second := 0, 0
	for {
		first = nums[first]
		second = nums[nums[second]]
		if first == second {
			break
		}
	}
	for first = 0; first != second; {
		first = nums[first]
		second = nums[second]
	}
	return first
}

func main65() {
	nums := []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicate(nums))
}
