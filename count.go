package main

import "fmt"

// https://leetcode.cn/problems/count-of-integers/

func count(num1 string, num2 string, min_sum int, max_sum int) int {
	const MOD = 1000000007
	const N = 23
	const M = 401

	d := make([][]int, N)
	for i := range d {
		d[i] = make([]int, M)
		for j := range d[i] {
			d[i][j] = -1
		}
	}

	var dfs func(num string, i int, j int, limit bool) int
	dfs = func(num string, i int, j int, limit bool) int {
		if j > max_sum {
			return 0
		}
		if i == -1 {
			if j >= min_sum {
				return 1
			}
			return 0
		}
		if !limit && d[i][j] != -1 {
			return d[i][j]
		}

		res := 0
		var up int
		if limit {
			up = int(num[i] - '0')
		} else {
			up = 9
		}

		for x := 0; x <= up; x++ {
			res = (res + dfs(num, i-1, j+x, limit && x == up)) % MOD
		}

		if !limit {
			d[i][j] = res
		}
		return res
	}

	get := func(num string) int {
		num = reverse(num)
		return dfs(num, len(num)-1, 0, true)
	}
	// 求解 num - 1，先把最后一个非 0 字符减去 1，再把后面的 0 字符变为 9
	sub := func(num string) string {
		i := len(num) - 1
		arr := []byte(num)
		for arr[i] == '0' {
			i--
		}
		arr[i]--
		i++
		for ; i < len(num); i++ {
			arr[i] = '9'
		}
		return string(arr)
	}

	return (get(num2) - get(sub(num1)) + MOD) % MOD
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(count("1", "12", 1, 8))
}
