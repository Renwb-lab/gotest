package main

import "fmt"

func robot(command string, obstacles [][]int, x int, y int) bool {
	px, py := 0, 0
	helper := make(map[string]struct{})
	helper[fmt.Sprintf("%d_%d", 0, 0)] = struct{}{}
	for _, c := range command {
		if c == 'U' {
			py += 1
		} else {
			px += 1
		}
		helper[fmt.Sprintf("%d_%d", px, py)] = struct{}{}
	}
	loop := x / px
	if loop > y/py {
		loop = y / py
	}
	key := fmt.Sprintf("%d_%d", x-loop*px, y-loop*py)
	if _, ok := helper[key]; !ok {
		return false
	}
	for _, line := range obstacles {
		if line[0] > x || line[1] > y {
			continue
		}
		loop := line[0] / px
		if loop > line[1]/py {
			loop = line[1] / py
		}
		if _, ok := helper[fmt.Sprintf("%d_%d", line[0]-loop*px, line[1]-loop*py)]; ok {
			return false
		}
	}

	return true
}

func main302() {
	fmt.Println(robot("RRU", [][]int{{5, 5},
		{9, 4},
		{9, 7},
		{6, 4},
		{7, 0},
		{9, 5},
		{10, 7},
		{1, 1},
		{7, 5}}, 1486, 743))
}
