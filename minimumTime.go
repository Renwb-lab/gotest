package main

import (
	"fmt"
	"sort"
)

func minimumTime(nums1 []int, nums2 []int, x int) int {
	s1, s2, n := 0, 0, len(nums1)
	id := make([][2]int, n)
	for i := range id {
		id[i] = [2]int{nums1[i], nums2[i]}
		s1 += nums1[i]
		s2 += nums2[i]
	}

	sort.Slice(id, func(i, j int) bool { return id[i][1] < id[j][1] })
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	// 定义 f[i+1][j]f[i+1][j]f[i+1][j] 表示从 0,1,2,⋯ ,i 中选 j 个下标（j≤i+1j），减少量最大是多少。
	// f[i+1][j]=max(f[i][j],f[i][j−1]+nums[i]+nums2[i]⋅j)
	f := make([]int, n+1)
	for i, p := range id {
		a, b := p[0], p[1]
		for j := i + 1; j > 0; j-- {
			f[j] = max(f[j], f[j-1]+a+b*j)
		}
	}
	for t, v := range f {
		if s1+s2*t-v <= x {
			return t
		}
	}
	return -1
}

func main269() {
	fmt.Println(minimumTime([]int{1, 2, 3}, []int{1, 2, 3}, 4))
}
