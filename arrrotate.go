package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xm42hs/

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

func arrRotate(nums []int, k int) {
	if k == 0 {
		return
	}
	k %= len(nums)
	reverse := func(arr []int) {
		for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func main51() {
	nums := []int{-1, -100, 3, 99}
	arrRotate(nums, 2)
	for _, num := range nums {
		fmt.Printf("%d\t", num)
	}
}
