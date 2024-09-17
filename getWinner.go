package main

import (
	"fmt"
)

func getWinner(arr []int, k int) int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	pre := max(arr[0], arr[1])
	if k == 1 {
		return pre
	}
	maxValue, cnt := pre, 1
	for _, x := range arr[2:] {
		if pre > x {
			cnt += 1
			if cnt == k {
				return pre
			}
		} else {
			pre, cnt = x, 1
		}
		maxValue = max(maxValue, pre)
	}
	return maxValue
}

func main243() {
	fmt.Println(getWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2))
}
