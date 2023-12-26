package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	flag := false
	res := -1
	var dfs func(i, j int, sum int)
	dfs = func(i, j, sum int) {
		if i >= l {
			if !flag || sum < res {
				res = sum
				flag = true
			}
			return
		}
		sum += triangle[i][j]
		dfs(i+1, j, sum)
		if j < i+1 {
			dfs(i+1, j+1, sum)
		}
	}
	dfs(0, 0, 0)
	return res
}

func minimumTotal2(triangle [][]int) int {
	l := len(triangle)
	helper := make([][]int, l, l)
	for i := 0; i < l; i += 1 {
		helper[i] = make([]int, i+1, i+1)
	}
	helper[0][0] = triangle[0][0]
	for i := 1; i < l; i += 1 {
		helper[i][0] = helper[i-1][0] + triangle[i][0]
		// i = 2, j = 1
		for j := 1; j < i; j += 1 {
			if helper[i-1][j] < helper[i-1][j-1] {
				helper[i][j] = helper[i-1][j] + triangle[i][j]
			} else {
				helper[i][j] = helper[i-1][j-1] + triangle[i][j]
			}
		}
		helper[i][i] = helper[i-1][i-1] + triangle[i][i]
	}
	ans := helper[l-1][0]
	for i := 1; i < l; i += 1 {
		if ans > helper[l-1][i] {
			ans = helper[l-1][i]
		}
	}
	return ans
}

func minimumTotal3(triangle [][]int) int {
	l := len(triangle)
	helper := make([]int, l, l)
	helper[0] = triangle[0][0]
	for i := 1; i < l; i += 1 {
		helper[i] = helper[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j -= 1 {
			if helper[j] < helper[j-1] {
				helper[j] = helper[j] + triangle[i][j]
			} else {
				helper[j] = helper[j-1] + triangle[i][j]
			}
		}
		helper[0] = helper[0] + triangle[i][0]
	}
	ans := helper[0]
	for i := 1; i < l; i += 1 {
		if ans > helper[i] {
			ans = helper[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumTotal3([][]int{
		{-1},
		{2, 3},
		{1, -1, -3},
	}))
}
