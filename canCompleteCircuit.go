package main

import "fmt"

// 加油站
// https://leetcode.cn/leetbook/read/top-interview-questions/xmguej/
func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); {
		gasSum, costSum, cnt := 0, 0, 0
		for j := i; cnt < len(gas); {
			gasSum += gas[j]
			costSum += cost[j]
			if gasSum < costSum {
				break
			}
			j = (j + 1) % len(gas)
			cnt += 1
		}
		if cnt == len(gas) {
			return i
		}
		i = i + cnt + 1
	}
	return -1
}

func main97() {
	// gas := []int{1, 2, 3, 4, 5}
	// cost := []int{3, 4, 5, 1, 2}
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
}
