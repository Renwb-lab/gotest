package main

import (
	"fmt"
	"sort"
)

func numSubseq(nums []int, target int) int {
	mod := 1000000007
	maxN := 100005
	// 只是提前计算了pow(2, x)
	helper := make([]int, maxN, maxN)
	helper[0] = 1
	for i := 1; i < len(nums); i += 1 {
		helper[i] = helper[i-1] * 2 % mod
	}
	// 排序
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	// 计算
	// 1
	// 1,2,
	// 1,5
	// 1,2,5
	ans := 0
	for i := 0; i < len(nums); i += 1 {
		n := target - nums[i]
		pos := sort.Search(len(nums), func(x int) bool {
			return nums[x] >= n+1
		})
		if pos-1 >= i {
			// 包含最小值，[i, pos)区间内满足条件的组合数
			ans = (ans + helper[pos-1-i]) % mod
		}
	}
	return ans
}

func main286() {
	fmt.Println(numSubseq([]int{2, 3, 3, 4, 6, 7}, 12))
}
