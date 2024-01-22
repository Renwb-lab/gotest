package main

import (
	"fmt"
	"strconv"
)

func maximumSwap2(num int) int {
	arr := []byte(strconv.Itoa(num))
	l := len(arr)
	res := num
	for i := 0; i < l; i += 1 {
		for j := i + 1; j < l; j += 1 {
			arr[i], arr[j] = arr[j], arr[i]
			t, _ := strconv.Atoi(string(arr))
			if res < t {
				res = t
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return res
}

func maximumSwap(num int) int {
	arr := []byte(strconv.Itoa(num))
	l := len(arr)
	idx1, idx2, maxIdx := -1, -1, l-1
	// maxIdx: 表示从后向前遍历截止到i处最大的数值下标
	// idx1: 表示从后向前遍历小于maxIdx的下标
	// idx2: 表示当idx1小于maxIdx时maxIdx下标
	// 98368
	for i := l - 1; i >= 0; i -= 1 {
		if arr[i] > arr[maxIdx] {
			maxIdx = i
		} else if arr[i] < arr[maxIdx] {
			idx1, idx2 = i, maxIdx
		}
	}
	if idx1 == -1 {
		return num
	}
	arr[idx1], arr[idx2] = arr[idx2], arr[idx1]
	res, _ := strconv.Atoi(string(arr))
	return res
}

func main() {
	fmt.Println(maximumSwap(2736))
}
