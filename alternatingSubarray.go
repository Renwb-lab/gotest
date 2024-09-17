package main

import "fmt"

// https://leetcode.cn/problems/longest-alternating-subarray/submissions/497863523/
func alternatingSubarray(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	res, x, y := 1, 1, 1
	if nums[1]-nums[0] == 1 {
		y += 1
	}
	res = max(res, y)
	for i := 2; i < len(nums); i += 1 {
		if y > 1 && nums[i] == nums[i-2] {
			x = y + 1
		} else {
			if nums[i]-nums[i-1] == 1 {
				x = 2
			} else {
				x = 1
			}
		}
		res = max(res, x)
		x, y = y, x
	}
	if res == 1 {
		return -1
	}
	return res
}

func main07109() {
	fmt.Println(alternatingSubarray([]int{2, 3, 4, 3, 4}))
}
