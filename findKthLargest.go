package main

import "fmt"

// https://leetcode.cn/problems/kth-largest-element-in-an-array/

func findKthLargest(nums []int, k int) int {
	k -= 1

	middleIndex := func(nums []int, b, e int) int {
		t, idx := nums[e], b-1
		for i := b; i < e; i += 1 {
			if nums[i] >= t {
				idx += 1
				nums[idx], nums[i] = nums[i], nums[idx]
			}
		}
		idx += 1
		nums[idx], nums[e] = nums[e], nums[idx]
		return idx
	}
	b, e := 0, len(nums)-1
	for {
		mid := middleIndex(nums, b, e)
		if mid == k {
			return nums[k]
		} else if mid < k {
			b = mid + 1
		} else {
			e = mid - 1
		}
	}
}

func main24() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
