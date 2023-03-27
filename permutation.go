package main

import (
	"fmt"
)

// https://leetcode.cn/problems/permutation-ii-lcci/?favorite=xb9lfcwi

func permutation(s string) []string {
	ret := make([]string, 0)
	sbyte := []byte(s)
	end := len(sbyte)

	var cursion func(idx int)
	cursion = func(idx int) {
		if idx == end {
			ret = append(ret, string(sbyte))
		}
		visited := make(map[byte]int)
		for i := idx; i < end; i += 1 {
			if _, ok := visited[sbyte[i]]; ok {
				continue
			}
			visited[sbyte[i]] = 1
			sbyte[idx], sbyte[i] = sbyte[i], sbyte[idx]
			cursion(idx + 1)
			sbyte[idx], sbyte[i] = sbyte[i], sbyte[idx]
		}
	}

	cursion(0)
	return ret
}

func main9() {
	s := "jawa"
	ret := permutation(s)
	for _, line := range ret {
		fmt.Println(line)
	}
}
