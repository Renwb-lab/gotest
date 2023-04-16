package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/x27sii/
// 缺失数字

func missingNumber(nums []int) int {
	n := len(nums)
	pre, sum := (n*(n+1))>>1, 0
	for _, v := range nums {
		sum += v
	}
	return pre - sum
}

func main78() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
}
