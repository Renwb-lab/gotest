package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xatgye/
// 滑动窗口最大值

// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7

func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	var (
		queue []int // 降序队列，注意保存的是数组下标
		ret   = make([]int, l-k+1)
	)
	for i := 0; i < l; i++ {
		for len(queue) > 0 && queue[0] <= i-k {
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i >= k-1 {
			ret[i-k+1] = nums[queue[0]]
		}
	}
	return ret
}

func main87() {
	nums := []int{-7, -8, 7, 5, 7, 1, 6, 0}
	ret := maxSlidingWindow(nums, 4)
	for _, v := range ret {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}
