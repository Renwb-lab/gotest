package main

// 这段代码演示了如何使用堆接口构建一个整数堆。

import (
	"container/heap"
	"fmt"
)

type NumItem struct {
	Base   int
	Offset int
	Num    int
}

// IntHeap 是一个由整数组成的最小堆。
type ItemHeap []NumItem

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].Num < h[j].Num }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ItemHeap) Push(x interface{}) {
	// Push 和 Pop 使用 pointer receiver 作为参数，
	// 因为它们不仅会对切片的内容进行调整，还会修改切片的长度。
	*h = append(*h, x.(NumItem))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeArr(arrayList [][]int) []int {
	l := 0
	h := &ItemHeap{}
	for i := range arrayList {
		if len(arrayList[i]) > 0 {
			*h = append(*h, NumItem{i, 0, arrayList[i][0]})
		}
		l += len(arrayList[i])
	}
	heap.Init(h)
	res, idx := make([]int, l), 0
	for len(*h) > 0 {
		item := heap.Pop(h).(NumItem)
		base, offset, num := item.Base, item.Offset, item.Num
		res[idx], idx = num, idx+1
		if offset+1 < len(arrayList[base]) {
			heap.Push(h, NumItem{base, offset + 1, arrayList[base][offset+1]})
		}
	}
	return res
}

// 这个示例会将一些整数插入到堆里面， 接着检查堆中的最小值，
// 之后按顺序从堆里面移除各个整数。
func main117() {
	arr1, arr2, arr3 := []int{1, 3, 5}, []int{2, 6, 8}, []int{5, 9}
	res := mergeArr([][]int{arr1, arr2, arr3})
	for i := range res {
		fmt.Printf("%d\t", res[i])
	}
	fmt.Println()
}
