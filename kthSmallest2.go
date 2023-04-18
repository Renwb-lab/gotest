package main

// https://leetcode.cn/leetbook/read/top-interview-questions/xaicbc/
// 有序矩阵中第K小的元素

func kthSmallestMatrix(matrix [][]int, k int) int {
	// [1, 5, 10],
	// [10,11,13],
	// [12,13,15]
	n := len(matrix)
	check := func(val int) int {
		num := 0
		for i, j := n-1, 0; i >= 0 && j < n; {
			if matrix[i][j] <= val {
				num += i + 1
				j += 1
			} else {
				i -= 1
			}
		}
		return num
	}
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)>>1
		num := check(mid)
		if num < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
