package main

import "fmt"

// https://leetcode.cn/problems/ugly-number-ii/description/
func nthUglyNumber(n int) int {
	min := func(args ...int) int {
		r := args[0]
		for _, n := range args[1:] {
			if n < r {
				r = n
			}
		}
		return r
	}
	if n == 1 {
		return 1
	}
	helper := make([]int, n+1)
	helper[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i += 1 {
		n := min(helper[p2]*2, helper[p3]*3, helper[p5]*5)
		helper[i] = n
		if n == helper[p2]*2 {
			p2 += 1
		}
		if n == helper[p3]*3 {
			p3 += 1
		}
		if n == helper[p5]*5 {
			p5 += 1
		}
	}
	return helper[n]
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
