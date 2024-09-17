package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/removing-minimum-number-of-magic-beans/
func minimumRemoval(beans []int) int64 {
	sort.Slice(beans, func(i, j int) bool { return beans[i] < beans[j] })
	sum, mx := 0, 0
	len := len(beans)
	for i, v := range beans {
		sum += v
		if (len-i)*v > mx {
			mx = (len - i) * v
		}
	}
	return int64(sum - mx)

}

func main272() {
	fmt.Println(minimumRemoval([]int{2, 10, 3, 2}))
}
