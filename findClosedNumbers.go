package main

import "fmt"

// https://leetcode.cn/problems/closed-number-lcci/?favorite=xb9lfcwi

func findClosedNumbers(num int) []int {
	x := 0
	for u := num; u > 0; u = u & (u - 1) {
		x += 1
	}
	isFeasible := func(other int) bool {
		x1 := 0
		for u := other; u > 0; u = u & (u - 1) {
			x1 += 1
		}
		return x1 == x
	}
	if x == 31 {
		return []int{-1, -1}
	}
	big := num + 1
	for big > 0 {
		if isFeasible(big) {
			break
		}
		big += 1
	}
	if big < 0 {
		big = -1
	}
	small := num - 1
	for small > 0 {
		if isFeasible(small) {
			break
		}
		small -= 1
	}
	if small == 0 {
		small = -1
	}
	return []int{big, small}
}

func main8() {
	ret := findClosedNumbers(2147483647)
	for _, v := range ret {
		fmt.Println(v)
	}
}
