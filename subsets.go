package main

import (
	"fmt"
)

// https://leetcode.cn/problems/power-set-lcci/?favorite=xb9lfcwi

func subsetsV1(nums []int) [][]int {
	res := make([][]int, 0)
	line := make([]int, 0)
	res = append(res, line)
	for i := 0; i < len(nums); i += 1 {
		for _, old := range res {
			line = append([]int{}, old...)
			line = append(line, nums[i])
			res = append(res, line)
		}
	}
	return res
}

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	ret = append(ret, []int{})
	l := len(nums)

	var newLine []int
	var recursion func(idx int)
	recursion = func(idx int) {
		if idx == l {
			return
		}
		for i := idx; i < l; i += 1 {
			newLine = []int{}
			for _, line := range ret {
				newLine := append(newLine, line...)
				newLine = append(newLine, nums[i])
				ret = append(ret, newLine)
			}
		}
	}
	recursion(0)
	return ret
}

func main316() {
	nums := []int{9, 0, 3, 5, 7}
	ret := subsetsV1(nums)
	for _, line := range ret {
		for _, n := range line {
			fmt.Printf("%d\t", n)
		}
		fmt.Println()
	}

}
