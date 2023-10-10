package main

import "fmt"

func maximumSumOfHeights(maxHeights []int) int64 {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	l := len(maxHeights)
	st := []int{-1} // 存放的是下标
	left := make([]int, l+1)
	left[0] = 0 // left[i] 表示[0, i)的左边之和
	for i := 0; i < l; i += 1 {
		pre := 0
		// st 单调递减
		for len(st) > 1 && maxHeights[i] < maxHeights[st[len(st)-1]] {
			j := st[len(st)-1] // 出去值的下标
			st = st[:len(st)-1]
			pre -= maxHeights[j] * (j - st[len(st)-1])
		}
		pre += maxHeights[i] * (i - st[len(st)-1])
		st = append(st, i)
		left[i+1] = left[i] + pre
	}
	st = []int{l}
	right := make([]int, l+1)
	right[l] = 0
	for i := l - 1; i >= 0; i -= 1 {
		post := 0
		for len(st) > 1 && maxHeights[i] < maxHeights[st[len(st)-1]] {
			j := st[len(st)-1] // 出去值的下标
			st = st[:len(st)-1]
			post -= maxHeights[j] * (st[len(st)-1] - j)
		}
		post += maxHeights[i] * (st[len(st)-1] - i)
		st = append(st, i)
		right[i] = right[i+1] + post
	}
	res := 0
	for i := 0; i <= l; i += 1 {
		res = max(res, left[i]+right[i])
	}
	return int64(res)
}

func main152() {
	fmt.Println(maximumSumOfHeights([]int{5, 3, 4, 1, 1}))
}
