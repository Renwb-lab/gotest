package main

import "math"

func maximumTotalCost(nums []int) int64 {
	// dp[i][0] = 表示截止于i位置，且为0结尾的最大值
	// dp[i][1] = 表示截止于i位置，且为1结尾的最大值
	max64 := func(x, y int64) int64 {
		if x > y {
			return x
		}
		return y
	}
	n := len(nums)
	dp := make([][]int64, n)
	for i := range dp {
		dp[i] = make([]int64, 2)
	}
	dp[0][0] = int64(nums[0])
	dp[0][1] = math.MinInt32

	for i := 1; i < n; i += 1 {
		dp[i][0] = int64(max64(dp[i-1][0], dp[i-1][1])) + int64(nums[i])
		dp[i][1] = dp[i-1][0] - int64(nums[i])
	}
	return max64(dp[n-1][0], dp[n-1][1])
}

func maximumTotalCostV1(nums []int) int64 {
	max64 := func(x, y int64) int64 {
		if x > y {
			return x
		}
		return y
	}
	n := len(nums)
	a, b := int64(nums[0]), int64(math.MinInt64)

	for i := 1; i < n; i += 1 {
		c := int64(max64(a, b)) + int64(nums[i])
		d := a - int64(nums[i])
		a, b = c, d
	}
	return max64(a, b)
}
