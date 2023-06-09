package main

import "fmt"

// https://leetcode.cn/problems/candy/?company_slug=didi
// 135. 分发糖果
func candy(ratings []int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(ratings)
	left := make([]int, n)
	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}
	return
}

func main125() {
	fmt.Println(candy([]int{1, 2, 87, 87, 87, 2, 1}))
}
