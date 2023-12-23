package main

import (
	"fmt"
	"math"
)

func feasible(n int) bool {
	if n < 10 {
		return true
	}
	a, b := n, 0
	for a > 0 {
		t := a % 10
		b = b*10 + t
		a = a / 10
	}
	return b == n
}
func minimumCost(nums []int) int64 {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	x := float64(sum) / float64(len(nums))
	avg := int(math.Floor(x + 0.5))
	t := 0
	for i := 0; ; i += 1 {
		if feasible(avg + i) {
			t = avg + i
			break
		}
		if feasible(avg - i) {
			t = avg - i
			break
		}
	}
	r := int64(0)
	for _, v := range nums {
		if v > t {
			r += int64(v - t)
		} else {
			r += int64(t - v)
		}
	}
	return r
}

func main() {
	fmt.Println(minimumCost([]int{5, 2, 1}))
}
