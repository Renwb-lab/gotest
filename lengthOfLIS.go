package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/x29fxj/
// 最长上升子序列
func lengthOfLIS(nums []int) int {
	helper := make([]int, len(nums))
	ret := 0
	for _, v := range nums {
		left, right := 0, ret
		for left < right {
			mid := left + (right-left)>>1
			if helper[mid] < v {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left >= ret {
			ret += 1
		}
		helper[left] = v
	}
	return ret
}

func main59() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
