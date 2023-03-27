package main

import "fmt"

// https://leetcode.cn/problems/eight-queens-lcci/?favorite=xb9lfcwi

func solveNQueens(n int) [][]string {
	helper := make([][]byte, 0)
	ret := make([][]string, 0)
	// 初始化
	for i := 0; i < n; i += 1 {
		line := make([]byte, 0)
		for j := 0; j < n; j += 1 {
			line = append(line, '.')
		}
		helper = append(helper, line)
	}
	isFeasible := func(x, y int) bool {
		// 行不需要判断
		// 列需要判断
		num := 0
		for i := 0; i < n; i += 1 {
			if helper[i][y] == 'Q' {
				num += 1
			}
		}
		if num > 1 {
			return false
		}
		// 正对角线判断
		num = 0
		for x1, y1 := x, y; x1 >= 0 && y1 >= 0; x1, y1 = x1-1, y1-1 {
			if helper[x1][y1] == 'Q' {
				num += 1
			}
		}
		if num > 1 {
			return false
		}
		// 反对角线判断
		num = 0
		for x1, y1 := x, y; x1 >= 0 && y1 < n; x1, y1 = x1-1, y1+1 {
			if helper[x1][y1] == 'Q' {
				num += 1
			}
		}
		return num <= 1
	}
	var dp func(line int)
	dp = func(line int) {
		// 循环结束
		if line == n {
			oneRet := make([]string, 0)
			for i := 0; i < n; i += 1 {
				oneRet = append(oneRet, string(helper[i]))
			}
			ret = append(ret, oneRet)
			return
		}
		for i := 0; i < n; i += 1 {
			helper[line][i] = 'Q'
			if isFeasible(line, i) {
				dp(line + 1)
			}
			helper[line][i] = '.'
		}
	}

	dp(0)
	return ret
}

func main10() {
	ret := solveNQueens(4)
	for _, ss := range ret {
		for _, line := range ss {
			fmt.Println(line)
		}
		fmt.Println("=============")
	}
}
