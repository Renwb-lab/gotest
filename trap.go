package main

import "fmt"

// https://leetcode.cn/problems/trapping-rain-water/
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

func trap(height []int) int {
	left, right := 1, len(height)-2
	leftMax, rightMax, ret := 0, 0, 0
	for left <= right {
		if height[left-1] < height[right+1] {
			if leftMax < height[left-1] {
				leftMax = height[left-1]
			}
			if leftMax > height[left] {
				ret += leftMax - height[left]
			}
			left += 1
		} else {
			if rightMax < height[right+1] {
				rightMax = height[right+1]
			}
			if rightMax > height[right] {
				ret += rightMax - height[right]
			}
			right -= 1
		}
	}
	return ret
}

func main26() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
