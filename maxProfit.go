package main

import "fmt"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/?company_slug=ubiquanthr

func maxProfit(prices []int) int {
	size, ret := len(prices), 0
	if size == 0 {
		return ret
	}
	preMin := prices[0]
	for i := 1; i < size; i += 1 {
		if prices[i] > preMin {
			if prices[i]-preMin > ret {
				ret = prices[i] - preMin
			}
		} else {
			preMin = prices[i]
		}
	}
	return ret
}

func main32() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
