package main

import "fmt"

func rightBound(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		m := l + (r-l)/2
		if arr[m] > target {
			r = m
		} else if arr[m] < target {
			l = m + 1
		} else {
			l = m + 1
		}
	}
	return l - 1
}

func main299() {
	fmt.Println(rightBound([]int{2, 3, 6, 6, 6, 7, 8, 10}, 6))
	fmt.Println(rightBound([]int{2, 3, 6, 6, 6, 7, 8, 10}, 3))
	fmt.Println(rightBound([]int{2, 3, 6, 6, 6, 7, 8, 10}, 1))
	fmt.Println(rightBound([]int{2, 3, 6, 6, 6, 7, 8, 10}, 11))
}
