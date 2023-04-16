package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/x20yuc/
// 计数质数

var prime []int

func init() {
	prime = genPrimes(5000000)
}

func genPrimes(limit int) []int {
	res := make([]int, 0)
	notPrimes := make([]bool, limit)
	for i := 2; i < limit; i += 1 {
		if !notPrimes[i] {
			res = append(res, i)
			// 这里需要特别注意
			for j := i * i; j < limit; j += i {
				notPrimes[j] = true
			}
		}
	}
	return res
}
func countPrimes(n int) int {
	return sort.SearchInts(prime, n)
}

func main70() {
	fmt.Println(countPrimes(10))
}
