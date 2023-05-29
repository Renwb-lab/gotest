package main

import "fmt"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
// 188. 买卖股票的最佳时机 IV

func maxProfit4(k int, prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	type Item struct {
		buy  int
		sell int
	}
	helper := make([]Item, k)
	for i := 0; i < k; i += 1 {
		helper[i].buy = -prices[0]
	}

	for i := 1; i < len(prices); i += 1 {
		// 参考两次买卖的场景
		// buy1 = max(buy1, -prices[i])
		// sell1 = max(sell1, buy1+prices[i])
		// buy2 = max(buy2, sell1-prices[i])
		// sell2 = max(sell2, buy2+prices[i])
		helper[0].buy = max(helper[0].buy, -prices[i])
		helper[0].sell = max(helper[0].sell, helper[0].buy+prices[i])
		for j := 1; j < k; j += 1 {
			helper[j].buy = max(helper[j].buy, helper[j-1].sell-prices[i])
			helper[j].sell = max(helper[j].sell, helper[j].buy+prices[i])
		}
	}
	return helper[k-1].sell
}

func main113() {
	fmt.Println(maxProfit4(2, []int{3, 2, 6, 5, 0, 3}))
}
