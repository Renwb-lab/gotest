package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/xm6kpg/
// Fizz Buzz

func fizzBuzz(n int) []string {
	ret := []string{}
	for i := 1; i <= n; i += 1 {
		t := strconv.Itoa(i)
		if i%15 == 0 {
			t = "FizzBuzz"
		} else if i%5 == 0 {
			t = "Buzz"
		} else if i%3 == 0 {
			t = "Fizz"
		}
		ret = append(ret, t)
	}
	return ret
}

func main92() {
	ret := fizzBuzz(15)
	for _, v := range ret {
		fmt.Printf("%s\n", v)
	}
	fmt.Println()
}
