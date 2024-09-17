package main

import "fmt"

// https://leetcode.cn/problems/rotate-matrix-lcci/?favorite=xb9lfcwi

// 给定 matrix =
// [
//   [ 5, 1, 9,11],
//   [ 2, 4, 8,10],
//   [13, 3, 6, 7],
//   [15,14,12,16]
// ],

// 原地旋转输入矩阵，使其变为:
// [
//   [15,13, 2, 5],  //1	[0][1] = [i][j]
//   [14, 3, 4, 1],	 //2	[1][3] = 	[j][r-1-i]
//   [12, 6, 8, 9],  //4	[2][0] = 			[r-1-j][i]
//   [16, 7,10,11]	 //3	[3][2] = 		[l-1-i][r-1-j]
// ]

// 原地旋转输入矩阵，使其变为:
// [
//   [15,13, 2, 5],
//   [14, 3, 4, 1],
//   [12, 6, 8, 9],
//   [16, 7,10,11]
// ]

// [
//   [16, 7,10,11],
//   [12, 6, 8, 9],
//   [14, 3, 4, 1],
//   [15,13, 2, 5]
// ]

func rotateV2(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func rotate(matrix [][]int) {
	l, r := len(matrix), len(matrix[0])
	q := l / 2
	for i := 0; i < q; i += 1 {
		for j := i; j < r-1-i; j += 1 {
			t := matrix[i][j]
			matrix[i][j] = matrix[r-1-j][i]
			matrix[r-1-j][i] = matrix[l-1-i][r-1-j]
			matrix[l-1-i][r-1-j] = matrix[j][r-1-i]
			matrix[j][r-1-i] = t
		}
	}
}

func main300() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotateV2(matrix)
	for _, line := range matrix {
		for _, i := range line {
			fmt.Printf("%d\t", i)
		}
		fmt.Println()
	}
}
