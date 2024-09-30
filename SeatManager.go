package main

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/seat-reservation-manager/?envType=daily-question&envId=2024-09-30
type SeatManager struct {
	// 只保存空余的座位
	sort.IntSlice
}

func Constructorer(n int) SeatManager {
	arr := make(sort.IntSlice, n)
	for i := 0; i < n; i += 1 {
		arr[i] = i + 1
	}
	return SeatManager{
		arr,
	}
}

func (this *SeatManager) Push(x any) {
	this.IntSlice = append(this.IntSlice, x.(int))
}

func (this *SeatManager) Pop() any {
	x := this.IntSlice[len(this.IntSlice)-1]
	this.IntSlice = this.IntSlice[:len(this.IntSlice)-1]
	return x
}

func (this *SeatManager) Less(i, j int) bool {
	return this.IntSlice[i] < this.IntSlice[j]
}

func (this *SeatManager) Reserve() int {
	// 预约的话，则将空余的座位移除一个
	return heap.Pop(this).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this, seatNumber)
}
