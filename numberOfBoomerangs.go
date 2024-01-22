package main

import (
	"fmt"
)

// https://leetcode.cn/problems/number-of-boomerangs/description/
func numberOfBoomerangs(points [][]int) int {
	distance := func(i, j int) int {
		x := points[i][0] - points[j][0]
		y := points[i][1] - points[j][1]
		return x*x + y*y
	}
	length := len(points)
	mp := make([][]int, length)
	for i := range mp {
		mp[i] = make([]int, length)
	}
	for i := 0; i < length; i += 1 {
		for j := i + 1; j < length; j += 1 {
			d := distance(i, j)
			mp[i][j], mp[j][i] = d, d
		}
	}
	ans := 0
	for i := 0; i < length; i += 1 {
		helper := make(map[int]int)
		for j := 0; j < length; j += 1 {
			helper[mp[i][j]] += 1
		}
		for _, c := range helper {
			if c < 2 {
				continue
			}
			ans += (c - 1) * c
		}
	}

	return ans
}

func numberOfBoomerangs2(points [][]int) int {
	distance := func(i, j int) int {
		x := points[i][0] - points[j][0]
		y := points[i][1] - points[j][1]
		return x*x + y*y
	}
	ans := 0
	length := len(points)
	for i := 0; i < length; i += 1 {
		helper := make(map[int]int)
		for j := 0; j < length; j += 1 {
			helper[distance(i, j)] += 1
		}
		for _, c := range helper {
			ans += (c - 1) * c
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfBoomerangs([][]int{
		{0, 0},
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}))
}
