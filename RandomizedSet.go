package main

import "math/rand"

// https://leetcode.cn/leetbook/read/top-interview-questions/xagm3s/
// 常数时间插入、删除和获取随机元素

type RandomizedSet struct {
	num2Idx map[int]int // key: num, val: id
	nums    []int       // num
}

func Constructor6() RandomizedSet {
	return RandomizedSet{num2Idx: make(map[int]int), nums: make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.num2Idx[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.num2Idx[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.num2Idx[val]; ok {
		if idx != len(this.nums)-1 {
			// 将最后一个数字调整当前位置上
			otherVal := this.nums[len(this.nums)-1]
			// 移动数组的位置
			this.nums[idx] = otherVal
			// 移动map的信息
			this.num2Idx[otherVal] = idx
		}
		delete(this.num2Idx, val)
		this.nums = this.nums[:len(this.nums)-1]
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
