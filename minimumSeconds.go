package main

import "fmt"

// https://leetcode.cn/problems/minimum-seconds-to-equalize-a-circular-array/?envType=daily-question&envId=2024-05-26
func minimumSeconds(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	// 逆向思维，最终只会出现一个整数
	mp := make(map[int][]int)
	for i, n := range nums {
		mp[n] = append(mp[n], i)
	}
	n := len(nums)
	ans := n
	for _, pos := range mp {
		// 这是个逻辑上的循环数组，这步算的是头尾之间的距离
		mx := pos[0] + n - pos[len(pos)-1]
		for i := 1; i < len(pos); i += 1 {
			mx = max(mx, pos[i]-pos[i-1])
		}
		ans = min(ans, mx/2)
	}
	return ans
}

func main() {
	fmt.Print(minimumSeconds([]int{1, 2, 3, 4, 5}))
}
