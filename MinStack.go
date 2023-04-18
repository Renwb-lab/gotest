package main

// https://leetcode.cn/leetbook/read/top-interview-questions/xa7r55/
// 最小栈

type MinStack struct {
	s1 []int
	s2 []int
}

func Constructor8() MinStack {
	return MinStack{
		s1: make([]int, 0),
		s2: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.s1 = append(this.s1, val)
	if len(this.s2) == 0 {
		this.s2 = append(this.s2, val)
	} else {
		minv := this.s2[len(this.s2)-1]
		if minv < val {
			this.s2 = append(this.s2, minv)
		} else {
			this.s2 = append(this.s2, val)
		}
	}
}

func (this *MinStack) Pop() {
	this.s1 = this.s1[:len(this.s1)-1]
	this.s2 = this.s2[:len(this.s2)-1]
}

func (this *MinStack) Top() int {
	return this.s1[len(this.s1)-1]
}

func (this *MinStack) GetMin() int {
	return this.s2[len(this.s1)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
