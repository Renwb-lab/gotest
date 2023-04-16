package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xakn6r/
// 四数相加 II

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n := len(nums1)
	hashTable := make(map[int]int, n*n) // 记录已经出现的两数之和以及数量
	for _, a := range nums1 {
		for _, b := range nums2 {
			hashTable[a+b] += 1
		}
	}
	ret := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			// 因为知道四数相加为0，所以在这里直接判断-c-d是否出现在上述的hashTable中
			// 如果出现的话，直接相加对应的数量就可以了。
			// 所以，不需要再次遍历了。
			if r, ok := hashTable[-c-d]; ok {
				ret += r
			}
		}
	}

	return ret
}

func main77() {
	fmt.Println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}
