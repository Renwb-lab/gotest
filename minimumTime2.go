package main

import (
	"sort"
)

// https://leetcode.cn/problems/minimum-time-to-complete-trips/?envType=daily-question&envId=2024-10-07
func minimumTime2(time []int, totalTrips int) int64 {
	sort.Slice(time, func(i, j int) bool {
		return time[i] < time[j]
	})
	f := func(mid int) int {
		ans := 0
		for _, t := range time {
			ans += mid / t
		}
		return ans
	}
	left := 1
	right := totalTrips * time[0]
	for left < right {
		mid := left + (right-left)>>2
		times := f(mid)
		if times < totalTrips {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return int64(left)
}
