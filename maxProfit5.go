package main

import "fmt"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// 309. 最佳买卖股票时机含冷冻期

func maxProfit5(prices []int) int {
	// 状态 第一次操作（买入，卖出）, 冷冻，第二次操作（买入，卖出）
	// buy1, sell1, freezing, buy2, sell2
	// 选择
	// 状态
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// f[i][0]: 手上持有股票的最大收益
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	n := len(prices)
	helper := make([][3]int, n)
	helper[0][0] = -prices[0]
	for i := 1; i < len(prices); i += 1 {
		helper[i][0] = max(helper[i-1][0], helper[i-1][2]-prices[i])
		helper[i][1] = helper[i-1][0] + prices[i]
		helper[i][2] = max(helper[i-1][1], helper[i-1][2])
	}

	return max(helper[n-1][1], helper[n-1][2])
}

func maxProfit51(prices []int) int {
	// 状态 第一次操作（买入，卖出）, 冷冻，第二次操作（买入，卖出）
	// buy1, sell1, freezing, buy2, sell2
	// 选择
	// 状态
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// f[i][0]: 手上持有股票的最大收益
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	helper := make([]int, 3)
	helper[0] = -prices[0]
	for i := 1; i < len(prices); i += 1 {
		preHelper2 := helper[1]
		helper[0] = max(helper[0], helper[2]-prices[i])
		helper[1] = helper[0] + prices[i]
		helper[2] = max(preHelper2, helper[2])
	}

	return max(helper[1], helper[2])
}

func maxProfit52(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	buy, sell, cooldown := -100001, 0, 0
	// 常规理解：
	// 77行 preSell := sell 和 81行 cooldown = preSell 两行可知： cooldown 表示上一次卖出的最大收益
	// 这里理解【cooldown 表示上上一次卖出的最大收益】的原因是：
	// 外层的for循环在作怪，因为这里从i变成i+1时，sell 本身就表示上一次卖出的最大收益。
	// 在加上preSell的原因，所以cooldown 表示上上一次卖出的最大收益
	// 具体请参考maxProfit53
	for i := 0; i < len(prices); i += 1 {
		preSell := sell
		buy = max(buy, cooldown-prices[i])
		sell = max(sell, buy+prices[i])
		fmt.Printf("i: %d, buy: %d, sell: %d, preSell: %d, cooldown: %d\n", i, buy, sell, preSell, cooldown)
		cooldown = preSell
	}
	return sell
}

func maxProfit53(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(prices)
	buy, sell, cooldown := make([]int, n), make([]int, n), make([]int, n)
	buy[0] = -prices[0]
	for i := 1; i < len(prices); i += 1 {
		buy[i] = max(buy[i-1], cooldown[i-1]-prices[i])
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
		cooldown[i] = sell[i-1]
	}
	return max(sell[n-1], cooldown[n-1])
}

func maxProfit54(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(prices)
	buy, sell := make([]int, n), make([]int, n)
	buy[0], sell[0] = -prices[0], 0
	if n >= 2 {
		if prices[1] > prices[0] {
			buy[1], sell[1] = buy[0], prices[1]-prices[0]
		} else {
			buy[1], sell[1] = -prices[1], 0
		}
	}
	for i := 1; i < len(prices); i += 1 {
		buy[i] = max(buy[i-1], sell[i-2]-prices[i])
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
	}
	return sell[n-1]
}

func main114() {
	fmt.Println(maxProfit52([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit52([]int{1, 2}))
}
