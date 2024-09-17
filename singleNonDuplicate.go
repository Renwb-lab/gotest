package main

import "fmt"

// https://leetcode.cn/problems/single-element-in-a-sorted-array/description/
func singleNonDuplicate(nums []int) int {
	b, e := 0, len(nums)-1
	for b < e {
		m := b + (e-b)>>1
		if (m%2 == 0 && nums[m] == nums[m+1]) ||
			(m%2 == 1 && nums[m] == nums[m-1]) {
			b = m + 1
		} else {
			e = m
		}
	}
	return b
}

func main306() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
}
