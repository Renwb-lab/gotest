package main

import (
	"sort"
)

func findWinners(matches [][]int) [][]int {
	lossCount := map[int]int{}
	for _, m := range matches {
		// 如果没有设置过，就置0。
		if lossCount[m[0]] == 0 {
			lossCount[m[0]] = 0
		}
		lossCount[m[1]]++
	}

	ans := make([][]int, 2)
	for player, cnt := range lossCount {
		if cnt < 2 {
			ans[cnt] = append(ans[cnt], player)
		}
	}

	sort.Slice(ans[0], func(i, j int) bool { return ans[0][i] < ans[0][j] })
	sort.Slice(ans[1], func(i, j int) bool { return ans[1][i] < ans[1][j] })
	return ans
}

func main15() {

}
