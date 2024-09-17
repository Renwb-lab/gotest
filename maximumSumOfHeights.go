package main

import "fmt"

func maximumSumOfHeights(maxHeights []int) int64 {
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	length := len(maxHeights)
	left := make([]int, length) // 存放左侧小于当前值最近的下标
	st := make([]int, 0)        // 借助单调栈实现
	for i, x := range maxHeights {
		// 递增
		for len(st) > 0 && maxHeights[st[len(st)-1]] > x {
			// > 不是>=， 是希望计算最小的小于当前值的小标
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		} else {
			left[i] = -1
		}
		st = append(st, i)
	}
	st = make([]int, 0)
	right := make([]int, length)
	for i := length - 1; i >= 0; i -= 1 {
		x := maxHeights[i]
		// 递增
		for len(st) > 0 && maxHeights[st[len(st)-1]] >= x {
			// >= 不是>， 是希望计算最大的小于当前值的小标
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		} else {
			right[i] = length
		}
		st = append(st, i)
	}
	f := make([]int64, length)
	for i, x := range maxHeights {
		if i > 0 && x >= maxHeights[i-1] {
			f[i] = f[i-1] + int64(x)
		} else {
			j := left[i]
			f[i] = int64(x) * int64(i-j)
			if j != -1 {
				f[i] += f[j]
			}
		}
	}
	g := make([]int64, length)
	for i := length - 1; i >= 0; i -= 1 {
		x := maxHeights[i]
		if i < length-1 && x >= maxHeights[i+1] {
			g[i] = g[i+1] + int64(x)
		} else {
			j := right[i]
			g[i] = int64(j-i) * int64(x)
			if j != length {
				g[i] += g[j]
			}
		}
	}
	ans := int64(0)
	for i, x := range maxHeights {
		ans = max(ans, f[i]+g[i]-int64(x))
	}
	return ans
}

func main260() {
	fmt.Println(maximumSumOfHeights([]int{5, 3, 4, 1, 1}))
}
