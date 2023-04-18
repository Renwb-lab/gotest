package main

import "container/heap"

// https://leetcode.cn/leetbook/read/top-interview-questions/xalff2/
// 数据流的中位数

type Bigger []int

func (h Bigger) Len() int            { return len(h) }
func (h Bigger) Less(i, j int) bool  { return h[i] < h[j] }
func (h Bigger) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Bigger) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *Bigger) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type Small []int

func (h Small) Len() int            { return len(h) }
func (h Small) Less(i, j int) bool  { return h[i] > h[j] }
func (h Small) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Small) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *Small) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type MedianFinder struct {
	biger Bigger // 堆顶显示最小的元素
	small Small  // 堆顶显示最大的元素
}

func Constructor9() MedianFinder {
	return MedianFinder{
		biger: make([]int, 0),
		small: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	adjust := func() {
		if this.biger[0] < this.small[0] {
			n1 := heap.Pop(&this.biger)
			n2 := heap.Pop(&this.small)
			heap.Push(&this.biger, n2)
			heap.Push(&this.small, n1)
		}
	}
	if len(this.biger) == 0 && len(this.small) == 0 {
		this.biger = append(this.biger, num)
		return
	}
	if len(this.biger) == len(this.small) {
		heap.Push(&this.biger, num)
		adjust()
		return
	}
	if len(this.small) < len(this.biger) {
		heap.Push(&this.small, num)
		adjust()
		return
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.biger) > len(this.small) {
		return float64(this.biger[0])
	}
	return float64(this.small[0]+this.biger[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
