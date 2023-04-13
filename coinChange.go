package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/x2echt/
// 零钱兑换

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	helper := make([]int, amount+1)
	helper[0] = 1
	for i := 1; i <= amount; i += 1 {
		t := math.MaxInt32
		for _, v := range coins {
			if i > v {
				if helper[i-v]+helper[v] < t {
					t = helper[i-v] + helper[v]
				}
			} else if i == v {
				t = 1
			}
		}
		helper[i] = t
	}
	if helper[amount] == math.MaxInt32 {
		return -1
	}
	return helper[amount]
}

func main60() {
	nums := []int{2}
	fmt.Println(coinChange(nums, 3))
}
