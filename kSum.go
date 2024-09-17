package main

import (
	"fmt"
	"sort"
)

func kSum(nums []int, k int) int64 {
	fmt.Println(len(nums))
	helper := make([]int, 0)
	l := len(nums)
	var dfs func(idx int, sum int)
	dfs = func(idx int, sum int) {
		fmt.Println(idx, sum)
		if idx >= l {
			helper = append(helper, sum)
			return
		} else {
			dfs(idx+1, sum)
			dfs(idx+1, sum+nums[idx])
		}
	}
	dfs(0, 0)

	sort.Slice(helper, func(i, j int) bool {
		return helper[i] > helper[j]
	})
	// fmt.Println(helper)
	return int64(helper[k-1])
}

func main210() {
	fmt.Println(kSum([]int{
		153123449, -974739108, -408679566, -996444415,
		-978921261, 805907128, -102259288, -397930077,
		51033052, -193994032, 158654659, -486195972,
		-294264190, -65262667, 375941242, -890038230, 315970860,
		403847239, -32469129, -350561293, 192113942, 794248972,
		-632675681, 434029943, 746632801, 500370163, 164413366,
		346449701, 473890512,
	}, 1906))
}
