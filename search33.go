package main

import "fmt"

// https://leetcode.cn/problems/search-in-rotated-sorted-array/
// 33. 搜索旋转排序数组

func search33(nums []int, target int) int {
	// 按照有序数组的方式查询，只是说每次只能保证单边是有序的。
	b, e := 0, len(nums)-1
	for b <= e {
		m := b + (e-b)/2
		if nums[m] == target {
			return m
		}
		if nums[b] <= nums[m] {
			// 保证 nums[b ... m] 是增加的
			if target >= nums[b] && target < nums[m] {
				e = m - 1
			} else {
				b = m + 1
			}
		} else {
			// 保证 nums[m ... e] 是增加的
			if target > nums[m] && target <= nums[e] {
				b = m + 1
			} else {
				e = m - 1
			}
		}
	}
	return -1
}

func main() {
	arr := []int{5, 1, 3}
	fmt.Println(search33(arr, 0))
}
