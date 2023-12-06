package main

import "fmt"

// https://leetcode.cn/problems/minimum-increment-operations-to-make-array-beautiful/description/

func minIncrementOperations(nums []int, k int) int64 {
	// dp[i] 表示修改第 i 项并使前 i 项变为美丽数组的最小修改次数
	// 重点：更改第i项
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	helper := [3]int{}
	for i := 0; i < 3; i += 1 {
		helper[i] = max(0, k-nums[i])
	}
	for i := 3; i < len(nums); i += 1 {
		t := min(min(helper[0], helper[1]), helper[2]) + max(0, k-nums[i])
		helper[0], helper[1], helper[2] = helper[1], helper[2], t
	}
	return int64(min(min(helper[0], helper[1]), helper[2]))
}

func main158() {
	fmt.Println(minIncrementOperations([]int{2, 3, 0, 0, 2}, 4))
}
