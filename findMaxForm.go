package main

import "fmt"

// https://leetcode.cn/problems/ones-and-zeroes/
func findMaxForm(strs []string, m int, n int) int {
	compute := func(str string) (int, int) {
		zeros, ones := 0, 0
		for i := range str {
			if str[i] == '0' {
				zeros += 1
			} else {
				ones += 1
			}
		}
		return zeros, ones
	}

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	l := len(strs)
	dp := make([][][]int, l+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i := range strs {
		zeros, ones := compute(strs[i])
		for j := 0; j <= m; j += 1 {
			for k := 0; k <= n; k += 1 {
				dp[i+1][j][k] = dp[i][j][k]
				if j >= zeros && k >= ones {
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j-zeros][k-ones]+1)
				}
			}
		}
	}
	return dp[l][m][n]
}

func findMaxForm2(strs []string, m int, n int) int {
	// dfs
	compute := func(str string) []int {
		res := make([]int, 2)
		for i := range str {
			if str[i] == '0' {
				res[0] += 1
			} else {
				res[1] += 1
			}
		}
		return res
	}

	l := len(strs)
	mp := make([][]int, l)
	for i := 0; i < l; i += 1 {
		mp[i] = compute(strs[i])
	}
	num := 1 << l
	ans := 0
	for n := 1; n <= num; n += 1 {
		m1, n1, t := 0, 0, 0
		for i := 0; i < l; i += 1 {
			if n&(1<<i) > 0 {
				m1 += mp[i][0]
				n1 += mp[i][1]
				t += 1
			}
			if m1 > m || n1 > n {
				break
			}
		}
		if m1 <= m && n1 <= n && ans < t {
			ans = t
		}
	}
	return ans
}

func findMaxForm3(strs []string, m int, n int) int {
	compute := func(str string) []int {
		res := make([]int, 2)
		for i := range str {
			if str[i] == '0' {
				res[0] += 1
			} else {
				res[1] += 1
			}
		}
		return res
	}

	l := len(strs)
	mp := make([][]int, l)
	for i := 0; i < l; i += 1 {
		mp[i] = compute(strs[i])
	}
	num := 1 << l
	ans := 0
	for n := 1; n <= num; n += 1 {
		m1, n1, t := 0, 0, 0
		for i := 0; i < l; i += 1 {
			if n&(1<<i) > 0 {
				m1 += mp[i][0]
				n1 += mp[i][1]
				t += 1
			}
			if m1 > m || n1 > n {
				break
			}
		}
		if m1 <= m && n1 <= n && ans < t {
			ans = t
		}
	}
	return ans
}

func main() {
	fmt.Println(findMaxForm([]string{
		"10", "0001", "111001", "1", "0",
	}, 3, 4))
}
