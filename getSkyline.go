package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/xa4clm/
// 天际线问题

func getSkyline2(buildings [][]int) [][]int {
	helper := [][]int{}
	for i := 0; i < len(buildings); i += 1 {
		b, e, h := buildings[i][0], buildings[i][1], buildings[i][2]
		for j := b; j <= e; j += 1 {
			helper = append(helper, []int{j, h})
		}
	}
	sort.Slice(helper, func(i, j int) bool {
		return helper[i][0] < helper[j][0] || (helper[i][0] == helper[j][0] && helper[i][1] < helper[j][1])
	})
	begin, end := helper[0][0], helper[len(helper)-1][0]+1
	newHelper := make([]int, end-begin+1)
	for i := 0; i < len(helper); i += 1 {
		newHelper[helper[i][0]-begin] = helper[i][1]
	}
	ret := [][]int{{begin, newHelper[0]}}
	preH := newHelper[0]
	for i := 1; i < len(newHelper); i += 1 {
		if newHelper[i] > preH {
			preH = newHelper[i]
			ret = append(ret, []int{i + begin, newHelper[i]})
		} else if newHelper[i] < preH {
			preH = newHelper[i]
			ret = append(ret, []int{i + begin - 1, newHelper[i]})
		}
	}
	return ret
}

type pair struct{ right, height int }

type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// 基本思路：天际线出现的位置只会在buildings的左右两侧
// 所以判断buildings左右两侧是否需要添加的结果中即可。
// 特别注意：物体的右侧边界是不包含在内的。
func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	// 1. 找到左右边界，并排序
	boundaries := make([]int, 0, n*2)
	for _, b := range buildings {
		boundaries = append(boundaries, b[0], b[1])
	}
	sort.Ints(boundaries)

	// 2. 针对每个边界进行判断，这里使用优先队列，支持自动排序
	idx, h := 0, hp{}
	ret := [][]int{}
	for _, b := range boundaries {
		// 先将左边界小于当前边界的建筑加入
		for idx < n && buildings[idx][0] <= b {
			heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
			idx += 1
		}
		// 然后将右边界大于等于当前边界的建筑移除, 注意这里的【等于】
		for len(h) > 0 && h[0].right <= b {
			heap.Pop(&h)
		}
		// 获取到满足条件的最大的高度
		maxn := 0
		if len(h) > 0 {
			maxn = h[0].height
		}
		// 这里需要判断是否出现过
		if len(ret) == 0 || ret[len(ret)-1][1] != maxn {
			ret = append(ret, []int{b, maxn})
		}
	}
	return ret
}

func main75() {
	buildings := [][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}
	ret := getSkyline(buildings)
	for _, line := range ret {
		for _, v := range line {
			fmt.Printf("%d\t", v)
		}
		fmt.Println()
	}
}
