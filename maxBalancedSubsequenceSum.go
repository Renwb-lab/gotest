package main

import "fmt"

func maxBalancedSubsequenceSum(nums []int) int64 {
	l := len(nums)
	// 以 idx 结尾的最后元素和
	helper := make([]int64, l)
	r := int64(nums[0])
	for i := 0; i < l; i += 1 {
		helper[i] = int64(nums[i])
		if r < int64(nums[i]) {
			r = int64(nums[i])
		}
	}
	for i := 0; i < l; i += 1 {
		for j := i - 1; j >= 0; j -= 1 {
			if nums[i]-nums[j] >= i-j {
				if helper[i] < helper[j]+int64(nums[i]) {
					helper[i] = helper[j] + int64(nums[i])
				}
			}
		}
		if r < helper[i] {
			r = helper[i]
		}
	}
	return r
}

func main() {
	fmt.Println(maxBalancedSubsequenceSum([]int{3, 3, 5, 6}))
}
