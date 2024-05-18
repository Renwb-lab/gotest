package main

import "fmt"

//  https://leetcode.cn/leetbook/read/top-interview-questions/xmup75/

//  鸡蛋掉落

// 给你 k 枚相同的鸡蛋，并可以使用一栋从第 1 层到第 n 层共有 n 层楼的建筑。
// 已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都会碎，从 f 楼层或比它低的楼层落下的鸡蛋都不会破。
// 每次操作，你可以取一枚没有碎的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎，则可以在之后的操作中 重复使用 这枚鸡蛋。
// 请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？

func superEggDropV1(k int, n int) int {
	mp := make([][]int, k+1, k+1)
	for i := 0; i <= k; i += 1 {
		mp[i] = make([]int, n+1, n+1)
	}

	var dp func(k int, n int) int
	dp = func(k int, n int) int {
		if mp[k][n] != 0 {
			return mp[k][n]
		}
		if n == 0 {
			mp[k][n] = 0
			return 0
		}
		if n == 1 { // 只有一层，那只要尝试1次
			mp[k][n] = 1
			return 1
		}
		if k == 1 { // 只有一个鸡蛋，那就要尝试n次
			mp[k][n] = n
			return n
		}
		res := n
		for i := 1; i <= n; i += 1 {
			t := 0
			t1 := dp(k, i-1)   // 没有破
			t2 := dp(k-1, n-i) // 破了
			if t1 > t2 {
				t = t1 + 1
			} else {
				t = t2 + 1
			}
			if res > t {
				res = t
			}
		}
		mp[k][n] = res
		return res
	}

	return dp(k, n)
}

func superEggDrop(k int, n int) int {
	// dp[k][n] = m, 表示 k 个鸡蛋 n 层，最坏情况下至少需要尝试m次
	// dp[k][m] = n, 表示 k 个鸡蛋，尝试m次，最好情况下能够到达的楼层为n
	mp := make([][]int, k+1, k+1)
	for i := 0; i <= k; i += 1 {
		mp[i] = make([]int, n+1, n+1)
	}

	var dp func(k int, n int) int
	dp = func(k int, n int) int {
		if mp[k][n] != 0 {
			return mp[k][n]
		}
		if n == 0 {
			mp[k][n] = 0
			return 0
		}
		if n == 1 { // 只有一层，那只要尝试1次
			mp[k][n] = 1
			return 1
		}
		if k == 1 { // 只有一个鸡蛋，那就要尝试n次
			mp[k][n] = n
			return n
		}
		res := n
		l, r := 1, n
		for l <= r {
			m := l + (r-l)/2
			t1 := dp(k, m-1)   // 没有破
			t2 := dp(k-1, n-m) // 破了

			t := 0
			if t1 > t2 {
				t = t1
				r = m - 1
			} else {
				t = t2
				l = m + 1
			}
			if res > t+1 {
				res = t + 1
			}
		}
		mp[k][n] = res
		return res
	}

	return dp(k, n)
}

func superEggDropV2(k int, n int) int {
	// dp[k][m] = x 表示k个鸡蛋尝试m次，最好的情况下，能够找到的最大高度为x
	// dp[1][7] = 7 表示1个鸡蛋尝试7次，最好的情况下，能够找到的最大高度为7

	// dp[k][m] = dp[k][m-1] + dp[k-1][m-1] + 1
	dp := make([][]int, k+1)
	for i := 0; i <= k; i += 1 {
		dp[i] = make([]int, n+1)
	}
	// 不管多少个鸡蛋，如果只扔一次的话，那么最大高度就是1
	for i := 1; i <= k; i += 1 {
		dp[i][1] = 1
	}
	// 如果只有一个鸡蛋的话，有几次尝试次数，最大高度就是几
	for j := 1; j <= n; j += 1 {
		dp[1][j] = j
	}

	for j := 1; j <= n; j += 1 {
		for i := 1; i <= k; i += 1 {
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
		}
		if dp[k][j] >= n {
			return j
		}
	}

	return 0
}

func main() {
	fmt.Println(superEggDrop(2, 6))
}
