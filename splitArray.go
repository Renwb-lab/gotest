package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/split-array-largest-sum
func splitArray2(nums []int, m int) int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	n := len(nums)
	// helper[i][j] 表示nums[:i)数据, 取j个数据的结果
	// 动态规划一定要注意连续性
	helper := make([][]int, n+1)
	for i := 0; i <= n; i += 1 {
		helper[i] = make([]int, m+1)
		for j := 0; j <= m; j += 1 {
			helper[i][j] = math.MaxInt
		}
	}
	// 前缀和数组
	preSum := make([]int, n+1)
	for i := 0; i < n; i += 1 {
		preSum[i+1] = preSum[i] + nums[i]
	}

	helper[0][0] = 0
	for i := 1; i <= n; i += 1 {
		for j := 1; j <= min(i, m); j += 1 {
			for k := 0; k < i; k += 1 {
				helper[i][j] = min(helper[i][j],
					max(helper[k][j-1], preSum[i]-preSum[k]))
			}
		}
	}

	return helper[n][m]
}

func splitArray(nums []int, m int) int {
	check := func(x int) int {
		cnt, sum := 1, 0
		for _, v := range nums {
			sum += v
			if sum > x {
				cnt += 1
				sum = v
			}
		}
		return cnt
	}
	left, right := 0, 0
	for _, v := range nums {
		if left < v {
			left = v
		}
		right += v
	}
	right += 1
	for left < right { // [left, right)
		mid := left + (right-left)>>1
		// 分成cnt段的最大值为mid
		// 分的段越多，则意味着最大值越小。
		cnt := check(mid)
		if cnt <= m {
			right = mid // [left, mid)
		} else {
			left = mid + 1 // [mid+1, right)
		}
	}
	return left
}

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
}
