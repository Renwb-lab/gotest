package main

import "fmt"

// https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-ii/
func minimumSum(nums []int) int {
	l := len(nums)
	right := make([]int, l)
	right[l-1] = nums[l-1]
	// 8,6,1,5,3
	// 1,1,1,3,3
	for i := l - 2; i >= 0; i -= 1 {
		if nums[i] < right[i+1] {
			right[i] = nums[i]
		} else {
			right[i] = right[i+1]
		}
	}

	min := nums[0]
	r := -1
	flag := false
	for i := 1; i < l-1; i += 1 {
		if min < nums[i] && nums[i] > right[i+1] {
			t := min + nums[i] + right[i+1]
			if !flag || t < r {
				r = t
				flag = true
			}
		}
		if min > nums[i] {
			min = nums[i]
		}
	}
	return r
}

func main155() {
	fmt.Println(minimumSum([]int{49, 50, 48}))
}
