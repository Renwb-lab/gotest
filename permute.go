package main

import "fmt"

// https://leetcode.cn/problems/permutations/?company_slug=ubiquanthr

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	l := len(nums)
	var dp func(idx int)
	dp = func(idx int) {
		if idx == l {
			ret = append(ret, append([]int{}, nums...))
			return
		}
		for i := idx; i < l; i += 1 {
			nums[i], nums[idx] = nums[idx], nums[i]
			dp(idx + 1)
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
	dp(0)
	return ret
}

func main30() {
	nums := []int{1, 2, 3}
	ret := permute(nums)
	for _, line := range ret {
		for _, v := range line {
			fmt.Printf("%d\t", v)
		}
		fmt.Println()
	}
}
