package main

import "fmt"

func reverseArr(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < arr[r] {
			if arr[mid] < target && target <= arr[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if arr[l] <= target && target < arr[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}

func main297() {
	arr := []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	fmt.Println(reverseArr(arr, 10))
}
