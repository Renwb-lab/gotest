package main

import "fmt"

// 从调整排序后的数组中查找指定的数据
func findNumFromArr(arr []int, target int) int {
	b, e := 0, len(arr)-1
	for b <= e {
		m := b + (e-b)/2
		if arr[m] == target {
			return m
		}
		if arr[b] <= arr[m] {
			if target >= arr[b] && target < arr[m] {
				e = m - 1
			} else {
				b = m + 1
			}
		} else {
			if target > arr[m] && target <= arr[e] {
				b = m + 1
			} else {
				e = m - 1
			}
		}
	}
	return -1
}

func main105() {
	arr := []int{4, 5, 6, 1, 2, 3}
	fmt.Println(findNumFromArr(arr, 4))
}
