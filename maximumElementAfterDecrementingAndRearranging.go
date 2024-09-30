package main

import "sort"

// https://leetcode.cn/problems/maximum-element-after-decreasing-and-rearranging/submissions/569324270/
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	arr[0] = 1
	for i := 0; i < len(arr)-1; i += 1 {
		if arr[i+1]-arr[i] > 1 {
			arr[i+1] = arr[i] + 1
		}
	}
	return arr[len(arr)-1]
}
