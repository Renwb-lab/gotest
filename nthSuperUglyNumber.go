package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/super-ugly-number/
func nthSuperUglyNumber(n int, primes []int) int {
	min := func(args ...int) int {
		r := args[0]
		for _, n := range args[1:] {
			if n < r {
				r = n
			}
		}
		return r
	}
	helper := make([]int, n+1)
	helper[1] = 1
	m := len(primes)
	pointer := make([]int, m)
	for i := range pointer {
		pointer[i] = 1
	}
	for i := 2; i <= n; i += 1 {
		num := math.MaxInt64
		for j := range primes {
			num = min(num, helper[pointer[j]]*primes[j])
		}
		for j := range primes {
			if num == helper[pointer[j]]*primes[j] {
				pointer[j] += 1
			}
		}
		helper[i] = num
	}

	return helper[n]
}

func main() {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}
