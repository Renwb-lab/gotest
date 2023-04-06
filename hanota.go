package main

import "fmt"

// https://leetcode.cn/problems/hanota-lcci/?favorite=xb9lfcwi

func hanota(A []int, B []int, C []int) []int {
	l := len(A)
	var move func(n int, a *[]int, b *[]int, c *[]int)
	move = func(n int, a *[]int, b *[]int, c *[]int) {
		if n == 1 {
			*c = append(*c, A[n-1])
			A = A[:n-1]
			return
		} else {
			move(n-1, a, c, b)         // a先将 n-1通过c移动到b上
			*c = append(*c, (*a)[n-1]) // 移动最大的数字
			*a = (*a)[:n-1]            //原来的数据记得出去
			move(n-1, b, a, c)         // 最后b将n-1借助a移动到c上
		}
	}
	move(l, &A, &B, &C)
	return C
}

func main20() {
	A, B, C := []int{3, 2, 1}, []int{}, []int{}
	r := hanota(A, B, C)
	for _, v := range r {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}
