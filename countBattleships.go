package main

import "fmt"

func countBattleships(board [][]byte) int {
	ans := 0
	for i, line := range board {
		for j := range line {
			if board[i][j] == 'X' {
				if i-1 >= 0 && board[i-1][j] == 'X' {
					continue
				}
				if j-1 >= 0 && board[i][j-1] == 'X' {
					continue
				}
				ans += 1
			}
		}
	}
	return ans
}

func main071024() {
	fmt.Println(countBattleships([][]byte{
		{'X', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}))
}
