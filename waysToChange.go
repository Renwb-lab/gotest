package main

import "fmt"

// https://leetcode.cn/problems/coin-lcci/description/?favorite=xb9lfcwi

func waysToChange0(n int) int {
	// i = 0.....n
	// j = 0, 1, 2, 3
	// helper[i][j]: 表示使用coins[:j]个硬币可以i元的可能数量
	// helper[0][0] = helper[0][1] = helper[0][2] = helper[0][3] = 1
	// helper[i][j] =
	// 				   helper[i][j-1] 不使用第j个硬币
	// 				   helper[i][j-1] + helper[i-coins[j]][j] 使用第j个硬币
	// return helper[n][3] % 1000000007
	// 初始化
	helper := make([][]int, n+1, n+1)
	for i := 0; i <= n; i += 1 {
		helper[i] = make([]int, 4, 4)
	}
	// 初始状态
	for j := 0; j < 4; j += 1 {
		helper[0][j] = 1
	}

	coins := []int{1, 5, 10, 25}
	for j := 0; j < 4; j += 1 {
		for i := 1; i <= n; i += 1 {
			if j >= 1 {
				helper[i][j] = helper[i][j-1]
			}
			if i-coins[j] >= 0 {
				helper[i][j] += helper[i-coins[j]][j]
			}
		}
	}

	// 0 1 2 3
	// 0 1 2 3
	// 0 1 2 3
	// 0 1 2 3
	// 0 1 2 3
	// 0 1 2 3
	// 0 1 2 3
	// 0 1 2 3

	return helper[n][3] % 1000000007
}

func waysToChange(n int) int {
	// i = 0.....n
	// j = 0, 1, 2, 3
	// helper[i][j]: 表示使用coins[:j]个硬币可以i元的可能数量
	// helper[0][0] = helper[0][1] = helper[0][2] = helper[0][3] = 1
	// helper[i][j] =
	// 				   helper[i][j-1] 不使用第j个硬币
	// 				   helper[i][j-1] + helper[i-coins[j]][j] 使用第j个硬币
	// return helper[n][3] % 1000000007
	// 初始化
	helper := make([]int, n+1, n+1)
	// 初始状态
	helper[0] = 1

	coins := []int{1, 5, 10, 25}
	for j := 0; j < 4; j += 1 {
		for i := 1; i <= n; i += 1 {
			if i-coins[j] >= 0 {
				helper[i] += helper[i-coins[j]]
			}
		}
	}

	// 0 1 2 3
	// 0 1 2 3
	// 0 1 2 3
	// 0 1 2 3
	// 0 1 2 3
	// 0 1 2 3
	// 0 1 2 3
	// 0 1 2 3

	return helper[n] % 1000000007
}

func waysToChange1(n int) int {
	dp := make([][]int, 0)
	for i := 0; i < 4; i += 1 {
		dp = append(dp, make([]int, n+1))
	}
	for i := 0; i < 4; i += 1 {
		dp[i][0] = 1
	}
	for j := 0; j < n+1; j += 1 {
		dp[0][j] = 1
	}
	// dp[i][j] 表示有coins[0:i] 组成的j的可能数
	// f(i,v)=f(i−1,v)+f(i−1,v−c i​)+f(i−1,v−2c i​)⋯f(i−1,v−kc i​ )
	// f(i,v−c i​ )=f(i−1,v−c i​ )+f(i−1,v−2c i​ )+f(i−1,v−3c i​ )⋯f(i−1,v−kc i​ )
	// f(i, v) = f(i-1,v) + f(i, v-ci)
	// i: 0 -> 4
	// v: 1 -> n
	coins := []int{1, 5, 10, 25}
	for j := 1; j <= n; j++ {
		for i := 1; i < 4; i++ { // 优化前，i,j 的循环顺序是可以调整的
			dp[i][j] = dp[i-1][j]
			if j-coins[i] >= 0 {
				dp[i][j] += dp[i][j-coins[i]]
			}
		}
	}
	return dp[3][n] % 1000000007
}

func waysToChange2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	coins := []int{1, 5, 10, 25}
	for i := 0; i < 4; i++ {
		for j := 1; j <= n; j++ { // 优化后，i,j 的循环顺序是不可以调整的。
			//  因为i是严格递增的。
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[n] % 1000000007
}

func main() {
	fmt.Println(waysToChange(64))

	fmt.Println(waysToChange1(5))
	fmt.Println(waysToChange1(10))

	fmt.Println(waysToChange2(5))
	fmt.Println(waysToChange2(10))
}
