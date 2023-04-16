package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/wiggle-sort-ii/description/
// 摆动排序 II

func wiggleSort(nums []int) {
	arr := append([]int{}, nums...)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	x := (len(arr) + 1) / 2
	for i, j, k := 0, x-1, len(arr)-1; i < len(arr); i, j, k = i+2, j-1, k-1 {
		nums[i] = arr[j]
		if i+1 < len(arr) {
			nums[i+1] = arr[k]
		}
	}
}

func main63() {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	for _, v := range nums {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}
