package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/xmcbym/

// 两个数组的交集 II

func intersect(nums1 []int, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })
	sort.Slice(nums2, func(i, j int) bool { return nums2[i] < nums2[j] })
	m, n := len(nums1), len(nums2)
	x, y := 0, 0
	ret := []int{}
	for x < m && y < n {
		if nums1[x] < nums2[y] {
			x += 1
		} else if nums1[x] == nums2[y] {
			ret = append(ret, nums1[x])
			x += 1
			y += 1
		} else {
			y += 1
		}
	}
	return ret
}

func main53() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	ret := intersect(nums1, nums2)
	for _, n := range ret {
		fmt.Printf("%d\t", n)
	}
	fmt.Println()
}
