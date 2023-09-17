package main

import "fmt"

// https://leetcode.cn/problems/maximum-size-subarray-sum-equals-k/description/
// 325. 和等于 k 的最长子数组长度

func maxSubArrayLen(nums []int, k int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	helper := map[int]int{0: -1}
	pre, res := 0, 0
	for i, n := range nums {
		current := pre + n
		if _, ok := helper[current]; !ok {
			helper[current] = i
		}
		if idx, ok := helper[current-k]; ok {
			res = max(res, i-idx)
		}
		pre = current
	}
	return res
}

func main120() {
	fmt.Println(maxSubArrayLen([]int{-2, -1, 2, 1}, 1))
}
