package main

// https://leetcode.cn/problems/stack-of-plates-lcci/description/

type StackOfPlates struct {
	data [][]int
	cap  int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{data: make([][]int, 0), cap: cap}
}

func (this *StackOfPlates) Push(val int) {
	if this.cap == 0 {
		return
	}
	n := len(this.data)
	if n == 0 || len(this.data[n-1]) == this.cap {
		this.data = append(this.data, []int{val})
	} else {
		this.data[n-1] = append(this.data[n-1], val)
	}
}

func (this *StackOfPlates) Pop() int {
	return this.PopAt(len(this.data) - 1)
}

func (this *StackOfPlates) PopAt(index int) int {
	if this.cap == 0 {
		return -1
	}
	n := len(this.data)

	if index < 0 || index > n-1 {
		return -1
	} else {
		ln := len(this.data[index])
		res := this.data[index][ln-1]
		if len(this.data[index]) == 1 {
			if n-1 == index {
				this.data = this.data[:index]
			} else {
				this.data = append(this.data[:index], this.data[index+1:]...)
			}
		} else {
			this.data[index] = this.data[index][:ln-1]
		}
		return res
	}
}
