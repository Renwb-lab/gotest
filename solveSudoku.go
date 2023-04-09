package main

import "fmt"

// https://leetcode.cn/problems/sudoku-solver/?company_slug=ubiquanthr

//  编写一个程序，通过填充空格来解决数独问题。
//  数独的解法需 遵循如下规则：
//
//  数字 1-9 在每一行只能出现一次。
//  数字 1-9 在每一列只能出现一次。
//  数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//  数独部分空格内已填入了数字，空白格用 '.' 表示。

func solveSudoku(board [][]byte) {
	n := 9

	feasible := func(x, y int) bool {
		// line
		c, t := board[x][y], 0
		for _, b := range board[x] {
			if b == c {
				t += 1
			}
		}
		if t > 1 {
			return false
		}
		// row
		t = 0
		for i := 0; i < n; i += 1 {
			if board[i][y] == c {
				t += 1
			}
		}
		if t > 1 {
			return false
		}
		// square
		x1, y1 := x/3, y/3
		t = 0
		for i := 0; i < 3; i += 1 {
			for j := 0; j < 3; j += 1 {
				if board[x1*3+i][y1*3+j] == c {
					t += 1
				}
			}
		}
		return t == 1
	}

	var dp func(x, y int) bool
	dp = func(x, y int) bool {
		if x == n {
			return true
		}
		if board[x][y] != '.' {
			if y+1 < n {
				return dp(x, y+1)
			}
			return dp(x+1, 0)
		}
		for i := 1; i <= n; i += 1 {
			board[x][y] = byte('0' + i)
			if feasible(x, y) {
				if y+1 < n {
					if dp(x, y+1) {
						return true
					}
				} else {
					if dp(x+1, 0) {
						return true
					}
				}
			}
			board[x][y] = '.'
		}
		return false
	}
	dp(0, 0)
}

func main29() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	for _, line := range board {
		for _, b := range line {
			fmt.Printf("%c\t", b)
		}
		fmt.Println()
	}
}
