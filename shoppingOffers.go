package main

import "fmt"

func shoppingOffers(price []int, special [][]int, needs []int) int {
	sum := 0
	for i := range needs {
		sum += needs[i] * price[i]
	}
	ans := sum
	helper := make([]int, len(special), len(special))
	for i := range special {
		t := 0
		for j := range price {
			t += special[i][j] * price[j]
		}
		helper[i] = t
	}
	meet := func(idx int) bool {
		for i := range needs {
			if special[idx][i] > needs[i] {
				return false
			}
		}
		return true
	}
	dec := func(idx int) {
		for i := range needs {
			needs[i] -= special[idx][i]
		}
	}
	add := func(idx int) {
		for i := range needs {
			needs[i] += special[idx][i]
		}
	}
	l := len(needs)
	var dp func(idx int)
	dp = func(idx int) {
		if idx >= len(special) {
			return
		}
		if meet(idx) {
			if helper[idx] > special[idx][l] {
				sum += (special[idx][l] - helper[idx])
				if ans > sum {
					ans = sum
				}
				dec(idx)
				dp(idx)
				add(idx)
				sum -= (special[idx][l] - helper[idx])
			}
		}
		dp(idx + 1)
	}
	dp(0)
	return ans
}

func main307() {
	fmt.Println(shoppingOffers(
		[]int{9},
		[][]int{
			{1, 10},
			{2, 2},
		},
		[]int{3},
	))
}
