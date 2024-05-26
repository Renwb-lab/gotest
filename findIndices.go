package main

import "fmt"

// https://leetcode.cn/problems/find-indices-with-index-and-value-difference-ii/description/
func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	// maxIdx: [0, i] 最大的数值下标
	// minIdx: [0, i] 最小的数值下标
	maxIdx, minIdx := 0, 0
	for j := indexDifference; j < len(nums); j += 1 {
		i := j - indexDifference
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		} else if nums[i] < nums[minIdx] {
			minIdx = i
		}
		if nums[maxIdx]-nums[j] >= valueDifference {
			return []int{maxIdx, j}
		}
		if nums[j]-nums[minIdx] >= valueDifference {
			return []int{minIdx, j}
		}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println(findIndices([]int{5, 1, 4, 1}, 2, 3))
}
