package main

import "fmt"

func maxOperations(nums []int) int {
	n := len(nums)
	memo := make([][]int, n)

	helper := func(i, j, target int) int {
		for i := range memo {
			memo[i] = make([]int, n)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}

		var dfs func(int, int) int
		dfs = func(i, j int) int {
			if i >= j {
				return 0
			}
			if memo[i][j] != -1 {
				return memo[i][j]
			}

			ans := 0
			if nums[i]+nums[i+1] == target {
				ans = max(ans, 1+dfs(i+2, j))
			}
			if nums[j-1]+nums[j] == target {
				ans = max(ans, 1+dfs(i, j-2))
			}
			if nums[i]+nums[j] == target {
				ans = max(ans, 1+dfs(i+1, j-1))
			}
			memo[i][j] = ans
			return ans
		}
		return dfs(i, j)
	}

	res := 0
	res = max(res, helper(0, n-1, nums[0]+nums[n-1]))
	res = max(res, helper(0, n-1, nums[0]+nums[1]))
	res = max(res, helper(0, n-1, nums[n-2]+nums[n-1]))
	return res
}

func main263() {
	fmt.Println(maxOperations([]int{3, 2, 1, 2, 3, 4}))
}
