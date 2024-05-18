package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	arr := []byte(num)
	for i := 0; i < k; i += 1 {
		j := 0
		for ; j < len(arr) && arr[j] == '0'; j += 1 {
		}
		if j == len(arr) {
			return "0"
		}
		var newArr []byte
		startLoc := j
		for ; j+1 < len(arr); j += 1 {
			if arr[j] > arr[j+1] {
				newArr = append(arr[startLoc:j], arr[j+1:]...)
				break
			}
		}
		if len(newArr) == 0 {
			newArr = arr[startLoc : len(arr)-1]
		}
		arr = append([]byte{}, newArr...)
	}
	j := 0
	for ; j < len(arr) && arr[j] == '0'; j += 1 {
	}
	if j == len(arr) {
		return "0"
	}
	return string(arr[j:])
}

func removeKdigitsV1(num string, k int) string {
	stack := []byte{}
	for i := range num {
		c := num[i]
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > c {
			stack = stack[:len(stack)-1]
			k -= 1
		}
		stack = append(stack, c)
	}
	stack = stack[:len(stack)-k]
	res := strings.TrimLeft(string(stack), "0")
	if len(res) == 0 {
		return "0"
	}
	return res
}

func main() {
	fmt.Println(removeKdigitsV1("1234567890", 9))
}
