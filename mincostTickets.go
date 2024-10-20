package main

// https://leetcode.cn/problems/minimum-cost-for-tickets/
func mincostTickets(days []int, costs []int) int {
	// dp[i]表示到days[:0)所需的最小值
	l := len(days)
	dp := make([]int, l)
	for i := range dp {
		dp[i] = 0
	}
	if costs[1] > costs[2] {
		costs[1] = costs[2]
	}
	if costs[0] > costs[1] {
		costs[0] = costs[1]
	}
	for i := 0; i < l; i += 1 {
		// 选择一天的
		cur := costs[0]
		if i-1 >= 0 {
			cur += dp[i-1]
		}
		// 选择七天的
		for k := i - 1; k >= 0 && days[i]-days[k]+1 <= 7; k -= 1 {
			if k-1 >= 0 {
				cur = min(cur, dp[k-1]+costs[1])
			} else {
				cur = min(cur, costs[1])
			}
		}
		// 选择三十天的
		for k := i - 1; k >= 0 && days[i]-days[k]+1 <= 30; k -= 1 {
			if k-1 >= 0 {
				cur = min(cur, dp[k-1]+costs[2])
			} else {
				cur = min(cur, costs[2])
			}
		}
		dp[i] = cur
	}
	return dp[l-1]
}

func mincostTickets2(days []int, costs []int) int {
	memo := [366]int{}
	dayM := map[int]bool{}
	for _, d := range days {
		dayM[d] = true
	}

	var dp func(day int) int
	dp = func(day int) int {
		if day > 365 {
			return 0
		}
		if memo[day] > 0 {
			return memo[day]
		}
		if dayM[day] {
			memo[day] = min(min(dp(day+1)+costs[0], dp(day+7)+costs[1]), dp(day+30)+costs[2])
		} else {
			memo[day] = dp(day + 1)
		}
		return memo[day]
	}
	return dp(1)
}
