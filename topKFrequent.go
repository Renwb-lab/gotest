package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/xau4ci/
// 前 K 个高频元素

func topKFrequent(nums []int, k int) []int {
	type Pair struct {
		Val   int
		Times int
	}

	v2t := make(map[int]int, 0)
	for _, v := range nums {
		v2t[v] += 1
	}
	pairs := make([]Pair, 0, len(v2t))
	for k, v := range v2t {
		pairs = append(pairs, Pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Times > pairs[j].Times
	})
	ret := make([]int, 0, k)
	for i := 0; i < k; i += 1 {
		ret = append(ret, pairs[i].Val)
	}
	return ret
}

func main85() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	ret := topKFrequent(nums, k)
	for _, n := range ret {
		fmt.Printf("%d\t", n)
	}
	fmt.Println()
}
