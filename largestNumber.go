package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/xa1401/
// 最大数

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		str1 := fmt.Sprintf("%d%d", nums[i], nums[j])
		str2 := fmt.Sprintf("%d%d", nums[j], nums[i])
		return str1 > str2
	})
	ret := ""
	i := 0
	for ; i < len(nums); i += 1 {
		if nums[i] != 0 {
			break
		}
	}
	for ; i < len(nums); i += 1 {
		ret += fmt.Sprint(nums[i])
	}
	if len(ret) == 0 {
		return "0"
	}
	return ret
}

func main62() {
	nums := []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums))
}
