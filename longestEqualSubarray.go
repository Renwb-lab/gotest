package main

import "fmt"

// https://leetcode.cn/problems/find-the-longest-equal-subarray/?envType=daily-question&envId=2024-05-23
func longestEqualSubarray(nums []int, k int) int {
	m := make(map[int][]int)
	for i, n := range nums {
		m[n] = append(m[n], i)
	}

	ans := 0
	for _, v := range m {
		j := 0
		for i := 0; i < len(v); i += 1 {
			for v[i]-v[j]-(i-j) > k {
				j += 1
			}
			if i-j+1 > ans {
				ans = i - j + 1
			}
		}
	}
	return ans
}

func main248() {
	fmt.Println(longestEqualSubarray([]int{1, 3, 2, 3, 1, 3}, 2))
}
