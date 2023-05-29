package main

import "fmt"

// https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
// 剑指 Offer 11. 旋转数组的最小数字

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := left + (right-left)/2
		if numbers[left] == numbers[mid] && numbers[mid] == numbers[right] {
			left += 1
			right -= 1
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return numbers[left]
}

func main109() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(minArray([]int{2, 3, 4, 5, 1}))
	fmt.Println(minArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(minArray([]int{3, 3, 3, 1, 3}))
}
