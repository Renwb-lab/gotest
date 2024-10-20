package main

import "math"

// https://leetcode.cn/problems/minimum-speed-to-arrive-on-time/?envType=daily-question&envId=2024-10-12
func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	// 将 hour 乘 100 以转为整数
	hr := int(math.Round(hour * 100))
	// 时间必须要大于路程段数减 1
	if hr <= (n-1)*100 {
		return -1
	}
	// 二分
	l, r := 1, 10000000
	for l < r {
		mid := l + (r-l)/2
		// 判断当前时速是否满足时限
		t := 0
		// 前 n-1 段中第 i 段贡献的时间： ceil(dist[i] / mid)
		for i := 0; i < n-1; i++ {
			t += (dist[i]-1)/mid + 1
		}
		// 最后一段贡献的时间： dist[n-1] / mid
		t *= mid
		t += dist[n-1]
		if t*100 <= hr*mid { // 通分以转化为整数比较
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l // 满足条件的最小时速
}
