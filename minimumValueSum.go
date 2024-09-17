package main

import "fmt"

func minimumValueSum(nums []int, andValues []int) int {
	n, m := len(nums), len(andValues)
	res := -1
	var dfs func(x int, y int, sum int)
	dfs = func(x int, y int, sum int) {
		if x == n && y == m {
			if res == -1 || res > sum {
				res = sum
			}
			return
		}
		if x >= n || y >= m {
			return
		}
		andNum := 0
		for i := x; i < n; i += 1 {
			if i == x {
				andNum = nums[x]
			} else {
				andNum = andNum & nums[i]
			}
			if andNum == andValues[y] {
				dfs(i+1, y+1, sum+nums[i])
			}
		}

	}
	dfs(0, 0, 0)
	return res
}

func minimumValueSum2(nums []int, andValues []int) int {
	n, m := len(nums), len(andValues)
	var dfs func(x int, y int) int
	dfs = func(x int, y int) int {
		if x == n && y == m {
			return 0
		}
		if x >= n || y >= m {
			return -1
		}
		andNum := 0
		res := -1
		for i := x; i < n; i += 1 {
			if i == x {
				andNum = nums[x]
			} else {
				andNum = andNum & nums[i]
			}
			if andNum == andValues[y] {
				t := dfs(i+1, y+1)
				if t != -1 {
					t += nums[i]
					if res == -1 || res > t {
						res = t
					}
				}
			}
		}
		return res
	}
	return dfs(0, 0)
}

func main270() {
	fmt.Println(minimumValueSum2([]int{4, 4}, []int{4}))
}
