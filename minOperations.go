package main

// https://leetcode.cn/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-ii/
func minOperations(nums []int) int {
	times := 0
	isOne := func(v int) bool {
		if times%2 == 0 {
			return v == 1
		} else {
			return v == 0
		}
	}
	for _, v := range nums {
		if !isOne(v) {
			times += 1
		}
	}
	return times
}
