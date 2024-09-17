package main

import "fmt"

func numberOfSequence(n int, sick []int) int {
	// [0 1 2 3 4 5]
	ans := 0
	mod := 10000000007
	m := make(map[int]struct{})
	for _, s := range sick {
		m[s] = struct{}{}
	}
	var dfs func()
	st := make([]int, 0)
	dfs = func() {
		if len(m) == n-1 {
			ans = (ans + 1) % mod
			fmt.Println(st)
			return
		}
		for k := range m {
			if k-1 >= 0 {
				if _, ok := m[k-1]; !ok {
					m[k-1] = struct{}{}
					st = append(st, k-1)
					dfs()
					delete(m, k-1)
					st = st[:len(st)-1]
				}
			}
			if k+1 < n {
				if _, ok := m[k+1]; !ok {
					m[k+1] = struct{}{}
					st = append(st, k+1)
					dfs()
					delete(m, k+1)
					st = st[:len(st)-1]
				}
			}
		}
	}
	dfs()
	return ans % mod
}

func main328() {
	fmt.Println(numberOfSequence(5, []int{0, 4}))
}
