package main

import "fmt"

// 45. 跳跃游戏 II
// https://leetcode.cn/problems/jump-game-ii/?company_slug=baidu

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	begin, ret := 0, 0
	for begin <= len(nums)-1 {
		preVal := begin
		maxPos := 0
		ret += 1
		for i := 1; i <= nums[preVal]; i += 1 {
			pos := preVal + i
			if pos >= len(nums)-1 {
				return ret
			}
			if pos+nums[pos] > maxPos {
				maxPos = pos + nums[pos]
				begin = pos
			}
		}
		if preVal == begin {
			return -1
		}

	}
	return ret
}

type A struct {
	a int
	b string
	c byte
}

func main98() {
	fmt.Println(jump([]int{1, 2}))
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
	fmt.Println(jump([]int{3, 2, 1, 0, 4}))
	a := A{
		a: 1,
		b: "b",
		c: 'c',
	}
	fmt.Printf("%v\n", a)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%#v\n", a)
}
