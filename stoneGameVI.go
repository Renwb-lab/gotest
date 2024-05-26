package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/stone-game-vi/?envType=daily-question&envId=2024-05-26
func stoneGameVI(aliceValues []int, bobValues []int) int {
	// (ai, bi) (aj, bj)
	// case1: ai, bj -> diff: ai - bj
	// case2: bi, aj -> diff: aj - bi
	// if ai - bj > aj - bi ==> case1
	//    ai + bi > aj + bj
	// if ai - bj < aj - bi ==> case2
	//    ai + bi < aj + bj
	// 因为ai + bi只和下标相关，所以可以先进行相加，然后排序
	// 先手总是可以想办法拿到最大的值
	// 这里不能进行模拟，因为如果ai + bi == aj + bj的话，
	// 就需要分情况讨论，而排序则不需要
	n := len(aliceValues)
	helper := make([][]int, n, n)
	for i := 0; i < n; i += 1 {
		a, b := aliceValues[i], bobValues[i]
		helper[i] = []int{a + b, a, b}
	}
	sort.Slice(helper, func(i, j int) bool {
		return helper[i][0] > helper[j][0]
	})
	s1, s2 := 0, 0
	for i := 0; i < n; i += 1 {
		if i%2 == 0 {
			s1 += helper[i][1]
		} else {
			s2 += helper[i][2]
		}
	}
	if s1 > s2 {
		return 1
	} else if s1 == s2 {
		return 0
	}
	return -1
}

func main() {
	fmt.Println(stoneGameVI(
		[]int{2, 9, 1, 1, 1, 3, 5, 8, 8, 6, 8, 6, 2, 4},
		[]int{1, 9, 7, 8, 3, 4, 2, 7, 8, 10, 1, 7, 10, 4}))
}
