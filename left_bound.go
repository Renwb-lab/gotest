package main

import (
	"fmt"
)

func leftBound(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		m := l + (r-l)/2
		if arr[m] > target {
			r = m
		} else if arr[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func main258() {
	arr := []int{2, 3, 6, 6, 6, 7, 8, 10}
	fmt.Println(leftBound(arr, 6))
	fmt.Println(leftBound(arr, 3))
	fmt.Println(leftBound(arr, 1))
	fmt.Println(leftBound(arr, 11))
}
