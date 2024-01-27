package main

import "fmt"

func sumIndicesWithKSetBits(nums []int, k int) int {
	f := func(n, x int) int {
		r, t := 0, n
		for t > 0 {
			t = t & (t - 1)
			r += 1
		}
		if r == k {
			return x
		}
		return 0
	}
	res := 0
	for i := range nums {
		res += f(i, nums[i])
	}
	return res
}

func main() {
	fmt.Println(sumIndicesWithKSetBits([]int{5, 10, 1, 5, 2}, 1))
}
