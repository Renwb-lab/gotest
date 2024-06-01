package main

import "fmt"

func firstDayBeenInAllRooms(nextVisit []int) int {
	mod := int(1e9 + 7)
	n := len(nextVisit)
	dp := make([]int, n)
	for i := 1; i < n; i += 1 {
		dp[i] = (2*dp[i-1] - dp[nextVisit[i-1]] + 2) % mod
		if dp[i] < 0 {
			dp[i] += mod
		}
	}
	return dp[n-1]
}

func main() {
	fmt.Println(firstDayBeenInAllRooms([]int{0, 0, 2}))
}
