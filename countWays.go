package main

import (
	"fmt"
	"sort"
)

// 100040. 让所有学生保持开心的分组方法数
// https://leetcode.cn/contest/weekly-contest-363/problems/happy-students/

func countWays(nums []int) int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	l := len(nums)
	check := func(s int) bool {
		// 左边的都选
		if s == 0 {
			return nums[0] > 0
		}
		if s == l {
			return s > nums[s-1]
		}
		if s > nums[s-1] {
			if s < nums[s] {
				return true
			}
		}
		return false
	}

	r := 0
	for i := 0; i <= l; i += 1 {
		if check(i) {
			r += 1
		}
	}
	return r
}

func main134() {
	fmt.Print(countWays([]int{1, 1}))
}
