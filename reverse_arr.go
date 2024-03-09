package main

import "fmt"

func reverseArr(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < arr[r] {
			if arr[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if arr[mid] > arr[l] {
			if target < arr[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		fmt.Println(l, r, mid, arr[mid])
	}
	return -1
}

func main() {
	arr := []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	fmt.Println(reverseArr(arr, 5))
}
