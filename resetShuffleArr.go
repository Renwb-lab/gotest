package main

import "math/rand"

// https://leetcode.cn/leetbook/read/top-interview-questions/xmchc3/
// 打乱数组

type Solution struct {
	old []int
	cur []int
}

func Constructor4(nums []int) Solution {
	return Solution{
		old: append([]int{}, nums...),
		cur: append([]int{}, nums...),
	}
}

func (this *Solution) Reset() []int {
	return this.old
}

func (this *Solution) Shuffle() []int {
	size := len(this.cur)
	for i := size - 1; i >= 0; i -= 1 {
		idx := rand.Intn(i + 1)
		this.cur[i], this.cur[idx] = this.cur[idx], this.cur[i]
	}
	return this.cur
}
