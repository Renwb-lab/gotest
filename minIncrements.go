package main

import (
	"fmt"
)

// https://leetcode.cn/problems/make-costs-of-paths-equal-in-a-binary-tree/
// 2673. 使二叉树所有路径值相等的最小代价
func minIncrements(n int, cost []int) int {
	res := 0
	layer := func(begin, end int, input []int) (output []int) {
		for i, o := begin, 0; i < end; i, o = i+2, o+2 {
			diff := cost[i] + input[o] - (cost[i+1] + input[o+1])
			if diff > 0 {
				output = append(output, cost[i]+input[o])
				res += diff
				fmt.Println(cost[i]+input[o], cost[i+1], cost[i+1]+input[o+1], diff, res)

			} else {
				output = append(output, cost[i+1]+input[o+1])
				res -= diff
				fmt.Println(cost[i+1]+input[o+1], cost[i], cost[i]+input[o], -diff, res)
			}
		}
		return
	}
	half := n/2 + 1
	output := make([]int, half, half)
	for begin, end := n/2, n-1; end > 0; begin, end = begin/2, begin-1 {
		// begin, end: 3, 6		-> output: [3, 3], res: 3
		// begin, end: 1, 2		-> output: [3],    res: 6
		output = layer(begin, end, output)
		fmt.Println(begin, end, output, res)
		fmt.Println("===========")
	}

	return res
}

func main278() {
	fmt.Println(minIncrements(15, []int{
		764,
		1460, 2664,
		764, 2725, 4556, 5305,
		8829, 5064, 5929, 7660, 6321, 4830, 7055, 3761,
	}))
}
