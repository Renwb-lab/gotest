package main

import "fmt"

// https://leetcode.cn/contest/ubiquant2022/problems/xdxykd/

func numberOfPairs(nums []int) int {
	mod := 1000000007
	img := func(n int) int {
		r := 0
		for ; n > 0 && n%10 == 0; n = n / 10 {
		}
		for ; n > 0; n = n / 10 {
			r = r*10 + n%10
		}
		return r
	}
	m := make(map[int]int, 0)
	for _, v := range nums {
		diff := v - img(v)
		if n, ok := m[diff]; ok {
			m[diff] = n + 1
		} else {
			m[diff] = 1
		}
	}

	ret := 0
	for _, v := range m {
		ret = (ret + v*(v-1)/2%mod) % mod
	}
	return ret
}

func main27() {
	arr := []int{
		17, 28, 39, 71,
	}
	fmt.Println(numberOfPairs(arr))
}
