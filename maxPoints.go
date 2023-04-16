package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/x2n2g1/
// 直线上最多的点数

func gcd(x, y int) (int, int) {
	a, b := x, y
	if a < b {
		a, b = b, a
	}
	r := a % b
	for r > 0 {
		a, b = b, r
		r = a % b
	}
	//  a:9 b:6 -> r:3
	//  a:6 b:3 -> r:0
	return x / b, y / b
}

func computeKey(p1, p2 []int) string {
	toString := func(x, y int) string {
		// 这里存在问题，即相同斜率的直线可能是平行线， 所以使用这种方法不正确
		return fmt.Sprintf("%d____%d", x, y)
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	x, y := p1[0]-p2[0], p1[1]-p2[1]
	if x == 0 {
		return toString(0, 1)
	}
	if y == 0 {
		return toString(1, 0)
	}
	flag := true
	if (x < 0) != (y < 0) {
		flag = false
	}
	x, y = abs(x), abs(y)
	x, y = gcd(x, y)
	if flag {
		return toString(x, y)
	}
	return toString(-x, y)
}

func maxPoints2(points [][]int) int {
	ret := 1
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || (points[i][0] == points[j][0] && points[i][1] < points[j][1])
	})
	// 这里存在问题，即相同斜率的直线可能是平行线， 所以使用这种方法不正确
	helper := make(map[string]map[int]int)
	for i := 0; i < len(points); i += 1 {
		for j := i + 1; j < len(points); j += 1 {
			p1, p2 := points[i], points[j]
			key := computeKey(p1, p2)
			if _, ok := helper[key]; !ok {
				helper[key] = make(map[int]int)
			}
			helper[key][i] = 1
			helper[key][j] = 1
			if len(helper[key]) > ret {
				ret = len(helper[key])
			}
		}
	}
	return ret
}

func OnOneLine(p1, p2, p3 []int) bool {
	x1, x2 := p2[0]-p1[0], p3[0]-p2[0]
	y1, y2 := p2[1]-p1[1], p3[1]-p2[1]
	return x1*y2 == x2*y1
}

func maxPoints(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || (points[i][0] == points[j][0] && points[i][1] < points[j][1])
	})
	ret := 1
	for i := 0; i < len(points); i += 1 {
		for j := i + 1; j < len(points); j += 1 {
			p1, p2, temp := points[i], points[j], 2
			for k := j + 1; k < len(points); k += 1 {
				p3 := points[k]
				if OnOneLine(p1, p2, p3) {
					temp += 1
				}
			}
			if temp > ret {
				ret = temp
			}
		}
	}
	return ret
}

func main73() {
	points := [][]int{
		{3, 3},
		{1, 4},
		{1, 1},
		{2, 1},
		{2, 2},
	}
	fmt.Println(maxPoints(points))
}
