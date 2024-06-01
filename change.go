package main

import "fmt"

func change(amount int, coins []int) int {
	n := len(coins)
	helper := make([][]int, n+1)
	for i := 0; i <= n; i += 1 {
		helper[i] = make([]int, amount+1)
	}
	for j := 0; j <= amount; j += 1 {
		helper[0][j] = 0
	}
	for i := 0; i <= n; i += 1 {
		helper[i][0] = 1
	}
	// helper[i][j] 表示使用[0,i)组成金额为j的组合数
	// helper[x][0] 表示组成金额为0的组合数为1
	// helper[0][x] 表示没有coins的情况下组成金额为x的组合数为0
	// helper[i][j] =
	//     helper[i-1][j] // 当前不选
	//     helper[i-1][j-x]
	//     helper[i-1][j-2x]
	for i := 1; i <= n; i += 1 {
		for j := 1; j <= amount; j += 1 {
			x := coins[i-1]
			for k := 0; j-k*x >= 0; k += 1 {
				helper[i][j] += helper[i-1][j-k*x]
			}
		}
	}
	for i := 1; i <= n; i += 1 {
		for j := 1; j <= amount; j += 1 {
			fmt.Printf("%d\t", helper[i][j])
		}
		fmt.Println()
	}

	return helper[n][amount]
}

func change1(amount int, coins []int) int {
	n := len(coins)
	helper := make([][]int, 2)
	for i := 0; i < 2; i += 1 {
		helper[i] = make([]int, amount+1)
	}
	for j := 0; j <= amount; j += 1 {
		helper[0][j] = 0
	}
	for i := 0; i < 2; i += 1 {
		helper[i][0] = 1
	}
	for i := 1; i <= n; i += 1 {
		for j := 1; j <= amount; j += 1 {
			x := coins[i-1]
			helper[i%2][j] = 0
			for k := 0; j-k*x >= 0; k += 1 {
				helper[i%2][j] += helper[(i-1)%2][j-k*x]
			}
			fmt.Printf("%d\t", helper[i%2][j])
		}
		fmt.Println()
	}

	return helper[n%2][amount]
}

func change2(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(change1(5, []int{1, 2, 5}))
	fmt.Println(change2(5, []int{1, 2, 5}))
}
