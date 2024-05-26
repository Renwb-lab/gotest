package main

import "slices"

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

	slices.Sort(ans[0])
	slices.Sort(ans[1])
	return ans
}

func main() {

}
