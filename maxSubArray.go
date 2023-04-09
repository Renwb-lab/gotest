package main

import "fmt"

// https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/?company_slug=ubiquanthr

// 剑指 Offer 42. 连续子数组的最大和

// 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

// 要求时间复杂度为O(n)。

// 示例1:

// 输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

func maxSubArray(nums []int) int {
	ret, l := nums[0], len(nums)
	pre, cur := nums[0], 0
	for i := 1; i < l; i += 1 {
		cur = nums[i]
		if pre > 0 {
			cur += pre
		}
		if cur > ret {
			ret = cur
		}
		pre = cur
	}
	return ret
}

func main36() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
