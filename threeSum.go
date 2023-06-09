package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/3sum/
// 15. 三数之和

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	l := len(nums)
	for i := 0; i < l; i += 1 {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		j, k := i+1, l-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j += 1
				for j <= k && nums[j-1] == nums[j] {
					j += 1
				}
				k -= 1
				for j <= k && nums[k+1] == nums[k] {
					k -= 1
				}
			} else if sum < 0 {
				j += 1
			} else {
				k -= 1
			}
		}
	}
	return res
}

func main118() {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	for i := range res {
		for j := range res[i] {
			fmt.Printf("%d\t", res[i][j])
		}
		fmt.Println()
	}
}
