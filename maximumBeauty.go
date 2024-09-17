package main

import (
	"fmt"
	"sort"
)

// 2779. 数组的最大美丽值
func maximumBeauty(nums []int, k int) (ans int) {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	left := 0
	for right, x := range nums {
		for x-nums[left] > k*2 {
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func main249() {
	fmt.Println(maximumBeauty([]int{4, 6, 1, 2}, 2))
}
