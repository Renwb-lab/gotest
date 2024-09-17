package main

import (
	"fmt"
)

func main290() {
	n := 0
	_, _ = fmt.Scan(&n)
	one, two := []int{}, []int{}
	for i := 0; i < n; i += 1 {
		t := 0
		fmt.Scan(&t)
		if t%2 == 1 {
			one = append(one, t)
		} else {
			two = append(two, t)
		}
	}
	isPrime := func(x int) bool {
		for i := 2; i*i <= x+1; i += 1 {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	twomatch := make([]int, len(two), len(two))
	var match func(o int, v []bool) bool
	match = func(o int, v []bool) bool {
		for i, t := range two {
			if !v[i] && isPrime(o+t) {
				v[i] = true
				if twomatch[i] == 0 || match(twomatch[i], v) {
					twomatch[i] = o
					return true
				}
			}
		}
		return false
	}
	ans := 0
	for _, o := range one {
		v := make([]bool, len(two), len(two))
		if match(o, v) {
			ans += 1
		}
	}
	fmt.Println(ans)
}
