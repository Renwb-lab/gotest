package main

import "sort"

// https://leetcode.cn/problems/the-latest-time-to-catch-a-bus/?envType=daily-question&envId=2024-09-21
func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Slice(buses, func(i, j int) bool { return buses[i] < buses[j] })
	sort.Slice(passengers, func(i, j int) bool { return passengers[i] < passengers[j] })
	n, m := len(buses), len(passengers)
	idx := 0
	people := 0
	for i := 0; i < n; i += 1 {
		j := 0
		people = 0
		for ; j < capacity && idx+j < m && passengers[idx+j] <= buses[i]; j += 1 {
			people += 1
		}
		idx += j
	}
	idx -= 1
	res := buses[n-1]
	if people == capacity {
		res = passengers[idx]
	}
	for idx >= 0 && res == passengers[idx] {
		idx, res = idx-1, res-1
	}
	return res
}
