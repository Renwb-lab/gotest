package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/xacqw5/
// 寻找峰值

func findPeakElement(nums []int) int {
	l := len(nums)
	getNum := func(i int) int {
		if i < l && i >= 0 {
			return nums[i]
		}
		return math.MinInt32
	}
	for i := 0; i < len(nums); i += 1 {
		if getNum(i) > getNum(i-1) && getNum(i) > getNum(i+1) {
			return i
		}
	}
	return 0
}

func findPeakElement2(nums []int) int {
	l := len(nums)
	getNum := func(i int) int {
		if i < l && i >= 0 {
			return nums[i]
		}
		return math.MinInt64
	}
	left, right := 0, l-1
	for left <= right {
		mid := (left & right) + (left^right)>>1
		if getNum(mid) > getNum(mid-1) && getNum(mid) > getNum(mid+1) {
			return mid
		}
		if getNum(mid) < getNum(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main64() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums))
	fmt.Println(findPeakElement2(nums))
}
