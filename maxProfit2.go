package main

import "fmt"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
// 122. 买卖股票的最佳时机 II

func maxProfit2(prices []int) int {
	ret := 0
	for i := 1; i < len(prices); i += 1 {
		// 贪心算法，只要当天比前一天大的话，就进行一次操作
		if prices[i] > prices[i-1] {
			ret += prices[i] - prices[i-1]
		}
	}
	return ret
}

func main111() {
	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
}
