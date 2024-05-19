package main

import (
	"fmt"
	"sort"
)

func bestTeamScore(scores []int, ages []int) int {
	scoreAges := [][]int{}
	for i := range scores {
		scoreAges = append(scoreAges, []int{scores[i], ages[i]})
	}
	sort.Slice(scoreAges, func(i, j int) bool {
		if scoreAges[i][1] < scoreAges[j][1] {
			return true
		} else if scoreAges[i][1] == scoreAges[j][1] {
			return scoreAges[i][0] < scoreAges[j][0]
		} else {
			return false
		}
	})

	max := func(args ...int) int {
		x := args[0]
		for _, y := range args[1:] {
			if y > x {
				x = y
			}
		}
		return x
	}

	helper := make([]int, len(scoreAges))
	helper[0] = scoreAges[0][0]
	for i := 1; i < len(helper); i += 1 {
		helper[i] = scoreAges[i][0]
		t := 0
		for j := i - 1; j >= 0; j -= 1 {
			if scoreAges[j][1] < scoreAges[i][1] &&
				scoreAges[j][0] > scoreAges[i][0] {
				continue
			}
			t = max(t, helper[j])
		}
		helper[i] += t
	}
	return max(helper...)
}

func main() {
	fmt.Println(bestTeamScore(
		[]int{1, 3, 7, 3, 2, 4, 10, 7, 5},
		[]int{4, 5, 2, 1, 1, 2, 4, 1, 4}))
}
