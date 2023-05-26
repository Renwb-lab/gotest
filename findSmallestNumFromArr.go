package main

import "fmt"

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/
// 153. 寻找旋转排序数组中的最小值
func findSmallestNumFromArr(arr []int) int {
	b, e := 0, len(arr)-1
	for b < e {
		m := b + (e-b)/2
		if arr[m] > arr[e] {
			b = m + 1
		} else {
			e = m
		}
	}
	return arr[b]
}

func main106() {
	fmt.Println(findSmallestNumFromArr([]int{4, 5, 6, 1, 2, 3}))
	fmt.Println(findSmallestNumFromArr([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(findSmallestNumFromArr([]int{6, 1, 2, 3, 4, 5}))
	fmt.Println(findSmallestNumFromArr([]int{2, 3, 4, 5, 6, 1}))
}
