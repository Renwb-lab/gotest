package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/minimum-number-of-coins-to-be-added/description/

func minimumAddedCoins(coins []int, target int) int {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j]
	})
	// s 表示可以达到的最大值
	s, r := 0, 0
	for _, coin := range coins {
		// 判断是否能够到达coin的值，不能的话则需要添加。 注意是循环！！！
		for s+1 < coin && s < target {
			s += s + 1
			r += 1
		}
		// 再次进行判断
		if s >= target {
			break
		}
		s += coin // 将当前的值加入，扩大范围
	}
	// 当数组处理结束后仍未到达目标值的话，就需要不断扩充
	for s < target {
		s += s + 1
		r += 1
	}
	return r
}

func main275() {
	fmt.Println(minimumAddedCoins([]int{1, 4, 10}, 19))
}
