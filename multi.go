package main

import "fmt"

func multi(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	flag := 1
	if a < 0 {
		flag, a = -flag, -a
	}
	if b < 0 {
		flag, b = -flag, -b
	}
	var recursion func(m, n int) int
	recursion = func(m, n int) int {
		if n == 0 {
			return 0
		}
		sum := 0
		if n&1 == 1 {
			sum += m
		}
		sum += recursion(m<<1, n>>1)
		return sum
	}
	if a < b {
		a, b = b, a
	}
	ret := recursion(a, b)
	if flag == 1 {
		return ret
	}
	return -ret
}

func main3() {
	fmt.Println(multi(0, 1))
	fmt.Println(multi(0, -1))
	fmt.Println(multi(0, 0))
	fmt.Println(multi(5, 5))
	fmt.Println(multi(5, 3))
	fmt.Println(multi(5, -3))
	fmt.Println(multi(-5, -3))
	fmt.Println(multi(-5, -3))
	fmt.Println(multi(-5, -3))
	fmt.Println(multi(-5, -3))
}
