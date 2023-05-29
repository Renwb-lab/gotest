package main

import "fmt"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
// 123. 买卖股票的最佳时机 III

func maxProfit3(prices []int) int {
	// 状态 第一次操作（买入，卖出）第二次操作（买入，卖出）
	// 	buy1, sell1, buy2, sell2 参数的含义：每个状态下的最大收益
	// 选择/转化
	// 	buy1 = max(buy1'[没有操作], -prices[i][买入])
	//  sell1 = max(sell1'[没有操作], buy1'+prices[i][买入])
	//  buy2 = max(buy2'[没有操作], -prices[i][买入])
	//  sell2 = max(sell2'[没有操作], buy2'+prices[i][买入])
	// 可以使用buy1替代buy1'，因为buy1比buy1'多了一次买入的操作，
	// 转移到 sell1时，考虑的是在第 i 天卖出股票的情况，这样在同一天买入并且卖出收益为零 [保障这个正确性]
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	length := len(prices)
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < length; i += 1 {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

func main112() {
	fmt.Println(maxProfit3([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
