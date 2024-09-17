package main

import (
	"fmt"
)

// https://leetcode.cn/problems/water-and-jug-problem/
func canMeasureWater1(x int, y int, z int) bool {

	// 把 X 壶的水灌进 Y 壶，直至灌满或倒空；
	// 把 Y 壶的水灌进 X 壶，直至灌满或倒空；
	// 把 X 壶灌满；
	// 把 Y 壶灌满；
	// 把 X 壶倒空；
	// 把 Y 壶倒空。

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	st := make([][]int, 0)
	mp := make(map[string]struct{})
	st = append(st, []int{0, 0}) // 初始化为0， 0
	for len(st) > 0 {
		remainX, remainY := st[len(st)-1][0], st[len(st)-1][1]
		st = st[:len(st)-1]
		fmt.Println(remainX, remainY)
		if remainX == z || remainY == z || remainX+remainY == z {
			return true
		}
		key := fmt.Sprintf("%d-%d", remainX, remainY)
		if _, ok := mp[key]; ok {
			continue
		}
		mp[key] = struct{}{}
		// 把 X 壶灌满
		st = append(st, []int{x, remainY})
		// 把 Y 壶灌满
		st = append(st, []int{remainX, y})
		// 把 X 壶倒空
		st = append(st, []int{0, remainY})
		// 把 Y 壶倒空
		st = append(st, []int{remainX, 0})
		// 把 X 壶的水灌进 Y 壶，直至灌满或倒空
		st = append(st, []int{remainX - min(y-remainY, remainX), remainY + min(y-remainY, remainX)})
		// 把 Y 壶的水灌进 X 壶，直至灌满或倒空
		st = append(st, []int{remainX + min(x-remainX, remainY), remainY - min(x-remainX, remainY)})
	}
	return false
}

func canMeasureWater(x int, y int, z int) bool {
	gcd := func(a, b int) int {
		for a%b != 0 {
			a, b = b, a%b
		}
		return b
	}
	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}
	return z%gcd(x, y) == 0

}
func main071013() {
	fmt.Println(canMeasureWater(3, 5, 4))
}
