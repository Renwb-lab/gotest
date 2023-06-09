package main

import (
	"fmt"
)

// https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock/
// 2110. 股票平滑下跌阶段的数目

func getDescentPeriods(prices []int) int64 {
	n := len(prices)
	preNum, seqNum := prices[0], 1
	ret := 0
	for i := 1; i < n; i += 1 {
		if preNum-prices[i] == 1 {
			seqNum += 1
		} else {
			ret += seqNum * (seqNum + 1) / 2
			seqNum = 1
		}
		preNum = prices[i]
	}
	// 非常重要
	ret += seqNum * (seqNum + 1) / 2
	return int64(ret)
}

func main116() {
	fmt.Println(getDescentPeriods([]int{8, 6, 7, 7}))
}
