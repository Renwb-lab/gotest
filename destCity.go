package main

// https://leetcode.cn/problems/destination-city/
func destCity(paths [][]string) string {
	citys := make(map[string]int)
	for _, p := range paths {
		from, to := p[0], p[1]
		if _, ok := citys[from]; !ok {
			citys[from] = 0
		}
		citys[from] += 1
		if _, ok := citys[to]; !ok {
			citys[to] = -1
		}
	}
	for k, v := range citys {
		if v == -1 {
			return k
		}
	}

	return ""
}
