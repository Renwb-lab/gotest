package main

import (
	"encoding/json"
	"fmt"
)

var helper map[string]int

func getNewPizza(pizza []string, y int) []string {
	ret := make([]string, 0)
	for _, line := range pizza {
		ret = append(ret, line[y:])
	}
	return ret
}

func ways(pizza []string, k int) int {
	if helper == nil {
		helper = make(map[string]int)
	}
	pizzaByte, _ := json.Marshal(pizza)
	pizzaStr := string(pizzaByte) + fmt.Sprintf("____%d", k)
	if ret, ok := helper[pizzaStr]; ok {
		return ret
	}
	lines, rows := len(pizza), len(pizza[0])
	// 找到最左上的位置
	x, y := lines, rows
	for i, line := range pizza {
		for j, c := range line {
			if c == rune('A') {
				if i < x {
					x = i
				}
				if j < y {
					y = j
				}
			}
		}
	}
	if k <= 1 {
		if x == lines && y == rows {
			return 0
		}
		return 1
	}

	ret := 0
	// 先按照行进行切
	for x += 1; x < lines; x += 1 {
		ret += ways(pizza[x:], k-1)
		ret %= 1e9 + 7
	}
	// 后按照列进行切
	for y += 1; y < rows; y += 1 {
		newPizza := getNewPizza(pizza, y)
		ret += ways(newPizza, k-1)
		ret %= 1e9 + 7
	}
	helper[pizzaStr] = ret
	return ret
}

func main5() {
	pizza := []string{
		".A..A",
		"A.A..",
		"A.AA.",
		"AAAA.",
		"A.AA.",
	}
	ret := ways(pizza, 5)
	fmt.Printf("ret: %d\n", ret)
}
