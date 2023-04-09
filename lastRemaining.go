package main

import "fmt"

// https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/?company_slug=ubiquanthr

// 剑指 Offer 62. 圆圈中最后剩下的数字

// 0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。求出这个圆圈里剩下的最后一个数字。
//
// 例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。

// 示例 1：

// 输入: n = 5, m = 3
// 输出: 3
// 示例 2：

// 输入: n = 10, m = 17
// 输出: 2

func lastRemaining(n int, m int) int {
	helper := make([]int, n)
	for i := 0; i < n; i += 1 {
		helper[i] = i
	}
	for len(helper) > 1 {
		t, idx := 0, 0
		for t < m {
			for i := 0; t < m && i < len(helper); i += 1 {
				t, idx = t+1, i
			}
		}
		if idx == 0 {
			helper = helper[1:]
		} else if idx == len(helper) {
			helper = helper[:len(helper)-1]
		} else {
			helper = append(helper[idx+1:], helper[:idx]...)
		}

	}
	return helper[0]
}

// https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/solutions/176636/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-by-lee/
func lastRemaining2(n int, m int) int {
	var dp func(n, m int) int
	dp = func(n, m int) int {
		if n == 1 {
			return 0
		}
		x := dp(n-1, m)
		return (m + x) % n
	}
	return dp(n, m)
}

func lastRemaining3(n int, m int) int {
	// f(n,m) -> f(n-1, m) = x
	// f(n, m) = (m + x) % n
	// 第一个删除的数据： m % n
	f := 0
	for i := 2; i <= n; i += 1 {
		f = (m + f) % i
	}
	return f
}

func main38() {
	n, m := 5, 3
	fmt.Println(lastRemaining(n, m))
	fmt.Println(lastRemaining2(n, m))
	fmt.Println(lastRemaining3(n, m))
}
