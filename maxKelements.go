package main

import (
	"container/heap"
	"sort"
)

func maxKelements(nums []int, k int) (ans int64) {
	h := hp2{nums}
	heap.Init(&h)
	for ; k > 0; k -= 1 {
		ans += int64(h.IntSlice[0])
		h.IntSlice[0] = (h.IntSlice[0] + 2) / 3
		heap.Fix(&h, 0)
	}
	return
}

type hp2 struct{ sort.IntSlice }

func (h hp2) Less(c, p int) bool { return h.IntSlice[c] > h.IntSlice[p] }
func (hp2) Push(any)             {} //在上述代码中没有使用, 所以不要实现
func (hp2) Pop() (_ any)         { return }
