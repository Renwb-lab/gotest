package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xajl22/
// 计算右侧小于当前元素的个数

// 线段树：跟节点能够代表所有访问的数字
// 更新一个区间时，是从顶向下更新的，如果能够覆盖当前节点，则更新当前节点的值。
// 如果不能覆盖的话，则判断是否能够覆盖左右节点。以此类推。
//
// 查找一个区间的值时，也是从顶向下更新的，如果能够覆盖当前节点，则返回当前节点的值。
// 如果不能覆盖的话，则判断是否能够覆盖左右节点，进行求和。
type SegmentTree struct {
	left       *SegmentTree
	right      *SegmentTree
	leftValue  int
	rightValue int
	cnt        int
}

func NewSegmentTree(leftValue, rightValue int) *SegmentTree {
	return &SegmentTree{left: nil, right: nil, cnt: 0, leftValue: leftValue, rightValue: rightValue}
}

func (s *SegmentTree) Update(l, r int) {
	if l <= s.leftValue && r >= s.rightValue {
		s.cnt += 1
		return
	}
	mid := (s.leftValue & s.rightValue) + (s.leftValue^s.rightValue)>>1
	// mid := s.leftValue + (s.rightValue-s.leftValue)/2
	if s.left == nil {
		s.left = NewSegmentTree(s.leftValue, mid)
	}
	if s.right == nil {
		s.right = NewSegmentTree(mid+1, s.rightValue)
	}

	if l <= mid {
		s.left.Update(l, r)
	}
	if r > mid {
		s.right.Update(l, r)
	}
	s.cnt += 1
}

func (s *SegmentTree) Query(l, r int) int {
	if l <= s.leftValue && r >= s.rightValue {
		return s.cnt
	}
	mid := (s.leftValue & s.rightValue) + (s.leftValue^s.rightValue)>>1
	// mid := s.leftValue + (s.rightValue-s.leftValue)/2
	if s.left == nil {
		s.left = NewSegmentTree(s.leftValue, mid)
	}
	if s.right == nil {
		s.right = NewSegmentTree(mid+1, s.rightValue)
	}

	ret := 0
	if l <= mid {
		ret += s.left.Query(l, r)
	}
	if r > mid {
		ret += s.right.Query(l, r)
	}
	return ret
}

func countSmaller(nums []int) []int {
	size := len(nums)
	ret := make([]int, size)
	minv, maxv := -10005, 10005
	// minv, maxv := 0, 7
	s := NewSegmentTree(minv, maxv)
	for i := size - 1; i >= 0; i -= 1 {
		ret[i] = s.Query(minv, nums[i]-1) // 求解时的区间非常重要
		// ret[i] = s.Query(nums[i], maxv)
		s.Update(nums[i], nums[i])
	}
	return ret
}

func main66() {
	nums := []int{5, 2, 2, 6, 1}
	ret := countSmaller(nums)
	for _, v := range ret {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}
