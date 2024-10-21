package main

import (
	"sort"
)

// https://leetcode.cn/problems/smallest-range-ii/
func smallestRangeII(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	l := len(nums)
	ans := nums[l-1] - nums[0]
	for i := 1; i < l; i += 1 {
		// [0, i-1]加k   --> 最小值为nums[0]+k, 最大值为nums[i-1]+k
		// [i, n-1]减k	 --> 最小值为nums[i]-k, 最大值为nums[n-1]-k
		// 找出最大值和最小值进行计算
		// 依次枚举所有的情况，计算结果
		mx := max(nums[i-1]+k, nums[l-1]-k)
		mn := min(nums[0]+k, nums[i]-k)
		ans = min(ans, mx-mn)
	}
	return ans
}
