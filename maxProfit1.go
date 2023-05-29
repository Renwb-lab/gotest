package main

import (
	"fmt"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
// 121. 买卖股票的最佳时机

func maxProfit1(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	ret, preMin := 0, prices[0]
	for i := 1; i < length; i += 1 {
		// 判断以当前节点进行卖出的话，获取的收益
		if prices[i]-preMin > ret {
			ret = prices[i] - preMin
		}
		// 更新最小的数值
		if prices[i] < preMin {
			preMin = prices[i]
		}
	}
	return ret
}

func main110() {
	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
}
