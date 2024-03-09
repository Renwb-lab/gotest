package main

import "fmt"

func innerMerge(arr []int, left, mid, right int) {
	helper := make([]int, 0, right-left+1)
	i, j := left, mid+1
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			helper = append(helper, arr[i])
			i += 1
		} else {
			helper = append(helper, arr[j])
			j += 1
		}
	}
	for ; i <= mid; i += 1 {
		helper = append(helper, arr[i])
	}
	for ; j <= right; j += 1 {
		helper = append(helper, arr[j])
	}
	for i, n := range helper {
		arr[left+i] = n
	}
}

func mergeSort(arr []int) {
	var dp func(left, right int)
	dp = func(left, right int) {
		if left >= right {
			return
		}
		mid := left + (right-left)/2
		dp(left, mid)
		dp(mid+1, right)
		innerMerge(arr, left, mid, right)
	}
	dp(0, len(arr)-1)
}

func mergeSortV2(arr []int) {
	length := len(arr)
	half := length / 2
	for i := 2; i <= (half+1)*2; i = i * 2 {
		for l := 0; l < length; {
			r := l + i - 1
			m := l + (r-l)/2
			if r >= length {
				r = length - 1
			}
			if m >= length {
				m = length - 1
			}
			innerMerge(arr, l, m, r)
			fmt.Println(arr, l, m, r)
			l = r + 1
		}
	}
}

func main() {
	arr := []int{5, 6, 2, 2, 3, 2, 3, 9, 0, 8}
	mergeSortV2(arr)
	fmt.Println(arr)
}
