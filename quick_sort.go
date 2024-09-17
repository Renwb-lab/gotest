package main

import "fmt"

func quickSort(arr []int) {
	var dp func(l, r int)
	dp = func(l, r int) {
		if l >= r {
			return
		}
		// mid := findMiddileV1(arr, l, r)
		mid := findMiddileV2(arr, l, r)
		dp(l, mid-1)
		dp(mid+1, r)
	}
	dp(0, len(arr)-1)
}

func findMiddileV1(arr []int, l, r int) int {
	// arr[l, idx) < t
	t, idx := arr[r], l-1       // r 是最后一个
	for i := l; i < r; i += 1 { // 不能遍历到最后一个
		if arr[i] < t {
			idx += 1
			arr[idx], arr[i] = arr[i], arr[idx] // 是交换
		}
	}
	idx += 1
	arr[idx], arr[r] = arr[r], arr[idx] // 和最后一个交换
	return idx
}

func findMiddileV2(arr []int, l, r int) int {
	t := arr[r]
	for l < r {
		for l < r && arr[l] <= t {
			l += 1
		}
		arr[r] = arr[l]
		for l < r && arr[r] > t {
			r -= 1
		}
		arr[l] = arr[r]
	}
	arr[l] = t
	return l
}

func main293() {
	arr := []int{5, 6, 2, 2, 3}
	quickSort(arr)
	fmt.Println(arr)
}
