package main

import (
	"fmt"
	"sort"
)

func findMinimumTime(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][1] < tasks[j][1] })
	run := make([]bool, tasks[len(tasks)-1][1]+1, tasks[len(tasks)-1][1]+1)
	res := 0
	for _, t := range tasks {
		s, e, d := t[0], t[1], t[2]
		for _, b := range run[s : e+1] {
			if b {
				d -= 1
			}
		}
		for i := e; d > 0; i -= 1 {
			if !run[i] {
				run[i] = true
				d, res = d-1, res+1
			}
		}
	}
	return res
}

func main241() {
	fmt.Println(findMinimumTime(
		[][]int{
			{1, 3, 2},
			{2, 5, 3},
			{5, 6, 2},
		},
	))
}
