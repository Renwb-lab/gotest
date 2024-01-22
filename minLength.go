package main

import (
	"fmt"
)

func minLength(s string) int {
	helper := func(arr []byte) (res []byte, flag bool) {
		i, l := 0, len(arr)
		for i < l {
			if i+1 < l && ((arr[i] == 'A' && arr[i+1] == 'B') ||
				(arr[i] == 'C' && arr[i+1] == 'D')) {
				flag, i = true, i+2
			} else {
				res = append(res, arr[i])
				i += 1
			}
		}
		return
	}

	arr, flag := []byte(s), false
	for {
		arr, flag = helper(arr)
		if !flag {
			break
		}
	}
	return len(arr)
}

func main() {
	fmt.Println(minLength("ABFCACDB"))
}
