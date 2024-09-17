package main

import (
	"container/heap"
	"sort"
)

type HP struct {
	sort.IntSlice
	size int
}

func (h *HP) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *HP) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func (h *HP) push(v int) {
	h.size += 1
	heap.Push(h, v)
}

func (h *HP) pop() int {
	h.size -= 1
	return heap.Pop(h).(int)
}

func (h *HP) prune() {
	for h.Len() > 0 {
		num := h.IntSlice[0]
		if h == small {
			num = -num
		}
		if d, has := delayed[num]; has {
			if d > 1 {
				delayed[num] -= 1
			} else {
				delete(delayed, num)
			}
			heap.Pop(h)
		} else {
			break
		}
	}
}

var delayed map[int]int
var small, large *HP

func medianSlidingWindow(nums []int, k int) []float64 {
	delayed = map[int]int{}
	small = &HP{}
	large = &HP{}
	makeBalance := func() {
		if small.size > large.size+1 {
			large.push(-small.pop())
			small.prune()
		} else if small.size < large.size {
			small.push(-large.pop())
			large.prune()
		}
	}
	insert := func(num int) {
		if small.Len() == 0 || num <= -small.IntSlice[0] {
			small.push(-num)
		} else {
			large.push(num)
		}
		makeBalance()
	}
	erase := func(num int) {
		delayed[num] += 1
		if num <= -small.IntSlice[0] {
			small.size -= 1
			if num == -small.IntSlice[0] {
				small.prune()
			}
		} else {
			large.size -= 1
			if num == large.IntSlice[0] {
				large.prune()
			}
		}
		makeBalance()
	}
	getMedian := func() float64 {
		if k&1 > 0 {
			return float64(-small.IntSlice[0])
		} else {
			return float64(-small.IntSlice[0]+large.IntSlice[0]) / 2
		}
	}
	for _, num := range nums[:k] {
		insert(num)
	}
	n := len(nums)
	ans := make([]float64, 1, n-k+1)
	ans[0] = getMedian()
	for i := k; i < n; i += 1 {
		insert(nums[i])
		erase(nums[i-k])
		ans = append(ans, getMedian())
	}
	return ans
}
